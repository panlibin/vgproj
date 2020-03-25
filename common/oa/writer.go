package oa

import (
	"fmt"
	"strings"
	"sync/atomic"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo"
	"github.com/panlibin/virgo/database"
)

// Writer oa记录器
type Writer struct {
	p       virgo.IProcedure
	pDb     *database.Mysql
	connNum int32
	dbIdx   int32
}

// NewWriter 新建oa记录器
func NewWriter(p virgo.IProcedure, pDb *database.Mysql, connNum int32) *Writer {
	return &Writer{
		p:       p,
		pDb:     pDb,
		connNum: connNum,
	}
}

// Write 写数据
func (w *Writer) Write(tableName string, args ...interface{}) {
	lenArgs := len(args)
	if lenArgs <= 0 {
		logger.Errorf("oa write %s, args empty", tableName)
		return
	}

	strPlaceholder := "?"
	if lenArgs > 1 {
		strPlaceholder += strings.Repeat(",?", lenArgs-1)
	}

	strSQL := fmt.Sprintf("insert into `%s` values(default,%s)", tableName, strPlaceholder)
	dbIdx := atomic.AddInt32(&w.dbIdx, 1)
	if dbIdx >= w.connNum {
		dbIdx -= w.connNum
	}
	w.pDb.AsyncExec(nil, nil, uint32(dbIdx), strSQL, args...)
}
