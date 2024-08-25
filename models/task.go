package models

type Task struct {
	ID          int `gorm:"primary key"`
	Title       string
	Description string
	UserID      int  `gorm:"references users(id)" json:"User_ID"`
	User        User `gorm:"foreignKey:UserID" json:"user_name"`
	IsDone      bool `gorm:"default false" json:"Is_Done"`
	IsDeleted   bool `gorm:"default false" json:"Is_Deleted"`
}

func (Task) TableName() string {
	return "tasks"
}
