package route

import (
	"github.com/zyin-c/zyinc-daemon/pkg/handlers"
	"github.com/zyin-c/zyinc-daemon/pkg/server"
)

func setUpHandlers(app *server.Server) {
	app.On("up", handlers.UpEventHandler)
	app.On("stop", handlers.StopEventHandler)
	app.On("restart", handlers.RestartEventHandler)
	app.On("down", handlers.KillEventHandler)
	app.On("kill", handlers.KillEventHandler)
	app.On("ps", handlers.PsEventHandler)
}
