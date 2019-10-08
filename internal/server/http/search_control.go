package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"shiji_server/internal/dao"
)

func searchHistoryGetter(c *bm.Context) {
}

func searchHistoryUpdater(c *bm.Context) {
}

func searchMapGetter(c *bm.Context) {

}

func searchAnsGetter(c *bm.Context) {
	key := c.Request.Form.Get("key")
	daoIns := dao.New()
	// 这里直接截断
	ans, err := daoIns.GetSearchAns(c, key)
	// 最后不管是否错误都选择转成json显示
	c.JSON(&ans, err)
}

func searchWordVecGetter(c *bm.Context) {
}
