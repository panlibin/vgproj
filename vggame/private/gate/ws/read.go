package ws

import (
	"encoding/binary"
	"fmt"
	"vgproj/common/util"
)

// Errors used by frame reader.
var (
	ErrHeaderLengthMSB        = fmt.Errorf("header error: the most significant bit must be 0")
	ErrHeaderLengthUnexpected = fmt.Errorf("header error: unexpected payload length bits")
	ErrHeaderNotReady         = fmt.Errorf("header error: not enough")
)

// ReadHeader reads a frame header from r.
func ReadHeader(b []byte) (h Header, extra int) {
	h.Fin = b[0]&bit0 != 0
	h.Rsv = (b[0] & 0x70) >> 4
	h.OpCode = OpCode(b[0] & 0x0f)

	if b[1]&bit0 != 0 {
		h.Masked = true
		extra += 4
	}

	length := b[1] & 0x7f
	h.Length = int64(length)
	switch {
	case length == 126:
		extra += 2

	case length == 127:
		extra += 8
	}
	return
}

func ReadHeaderEx(b []byte, h *Header) {
	switch {
	case h.Length == 126:
		h.Length = int64(binary.BigEndian.Uint16(b[:2]))
		b = b[2:]

	case h.Length == 127:
		if b[0]&0x80 != 0 {
			b[0] &= 0x7f
		}
		h.Length = int64(binary.BigEndian.Uint64(b[:8]))
		b = b[8:]
	}

	if h.Masked {
		copy(h.Mask[:], b[:4])
	}
}

// ParseCloseFrameData parses close frame status code and closure reason if any provided.
// If there is no status code in the payload
// the empty status code is returned (code.Empty()) with empty string as a reason.
func ParseCloseFrameData(payload []byte) (code StatusCode, reason string) {
	if len(payload) < 2 {
		// We returning empty StatusCode here, preventing the situation
		// when endpoint really sent code 1005 and we should return ProtocolError on that.
		//
		// In other words, we ignoring this rule [RFC6455:7.1.5]:
		//   If this Close control frame contains no status code, _The WebSocket
		//   Connection Close Code_ is considered to be 1005.
		return
	}
	code = StatusCode(binary.BigEndian.Uint16(payload))
	reason = util.BytesToString(payload[2:])
	return
}
