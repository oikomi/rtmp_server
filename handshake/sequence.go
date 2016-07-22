package handshake

import "io"

type Sequence interface {
	Read(r io.Reader) error
	WriteTo(w io.Writer) error
	Next() Sequence
}
