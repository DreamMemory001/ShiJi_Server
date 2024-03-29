package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"shiji_server/internal/model"
	"shiji_server/utils"
)

var (
	_GetUser    = "SELECT `Id`, `Name`, `Password` FROM `user` where `Email` = ?"
	_AddUser    = "INSERT INTO `user` (`Name`, `Email`, `Password`) VALUES(?, ?, ?)"
	_UpdateUser = ""
)

func (d *DaoStruct) GetUser(c context.Context, user *model.User) (id uint, name, password string, err error) {
	err = d.db.QueryRow(c, _GetUser, user.Email).Scan(&id, &name, &password)
	if err != nil && err != sql.ErrNoRows {
		log.Error("user.Query error(%v)", err)
		// 这里直接返回错误跟userGot
		return
	}
	return
}

func (d *DaoStruct) AddUser(c context.Context, user *model.User) (num int64, err error) {
	var res sql.Result
	encodePwd, err := utils.PwdEncode(user.Password)
	if err != nil {
		log.Error("pwd encode error (%v)", err)
		return
	}

	if res, err = d.db.Exec(c, _AddUser, user.Name, user.Email, encodePwd); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return res.LastInsertId()
}
