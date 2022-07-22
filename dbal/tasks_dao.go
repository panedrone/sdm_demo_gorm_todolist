package dbal

// This code was generated by a tool. Don't modify it manually.
// http://sqldalmaker.sourceforge.net

import "sdm_demo_gorm_todolist/models"

type TasksDao struct {
	ds DataStore
}

// (C)RUD: tasks
// Generated values are passed to DTO/model.

func (dao *TasksDao) CreateTask(p *models.Task) (err error) {
	err = dao.ds.Db().Create(p).Error
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTaskList() (res []*models.Task, err error) {
	err = dao.ds.Db().Find(res).Error
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTask(tId int64) (res *models.Task, err error) {
	res = &models.Task{}
	err = dao.ds.Db().Take(res, tId).Error
	return
}

// CR(U)D: tasks
// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) UpdateTask(p *models.Task) (res int64, err error) {
	db := dao.ds.Db().Save(p)
	err = db.Error
	res = db.RowsAffected
	return
}

// CRU(D): tasks
// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) DeleteTask(tId int64) (res int64, err error) {
	p := &models.Task{
		TId: tId,
	}
	db := dao.ds.Db().Delete(p)
	err = db.Error
	res = db.RowsAffected
	return
}
