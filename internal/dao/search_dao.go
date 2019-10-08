package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"shiji_server/internal/model"
)

var (
	_GetSearchHistory    = "SELECT `History` FROM `search_history` WHERE `id` = ?"
	_UpdateSearchHistory = "UPDATE `search_history` SET `History` = ? WHERE `Id` = ?"
	_AddSearchHistory = "INSERT INTO `search_history` (`Id`, `History`) VALUES(?, ?)"
)

func (d *DaoStruct) GetSearchHistory(c context.Context, history *model.SearchHistory) (err error) {
	err = d.db.QueryRow(c, _GetUser, history.Id).Scan(&history.History)
	if err != nil && err != sql.ErrNoRows {
		log.Error("user.Query error(%v)", err)
		// 这里直接返回错误跟userGot
		return
	}
	return
}

func (d *DaoStruct) UpdateSearchHistory(c context.Context, history *model.SearchHistory) (err error) {
	if _, err = d.db.Exec(c, _UpdateSearchHistory, history.History, history.Id); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return
}

func (d *DaoStruct) AddSearchHistory(c context.Context, user *model.User) (err error) {
	if _, err = d.db.Exec(c, _AddSearchHistory, user.Id, ""); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return
}

func (d *DaoStruct) GetSearchMap(c context.Context, user *model.User) (err error) {
	if _, err = d.db.Exec(c, _AddSearchHistory, user.Id, ""); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return
}

func (d *DaoStruct) GetSearchVectot(c context.Context, user *model.User) (err error) {
	if _, err = d.db.Exec(c, _AddSearchHistory, user.Id, ""); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return
}

func (d *DaoStruct) GetSearchAns(c context.Context, user *model.User) (err error) {
	if _, err = d.db.Exec(c, _AddSearchHistory, user.Id, ""); err != nil {
		log.Error("incr user base err(%v)", err)
		// 返回错误
		return
	}

	// 这里直接返回id
	return
}
