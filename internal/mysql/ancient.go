package mysql

import "github.com/jinzhu/gorm"

type Ancient struct {
	ID      int    `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	Title   string `gorm:"type:varchar(64);index:title_idx;unique_index:tad_idx;not null" json:"title,omitempty"`
	Author  string `gorm:"type:varchar(64);index:author_idx;unique_index:tad_idx;not null" json:"author,omitempty"`
	Dynasty string `gorm:"type:varchar(32);index:dynasty_idx;unique_index:tad_idx;not null" json:"dynasty,omitempty"`
	Content string `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
}

func (m *Mysql) SelectAncients(offset int, limit int) ([]*Ancient, error) {
	var ancients []*Ancient

	if err := m.db.Select("id, title, author, dynasty").Order("id").Offset(offset).Limit(limit).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return ancients, nil
}

func (m *Mysql) SelectAuthors(offset int, limit int) ([]string, error) {
	var authors []string

	if err := m.db.Select("distinct title").Offset(offset).Limit(limit).Find(&authors).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return authors, nil
}

func (m *Mysql) SelectDynastys(offset int, limit int) ([]string, error) {
	var dynastys []string

	if err := m.db.Select("distinct dynasty").Offset(offset).Limit(limit).Find(&dynastys).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return dynastys, nil
}

func (m *Mysql) SelectAncientByID(id int) (*Ancient, error) {
	ancient := &Ancient{}
	if err := m.db.Where("id=?", id).First(ancient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return ancient, nil
}

func (m *Mysql) SelectAncientByAuthor(author string, offset int, limit int) ([]*Ancient, error) {
	var ancients []*Ancient

	if err := m.db.Offset(offset).Limit(limit).Where(&Ancient{Author: author}).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return ancients, nil
}

func (m *Mysql) SelectAncientByDynasty(dynasty string, offset int, limit int) ([]*Ancient, error) {
	var ancients []*Ancient

	if err := m.db.Offset(offset).Limit(limit).Where(&Ancient{Dynasty: dynasty}).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return ancients, nil
}

func (m *Mysql) SelectAncientByTitleAndAuthor(ancient *Ancient, offset int, limit int) ([]*Ancient, error) {
	var ancients []*Ancient

	if err := m.db.Offset(offset).Limit(limit).Where(ancient).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return ancients, nil
}
