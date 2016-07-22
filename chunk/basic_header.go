package chunk

import (
	"bytes"
	"io"

	"github.com/oikomi/rtmp_server/spec"
)

type BasicHeader struct {
	FormatId byte
	StreamId uint32
}

func (h *BasicHeader) Read(r io.Reader) error {
	b, err := spec.ReadByte(r)
	if err != nil {
		return err
	}
	h.FormatId = b >> 6
	if b&0x3f == 0x3f {
		tail, err := spec.ReadBytes(r, int(2))
		if err != nil {
			return err
		}
		h.StreamId = spec.Uint32(tail) + 64
	} else if b&0x3f == 0 {
		tail, err := spec.ReadBytes(r, int(1))
		if err != nil {
			return err
		}
		h.StreamId = spec.Uint32(tail) + 64
	} else {
		h.StreamId = spec.Uint32([]byte{b & 0x3f})
	}
	return nil
}

func (h *BasicHeader) Write(w io.Writer) error {
	buf := make([]byte, 1)
	buf[0] = h.FormatId << 6
	switch {
	case h.StreamId < 64:
		buf[0] |= (byte(h.StreamId) & 0x3f)
	case h.StreamId < 320:
		buf = append(buf, byte(h.StreamId-64))
	default:
		tmp := new(bytes.Buffer)
		if _, err := spec.PutUint16(uint16(h.StreamId-64), tmp); err !=
			nil {

			return err
		}
		buf[0] |= 63
		buf = append(buf, tmp.Bytes()...)
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}
