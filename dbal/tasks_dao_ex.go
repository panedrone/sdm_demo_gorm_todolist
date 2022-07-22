package dbal

import "sdm_demo_gorm_todolist/models"

// this is an extension of generated TasksDao created by hand

func (dao *TasksDao) ReadGroupTasks(gId int64) (res []*models.Task, err error) {
	// don't fetch comments for list:
	err = dao.ds.Db().Table("tasks").Where("g_id = ?", gId).Order("t_id").
		Select("t_id", "t_date", "t_subject", "t_priority").Find(&res).Error
	return
}
