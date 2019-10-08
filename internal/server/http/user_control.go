package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"net/http"
	"shiji_server/internal/dao"
	"shiji_server/internal/model"
	"shiji_server/utils"
)

func covertBody2User(body *http.Request, user *model.User) {
	user.Name = body.Form.Get("Name")
	user.Email = body.Form.Get("Email")
	user.Password = body.Form.Get("Password")
}

func login(c *bm.Context) {
	var k model.User
	covertBody2User(c.Request, &k)
	daoIns := dao.New()
	id, name, password, err := daoIns.GetUser(c, &k)
	k.Id = id
	k.Name = name
	err = utils.PwdDecode(password, k.Password)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&k, err)
}

func register(c *bm.Context) {
	var k model.User
	covertBody2User(c.Request, &k)
	daoIns := dao.New()
	// 这里直接截断
	id, err := daoIns.AddUser(c, &k)
	k.Id = uint(id)
	err = daoIns.AddSearchHistory(c, &k)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&k, err)
}
