package dao

import (
	"context"
	"shiji_server/internal/model"
)

var (
	_GetSearchHistory    = "SELECT * FROM `search_history` where `id` = ?"
	_UpdateSearchHistory    = "INSERT INTO `user` (`Email`, `Password`) VALUES(?, ?)"
)

func (d *DaoStruct) GetSearchHistory(c context.Context, history *model.SearchHistory) (userGot model.User, err error) {
	err = d.db.QueryRow(c, _GetUser, user.Email).Scan(&userGot.Id, &userGot.Email, &userGot.Password)
	if err != nil && err != sql.ErrNoRows {
		log.Error("user.Query error(%v)", err)
		// 这里直接返回错误跟userGot
		return
	}
	return
}

