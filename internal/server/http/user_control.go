package http

import (
	"encoding/json"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"shiji_server/internal/model"
)

func covertBody2JSON(body *http.Request, c interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body.Body)

	if err != nil {
		log.Info("读取错误")
		return http.ErrServerClosed
	}

	if err = json.Unmarshal(bodyBytes, &c); err != nil {
		log.Info("转json错误")
		return http.ErrServerClosed
	}

	return nil
}

func login(c *bm.Context) {
	var k model.User

	err := covertBody2JSON(c.Request, &k)

	// 最后不管是否错误都选择转成json显示
	defer c.JSON(&k, err)
}

func register(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
