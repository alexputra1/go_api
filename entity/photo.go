package entity

type Photo struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:text" json:"photourl"`
	UserID   uint64 `gorm:"not null" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constrain:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}