package orm

import (
	"gorm.io/gorm"
	"time"
)

// Model is the interface for all models.
// It only requires an Identity() method to return the primary key field
// name and value.
type Model interface {
	// Identity returns the primary key field of the model.
	// A very common case is that the primary key field is ID.
	Identity() (fieldName string, value any)
}

// BasicModel implements Model interface with an auto increment primary key ID.
//
// BasicModel is actually the gorm.Model struct which contains the following
// fields:
//
//	ID, CreatedAt, UpdatedAt, DeletedAt
//
// It is a good idea to embed this struct as the base struct for all models:
//
//	type User struct {
//	  orm.BasicModel
//	}
type BasicModel gorm.Model

func (m BasicModel) Identity() (fieldName string, value any) {
	return "ID", m.ID
}

type BaseModel struct {
	gorm.Model
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
