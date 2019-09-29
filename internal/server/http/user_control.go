package http

import (
	"encoding/json"
	"errors"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"shiji_server/internal/model"
)

func login(c *bm.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	var runErr error = nil
	var k model.User

	// 最后不管是否错误都选择转成json显示
	defer c.JSON(&k, runErr)

	if err != nil {
		log.Info("读取错误")
		runErr = errors.New("-1")
		return
	}

	if err = json.Unmarshal(body, &k); err != nil {
		log.Info("转json错误")
		runErr = errors.New("-1")
		return
	}
}

func register(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
