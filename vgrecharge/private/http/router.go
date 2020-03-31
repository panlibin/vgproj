package http

import (
	"github.com/panlibin/virgo/util/nethelper"
)

// NewHTTPServer 创建HTTPServer
func NewHTTPServer() *nethelper.HTTPServer {
	pServer := nethelper.NewHTTPServer()
	// pServer.Handle("/register", handleRegister)
	// pServer.Handle("/login", handleLogin)
	// pServer.Handle("/server_list", handleServerList)

	return pServer
}
