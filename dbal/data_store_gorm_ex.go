package dbal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

var _ds = &_DS{} // completely private for this package

func (ds *_DS) initDb() (err error) {
	ds.db, err = gorm.Open(sqlite.Open("./todolist.sqlite3"), &gorm.Config{
		//		Logger: logger.Default.LogMode(logger.Info),
	})
	return
}

//func Db() *gorm.DB {
//	return _ds.db
//}

func OpenDB() error {
	return _ds.Open()
}

func CloseDB() error {
	return _ds.Close()
}

func NewGroupsDao() *GroupsDao {
	return &GroupsDao{ds: _ds}
}

func NewTasksDao() *TasksDao {
	return &TasksDao{ds: _ds}
}
