package mail

import (
	"sort"
	imail "vgproj/vggame/public/game/mail"
)

type MailArray []*imail.GlobalMailDef

func (m MailArray) Len() int {
	return len(m)
}

func (m MailArray) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MailArray) Less(i, j int) bool {
	return m[i].GlobalMailId < m[j].GlobalMailId
}

func (m MailArray) Search(globalMailId int32) int {
	return sort.Search(m.Len(), func(i int) bool {
		return m[i].GlobalMailId >= globalMailId
	})
}
