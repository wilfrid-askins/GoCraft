package packets

import (
	"GoCraft/pkg/gocraft/packets/client"
	"GoCraft/pkg/gocraft/packets/server"
	"GoCraft/pkg/gocraft/packets/types"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
)

type PacketHandler func(Packet)

var handlers = map[types.VarInt]map[types.VarInt]PacketHandler{}

type Session struct {
	state      types.VarInt
	summariser Summariser
	logger     *zap.SugaredLogger
}

type Summariser interface {
	GetSummary() string
}

func NewSession(summariser Summariser, logger *zap.SugaredLogger) Session {
	return Session{state: types.VarInt(HANDSHAKE), summariser: summariser, logger: logger}
}

func (s *Session) OnPacket(p Packet) {
	handler := handlers[s.state][p.GetID()]
	handler(p)
}

func (s *Session) GetState() types.VarInt {
	return s.state
}

func (s *Session) SetOutput(write WriteFunc, closeFunc ActionFunc, compress ActionFunc) {
	handlers[HANDSHAKE] = map[types.VarInt]PacketHandler{
		0x0: func(p Packet) {
			hs := p.(*client.Handshake)
			s.state = hs.NextState
			s.logger.Info("Handshake state change")
		},
	}

	handlers[STATUS] = map[types.VarInt]PacketHandler{
		0x0: func(p Packet) {
			s.logger.Info("Status packet received")
			summary := s.summariser.GetSummary()

			// Write response
			response := &server.Response{}
			response.JsonResponse = types.String(summary)

			if err := write(response); err != nil {
				s.logger.Error("failed to write status response", zap.String("msg", errors.Message(err)))
			}
		},
		0x1: func(p Packet) {
			ping := p.(*client.Ping)
			pong := &server.Pong{Payload: ping.Payload}
			err := write(pong)
			if err != nil {
				s.logger.Error("failed to write pong packet", zap.String("msg", errors.Message(err)))
			}
			s.logger.Info("Received Ping, sent Pong")
			closeFunc()
		},
	}

	handlers[LOGIN] = map[types.VarInt]PacketHandler{
		0x0: func(p Packet) {
			lp := p.(*client.LoginStart)
			s.logger.Info("Login start from " + string(lp.Name))

			// Encryption request
			//key := getPublicKey()
			//
			//encryptionReq := &server.EncryptionRequest{
			//	ServerID:          "                    ",
			//	PublicKeyLength:   types.VarInt(len(key)),
			//	PublicKey:         key,
			//	VerifyTokenLength: 4,
			//	VerifyToken:       []byte("aaaa"),
			//}
			//err := Write(s.out, encryptionReq)
			//if err != nil {
			//	s.logger.Error("failed to write encryption request", zap.String("msg", errors.Message(err)))
			//}

			//compression := &server.SetCompression{Threshold: 0}
			//if err := write(compression); err != nil {
			//	s.logger.Error("failed to write set compression", zap.String("msg", errors.Message(err)))
			//}
			//compress()

			success := &server.LoginSuccess{
				UUID:     "26173bb2-50fc-45dc-8372-94c47e1d2669",
				Username: "HotFix          ", // TODO
			}
			if err := write(success); err != nil {
				s.logger.Error("failed to write login success", zap.String("msg", errors.Message(err)))
			}

			// TODO send join game packet
			// held item change
			// player position and look
			// chunk data
			// etc
		},
		0x1: func(p Packet) {
			lp := p.(*client.EncryptionResponse)
			s.logger.Info("Encryption response ss: " + string(lp.SharedSecret))
			// Optional: Send set compression

			success := &server.LoginSuccess{
				UUID:     "31574af2-21ff-53hf-5832-94b63e5o6678",
				Username: "HotFix", // TODO
			}
			if err := write(success); err != nil {
				s.logger.Error("failed to write login success", zap.String("msg", errors.Message(err)))
			}
		},
	}

	handlers[PLAY] = map[types.VarInt]PacketHandler{
		0x3: func(packet Packet) {
			p := packet.(*client.ChatMessage)
			s.logger.Info("ChatMessage received: " + string(p.Message))
		},
	}
}
