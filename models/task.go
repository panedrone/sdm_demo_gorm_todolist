package models

// This code was generated by a tool. Don't modify it manually.
// Additional (hand-coded) methods can be implemented in a separate file like <this_file>_ex.go.
// http://sqldalmaker.sourceforge.net

type Task struct {
	TId       int64  `json:"t_id" gorm:"column:t_id;primary_key;auto_increment"`
	GId       int64  `json:"g_id" gorm:"column:g_id;not null"`
	TPriority int64  `json:"t_priority" gorm:"column:t_priority;not null"`
	TDate     string `json:"t_date" gorm:"column:t_date;not null"`
	TSubject  string `json:"t_subject" gorm:"column:t_subject;not null"`
	TComments string `json:"t_comments" gorm:"column:t_comments"`
}
