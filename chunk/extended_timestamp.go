package chunk

import (
	"io"

	"github.com/oikomi/rtmp_server/spec"
)

type ExtendedTimestamp struct {
	Delta uint32
}

func (t *ExtendedTimestamp) Read(r io.Reader) error {
	buf, err := spec.ReadBytes(r, 4)
	if err != nil {
		return err
	}

	t.Delta = spec.Uint32(buf)
	return nil
}

func (t *ExtendedTimestamp) Write(w io.Writer) error {
	if _, err := spec.PutUint32(t.Delta, w); err != nil {
		return err
	}
	return nil
}
