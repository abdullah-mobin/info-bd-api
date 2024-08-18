package model

type Division struct {
	Id        int        `json:"id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	Area      float32    `json:"area,omitempty"`
	Districts []District `json:"districts,omitempty" gorm:"foreignKey:DivisionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type District struct {
	Id         int     `json:"id,omitempty" gorm:"primaryKey"`
	Name       string  `json:"name,omitempty"`
	DivisionID int     `json:"division_id,omitempty"`
	Area       float32 `json:"area,omitempty"`
	// Division   Division `json:"division,omitempty" gorm:"foreignKey:DivisionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
