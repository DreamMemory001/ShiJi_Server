package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"shiji_server/internal/model"
)

var (
	_GetUser    = "SELECT * FROM `user`"
	_AddUser    = "INSERT INTO `user` (`Email`, `Password`) VALUES(?, ?)"
	_UpdateUser = ""
)

func (d *DaoStruct) GetUser(c context.Context, user *model.User) (re string, err error) {
	err = d.db.QueryRow(c, _GetUser).Scan(&re)
	if err != nil && err != sql.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		// 这里直接返回错误跟空re了
		return
	}
	return re, nil
}

func (d *DaoStruct) AddUser(c context.Context, user *model.User) (num int64, err error) {
	var res sql.Result

	if res, err = d.db.Exec(c, _AddUser, user.Email, user.Password); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return res.LastInsertId()
}
