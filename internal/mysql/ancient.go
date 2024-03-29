package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Ancient struct {
	ID      int    `gorm:"type:bigint(20);primary_key" json:"id"`
	Title   string `gorm:"type:varchar(64);index:title_idx;not null" json:"title,omitempty"`
	Author  string `gorm:"type:varchar(64);index:author_idx;not null" json:"author,omitempty"`
	Dynasty string `gorm:"type:varchar(32);index:dynasty_idx;not null" json:"dynasty,omitempty"`
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

	if err := m.db.Table("ancients").Offset(offset).Limit(limit).Pluck("DISTINCT author", &authors).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	fmt.Println(authors)

	return authors, nil
}

func (m *Mysql) SelectDynastys(offset int, limit int) ([]string, error) {
	var dynastys []string

	if err := m.db.Table("ancients").Offset(offset).Limit(limit).Pluck("DISTINCT dynasty", &dynastys).Error; err != nil {
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

	if err := m.db.Offset(offset).Limit(limit).Select("id, title, author, dynasty").Order("id").Where(&Ancient{Author: author}).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(ancients) == 0 {
		return nil, nil
	}

	return ancients, nil
}

func (m *Mysql) SelectAncientByDynasty(dynasty string, offset int, limit int) ([]*Ancient, error) {
	var ancients []*Ancient

	if err := m.db.Offset(offset).Limit(limit).Select("id, title, author, dynasty").Order("id").Where(&Ancient{Dynasty: dynasty}).Find(&ancients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(ancients) == 0 {
		return nil, nil
	}

	return ancients, nil
}
