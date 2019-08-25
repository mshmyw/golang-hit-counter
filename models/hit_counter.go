package models

type HitCounter struct {
	Model
	URL   string `gorm:"type:varchar(256)" json:"url,omitempty"`
	Count *int   `json:"count,omitempty"` // 指针是为了序列化的时候0值生效
	Memo   string `gorm:"type:varchar(256)" json:"memo,omitempty"`
}
