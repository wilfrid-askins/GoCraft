package packets

import (
	"GoCraft/pkg/gocraft/packets/client"
	"GoCraft/pkg/gocraft/packets/server"
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
)

type PacketHandler func(Packet)

var handlers = map[types.VarInt]map[types.VarInt]PacketHandler{}

type Session struct {
	state      types.VarInt
	out        *bufio.Writer
	summariser Summariser
	logger     *zap.Logger
}

type Summariser interface {
	GetSummary() string
}

func NewSession(summariser Summariser, logger *zap.Logger) Session {
	return Session{types.VarInt(HANDSHAKE), nil, summariser, logger}
}

func (s *Session) OnPacket(p Packet) {
	handler := handlers[s.state][p.GetID()]
	handler(p)
}

func (s *Session) GetState() types.VarInt {
	return s.state
}

func (s *Session) SetOutput(out *bufio.Writer) {
	s.out = out

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

			if err := Write(s.out, response); err != nil {
				s.logger.Error("failed to write status response", zap.String("msg", errors.Message(err)))
			}
		},
		0x1: func(p Packet) {
			ping := p.(*client.Ping)
			pong := &server.Pong{Payload: ping.Payload}
			err := Write(s.out, pong)
			if err != nil {
				s.logger.Error("failed to write pong packet", zap.String("msg", errors.Message(err)))
			}
			s.logger.Info("Received Ping, sent Pong")
			// TODO close connection
		},
	}

	handlers[LOGIN] = map[types.VarInt]PacketHandler{
		0x0: func(p Packet) {
			lp := p.(*client.LoginStart)
			s.logger.Info("Login start from " + string(lp.Name))

			encryptionReq := &server.EncryptionRequest{
				ServerID:          "                   ",
				PublicKeyLength:   3,
				PublicKey:         []byte("fgf"),
				VerifyTokenLength: 3,
				VerifyToken:       []byte("fgf"),
			}
			err := Write(s.out, encryptionReq)
			if err != nil {
				s.logger.Error("failed to write encryption request", zap.String("msg", errors.Message(err)))

			}
			// Login start
			// Optional: Send encryption request
		},
		0x1: func(p Packet) {
			lp := p.(*client.EncryptionResponse)
			s.logger.Info("Encryption response ss: " + string(lp.SharedSecret))
			// Optional: Send set compression

			success := &server.LoginSuccess{
				UUID:     "31574af2-21ff-53hf-5832-94b63e5o6678",
				Username: "HotFix", // TODO
			}
			err := Write(s.out, success)
			if err != nil {
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
