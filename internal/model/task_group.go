package model

type TaskGroup struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string
	GroupLead string
	Tasks     []Task `gorm:"foreignkey:GroupID;references:ID"`
}

type TaskGroupRequest struct {
	Name      *string `json:"name"`
	GroupLead *string `json:"group_lead"`
}

type TaskGroupResponse struct {
	Name      string `json:"name"`
	GroupLead string `json:"group_lead"`
	Tasks     []Task `json:"tasks"`
}
