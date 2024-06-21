package global

import (
	"time"

	"gorm.io/gorm"
)

// 软
type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// 硬
type GVA_MODEL_HARD struct {
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
