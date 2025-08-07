package model

type TaskGroup struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string
	GroupLead string
	Tasks     []Task `gorm:"foreignkey:GroupID;references:ID"`
}
