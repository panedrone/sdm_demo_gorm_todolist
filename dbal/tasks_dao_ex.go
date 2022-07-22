package dbal

import "sdm_demo_gorm_todolist/models"

func (dao *TasksDao) ReadGroupTasks(gId int64) (res []*models.Task, err error) {
	err = dao.ds.Db().Table("tasks").Where("g_id = ?", gId).Order("t_id").
		Select("t_id", "t_date", "t_subject", "t_priority").Find(&res).Error
	return
}
