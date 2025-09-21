package schema

import (
	"time"

	"github.com/lib/pq"
)

type Article struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"type:text;not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;default:now()" json:"updated_at"`
}

func (Article) TableName() string {
	return "articles"
}
