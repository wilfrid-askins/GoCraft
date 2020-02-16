package packets

import (
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
	"bytes"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
	"io"
	"net"
)

type (
	Handler struct {
		receiver Receiver
		logger *zap.Logger
	}

	Receiver interface {
		OnPacket(Packet)
		GetState() types.VarInt
		SetOutput(out *bufio.Writer)
	}
)

func NewHandler(receiver Receiver, logger *zap.Logger) Handler {
	return Handler{receiver, logger}
}

func (h *Handler) Listen(conn net.Conn) {
	input := bufio.NewReader(conn)
	defer conn.Close()

	output := bufio.NewWriter(conn)
	h.receiver.SetOutput(output)

	for {
		lenVal, err := types.VarIntDefault.Read(input)
		length := lenVal.(types.VarInt)

		if err == io.EOF {
			continue
		}

		buf := make([]byte, length)
		_, err = io.ReadFull(input, buf)
		if err != nil {
			h.logger.Error("failed to read packet data", zap.String("msg", err.Error()))
		}

		payload := bufio.NewReader(bytes.NewReader(buf))
		packetType, err := types.VarIntDefault.Read(payload)
		if err != nil {
			h.logger.Error("failed to read packet type", zap.String("msg", errors.Message(err)))
		}

		packetID := packetType.(types.VarInt)
		h.logger.Info("Received packet", zap.Int("id", int(packetID)), zap.Int("state", int(h.receiver.GetState())))
		packet := StateToPacketLookup[h.receiver.GetState()][packetID]

		h.logger.Info("Reading payload")
		err = packet.Read(payload)

		if err != nil {
			h.logger.Error("failed to read packet payload", zap.String("msg", errors.Message(err)))
		}

		h.receiver.OnPacket(packet)
	}
}
