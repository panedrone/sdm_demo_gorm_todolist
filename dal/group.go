package dal

// This code was generated by a tool. Don't modify it manually.
// http://sqldalmaker.sourceforge.net

type Group struct {
	GId   int64  `json:"g_id" gorm:"column:g_id;primary_key;auto_increment"`
	GName string `json:"g_name" gorm:"column:g_name;unique_index;not null"`
}