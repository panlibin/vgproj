package filler

import (
	"vgproj/proto/msg"
)

type Params struct {
	Type    int32
	Content string
}

func FormatToMessage(params []*Params) []*msg.CHARACTERS_PARAM {
	ret := make([]*msg.CHARACTERS_PARAM, len(params))
	for i, v := range params {
		p := new(msg.CHARACTERS_PARAM)
		p.ParamType = v.Type
		p.ParamValue = v.Content
		ret[i] = p
	}
	return ret
}
