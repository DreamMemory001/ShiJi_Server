package dao

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"
	"shiji_server/internal/model"
)

var (
	_GetSearchHistory    = "SELECT `History` FROM `search_history` WHERE `Id` = ?"
	_UpdateSearchHistory = "UPDATE `search_history` SET `History` = ? WHERE `Id` = ?"
	_AddSearchHistory = "INSERT INTO `search_history` (`Id`, `History`) VALUES(?, ?)"

	_GetSearchMap = ""
	_GetSearchVector = "SELECT BookName, SUM(Counts) AS Sums FROM (SELECT BookName, COUNT(CutWord = ?)" +
		" AS Counts FROM (SELECT BookId, CutWord, BookName FROM ancient_book_cut, ancient_book_all WHERE ancient_book_cut.BookId = ancient_book_all.Id) " +
		" AS p GROUP BY BookId) AS pp GROUP BY BookName"
	// 嵌套查询
	_GetSearchAns = "SELECT `BookName`, `Title`, `Content` FROM `ancient_book_all` WHERE `Id` = ANY(" +
		"SELECT `BookId` FROM `ancient_book_cut` WHERE `CutWord` = ?)"
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
	if _, err = d.db.Exec(c, _AddSearchHistory, user.Id, "|"); err != nil {
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

func (d *DaoStruct) GetSearchVector(c context.Context, searchWord string) (answer []model.SearchVectorResult, err error) {
	rows, err := d.db.Query(c, _GetSearchVector, searchWord)
	if err != nil {
		log.Error("query  error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmp model.SearchVectorResult
		if err = rows.Scan(&tmp.BookName, &tmp.Sums); err != nil {
			log.Error("scan demo log error(%v)", err)
			return
		}
		answer = append(answer, tmp)
	}
	return
}

// 查询所在词在数据库中的所在
func (d *DaoStruct) GetSearchAns(c context.Context, searchWord string) (answer []model.SearchAnsResult, err error) {
	rows, err := d.db.Query(c, _GetSearchAns, searchWord)
	if err != nil {
		log.Error("query  error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmp model.SearchAnsResult
		if err = rows.Scan(&tmp.BookName, &tmp.Title, &tmp.Content); err != nil {
			log.Error("scan demo log error(%v)", err)
			return
		}
		answer = append(answer, tmp)
	}
	return
}
