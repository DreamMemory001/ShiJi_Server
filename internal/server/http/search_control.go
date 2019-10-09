package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"shiji_server/internal/dao"
	"shiji_server/internal/model"
	"strconv"
)

func searchHistoryGetter(c *bm.Context) {
	var history model.SearchHistory
	Id := c.Request.Form.Get("key")
	uid, err := strconv.ParseUint(Id, 10, 32)
	history.Id = uint(uid)
	daoIns := dao.New()
	// 这里直接截断
	//log.Info(history.Id)
	err = daoIns.GetSearchHistory(c, &history)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&history, err)
}

func searchHistoryUpdater(c *bm.Context) {
}

func searchMapGetter(c *bm.Context) {
}

func searchAnsGetter(c *bm.Context) {
	key := c.Request.Form.Get("Key")
	daoIns := dao.New()
	// 这里直接截断
	ans, err := daoIns.GetSearchAns(c, key)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&ans, err)
}

func searchWordVecGetter(c *bm.Context) {
	key := c.Request.Form.Get("Key")
	daoIns := dao.New()
	// 这里直接截断
	ans, err := daoIns.GetSearchVector(c, key)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&ans, err)
}
