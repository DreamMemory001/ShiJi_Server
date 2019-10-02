package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"shiji_server/internal/model"
)

var (
	_GetUser    = ""
	_AddUser    = "INSERT INTO `user` (`Email`, `Password`) VALUES(?, ?)"
	_UpdateUser = ""
)

func (d *DaoStruct) GetUser(c context.Context, user *model.User) (re string, err error) {
	err = d.db.QueryRow(c, "SELECT * FROM `user`").Scan(&re)
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		// 这里直接返回错误跟空re了
		return
	}
	return re, nil
}

func (d *DaoStruct) AddUser(c context.Context, user *model.User) (err error) {
	var res sql.Result

	if res, err = d.db.Exec(c, _AddUser, user.Email, user.Password); err != nil {
		log.Error("incr user base err(%v)", err)
		return
	}
	return nil
}
