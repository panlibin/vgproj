package gate

import (
	"bytes"
	"errors"
	"vgproj/vggame/private/gate/ws"

	"github.com/panlibin/gnet"
)

const (
	wantHead    = 0
	wantPayload = 1
	wantClose   = 2
)

var (
	ErrTooBigHeader  = errors.New("too big header")
	ErrTooBigPayload = errors.New("too big payload")
)

type WsCodec struct {
	g *Gate
	u ws.Upgrader
}

func (c *WsCodec) Encode(conn gnet.Conn, buf []byte) (out []byte, err error) {
	out = buf
	return
}

func (c *WsCodec) Decode(conn gnet.Conn) (out []byte, err error) {
	ctx := conn.Context()
	if ctx == nil {
		var again bool
		out, err, again = c.upgrade(conn)
		if err == nil && !again {
			pConnection := c.g.OnNewConnection(conn)
			conn.SetContext(pConnection)
			return
		}
	} else {
		pConnection := ctx.(*Connection)
		for {
			// 读取包头
			if pConnection.wantType == wantHead {
				// 读取2字节包头
				if pConnection.wantLen == 0 {
					s, b := conn.ReadN(2)
					if s == 0 || b == nil {
						return
					}
					header, extra := ws.ReadHeader(b)

					pConnection.header = &header
					if extra == 0 {
						pConnection.wantType = wantPayload
					} else {
						pConnection.wantLen = extra
					}
					conn.ShiftN(2)
					// 读取额外包头
				} else {
					s, b := conn.ReadN(pConnection.wantLen)
					if s == 0 || b == nil {
						return
					}
					ws.ReadHeaderEx(b, pConnection.header)
					conn.ShiftN(pConnection.wantLen)
					pConnection.wantType = wantPayload
					pConnection.wantLen = 0
				}
				// 读取内容
			} else {
				if pConnection.header.Length > 1024*500 {
					err = ErrTooBigPayload
					pConnection.wantType = wantClose
					out = make([]byte, 0)
					return
				}

				s, b := conn.ReadN(int(pConnection.header.Length))
				if s == 0 || b == nil {
					return
				}

				if pConnection.header.Masked {
					ws.Cipher(b, pConnection.header.Mask, 0)
				}
				out = b
				conn.ShiftN(int(pConnection.header.Length))
				pConnection.wantType = wantHead
				pConnection.wantLen = 0
			}
		}
	}
	return
}

func (c *WsCodec) upgrade(conn gnet.Conn) (out []byte, err error, again bool) {
	b := conn.Read()

	idx := bytes.Index(b, []byte("\r\n\r\n"))
	if idx == -1 {
		if len(b) > 1024 {
			err = ErrTooBigHeader
			conn.ResetBuffer()
		} else {
			again = true
		}
		return
	}
	out, _, err = c.u.Upgrade(b[:idx])
	conn.ShiftN(idx + 4)
	return
}
