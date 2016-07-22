package handshake

import (
	"fmt"
	"io"
)

const (
	SupportedRTMPVersion byte = 3
)

type VerisonSequence struct {
	Supported byte
}

var _ Sequence = new(VerisonSequence)

func NewVersionSequence() *VerisonSequence {
	return &VerisonSequence{
		Supported: SupportedRTMPVersion,
	}
}

func (v *VerisonSequence) Read(r io.Reader) error {
	var b [1]byte
	if _, err := r.Read(b[:]); err != nil {
		return err
	}

	if b[0] != v.Supported {
		return fmt.Errorf(
			"rtmp/handshake: unsupported version %v", b[0])
	}

	return nil
}

func (v *VerisonSequence) WriteTo(w io.Writer) error {
	if _, err := w.Write([]byte{v.Supported}); err != nil {
		return err
	}

	return nil
}

// Next returns the ClientAckSequence, which is the next step in the RTMP
// handshake, according to the specification.
func (v *VerisonSequence) Next() Sequence {
	return NewClientAckSequence()
}
