package handshake

import "io"

type Param struct {
	Conn    io.ReadWriter
	Initial Sequence
}
