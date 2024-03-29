package http

import (
	"net/http"

	"shiji_server/internal/model"
	"shiji_server/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/test")
	{
		g.GET("/start", howToStart)
	}

	user_control := e.Group("/user")
	{
		user_control.POST("/login", login)
		user_control.POST("/register", register)
	}

	search_control := e.Group("/search")
	{
		search_control.POST("/get_history", searchHistoryGetter)
		search_control.POST("/update_history", searchHistoryUpdater)

		search_control.POST("/ans", searchAnsGetter)
		search_control.POST("/map", searchMapGetter)
		search_control.POST("/vec", searchWordVecGetter)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
