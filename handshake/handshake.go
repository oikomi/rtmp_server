package handshake

import "io"

type Handshaker struct {
	rw      io.ReadWriter
	current Sequence
}

func With(p *Param) *Handshaker {
	h := &Handshaker{
		rw: p.Conn,
	}
	if p.Initial != nil {
		h.current = p.Initial
	} else {
		h.current = NewVersionSequence()
	}
	return h
}

func (h *Handshaker) Handshake() (err error) {
	for ; h.current != nil; h.current = h.current.Next() {
		if err = h.current.Read(h.rw); err != nil {
			return
		}
		if err := h.current.WriteTo(h.rw); err != nil {
			return
		}
	}
	return
}
