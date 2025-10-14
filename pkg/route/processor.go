package route

import "github.com/zyin-c/zyinc-daemon/pkg/server"

func StartServer() error {
	app := server.NewServer(server.ServerConfig{})
	setUpHandlers(app)
	return app.Start()
}
