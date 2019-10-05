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
	return
}

