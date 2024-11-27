package model

import "gorm.io/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State *uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	query := db.Model(&t)
	if t.Name != "" {
		query = query.Where("name = ?", t.Name)
	}
	query = query.Where("state = ?", t.State)
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int64) ([]*Tag, error) {
	var tags []*Tag
	query := db.Model(&t)
	if pageOffset >= 0 && pageSize > 0 {
		query = query.Offset(int(pageOffset)).Limit(int(pageSize))
	}
	if t.Name != "" {
		query = query.Where("name = ?", t.Name)
	}
	query = query.Where("state = ?", t.State)
	if err := query.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&t).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}
