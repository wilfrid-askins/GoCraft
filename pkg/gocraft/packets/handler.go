package packets

import (
	"GoCraft/pkg/gocraft/packets/types"
	"bufio"
	"bytes"
	"compress/zlib"
	"github.com/icelolly/go-errors"
	"go.uber.org/zap"
	"io"
	"net"
)

type (
	Handler struct {
		receiver   Receiver
		logger     *zap.SugaredLogger
		compressed bool
	}

	WriteFunc func(packet Packet) error
	ActionFunc func()

	Receiver interface {
		OnPacket(Packet)
		GetState() types.VarInt
		SetOutput(write WriteFunc, close ActionFunc, compress ActionFunc)
	}
)

func NewHandler(receiver Receiver, logger *zap.SugaredLogger) Handler {
	return Handler{receiver: receiver, logger: logger}
}

func (h *Handler) Listen(conn net.Conn) {
	input := bufio.NewReader(conn)
	outWriter := bufio.NewWriter(conn)
	defer conn.Close()
	closed := false

	output := func(p Packet) error {
		return h.Write(outWriter, p)
	}
	close := func() {
		if err := conn.Close(); err != nil {
			h.logger.Errorw("failed to close connection", "msg", err.Error())
		}
		closed = true
	}
	compress := func() {
		h.compressed = true
	}
	h.receiver.SetOutput(output, close, compress)

	for !closed {
		lenVal, err := types.VarIntDefault.Read(input)
		if err == io.EOF {
			continue
		}

		packetLen := lenVal.(types.VarInt)
		packetBuf := make([]byte, packetLen)
		_, err = io.ReadFull(input, packetBuf)
		if err != nil {
			h.logger.Errorw("failed to read packet", "msg", err.Error())
			continue
		}

		packetRead := bufio.NewReader(bytes.NewReader(packetBuf))
		if !h.compressed {
			// uncompressed
			h.readPacketData(packetRead)
			continue
		}

		dataLen, err := types.VarIntDefault.Read(packetRead)
		if err != nil {
			h.logger.Errorw("failed to read data length", "msg", err.Error())
			continue
		}

		if dataLen == 0 {
			// uncompressed
			h.readPacketData(packetRead)
			continue
		}

		compressedData, err := zlib.NewReader(packetRead)
		if err != nil {
			h.logger.Errorw("failed to create zlib reader", "msg", err.Error())
			continue
		}

		data := make([]byte, 0)
		if _, err = io.Copy(bytes.NewBuffer(data), compressedData); err != nil {
			h.logger.Errorw("failed to read compressed data", "msg", err.Error())
			continue
		}

		dataRead := bufio.NewReader(bytes.NewReader(data))
		h.readPacketData(dataRead)
	}
}

func (h *Handler) readPacketData(data *bufio.Reader) {
	packetType, err := types.VarIntDefault.Read(data)
	if err != nil {
		h.logger.Errorw("failed to read packet type", "msg", errors.Message(err))
		return
	}

	packetID := packetType.(types.VarInt)
	h.logger.Infow("Received packet", "id", int(packetID), "state", int(h.receiver.GetState()))
	packet, pres := StateToPacketLookup[h.receiver.GetState()][packetID]

	if !pres {
		h.logger.Errorw("failed to find handler for packet", "packetID", packetID)
		return
	}

	// h.logger.Info("Reading data")
	err = packet.Read(data)
	if err != nil {
		h.logger.Errorw("failed to read packet data", "msg", errors.Message(err))
		return
	}

	h.receiver.OnPacket(packet)
}

func (h *Handler) Write(out *bufio.Writer, p Packet) error {
	dataBuf := &bytes.Buffer{}
	dataWriter := bufio.NewWriter(dataBuf)

	// write packetID
	pType := p.GetID()
	if err := pType.Write(dataWriter); err != nil {
		return err
	}

	// write payload
	if err := p.Write(dataWriter); err != nil {
		return err
	}
	if err := dataWriter.Flush(); err != nil {
		return err
	}

	if !h.compressed {
		// write length
		length := types.VarInt(dataBuf.Len())
		if err := length.Write(out); err != nil {
			return err
		}

		// write packetID + payload
		if _, err := out.Write(dataBuf.Bytes()); err != nil {
			return err
		}
		if err := out.Flush(); err != nil {
			return err
		}

		return nil
	}

	// compress data

	// packet length
	// data length
	// data

	panic("not yet implemented")

	return nil
}
