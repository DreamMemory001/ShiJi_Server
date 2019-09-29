package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"shiji_server/internal/model"
)

func login(c *bm.Context) {
	query := c.Request.URL.Query()
	id := query.Get("id")
	email := query.Get("email")
	password := query.Get("password")

	k := &model.User{
		Id:       id,
		Email:    email,
		Password: password,
	}
	c.JSON(k, nil)
}

func register(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
