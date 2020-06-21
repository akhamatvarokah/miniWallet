package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserBalanceHistory struct {
	ID            uint64      `gorm:"primary_key;auto_increment" json:"id"`
	UserBalance   UserBalance `json:"user_balance"`
	UserBalanceId uint64      `gorm:"not null" json:"user_balance_id"`
	BalanceBefore uint64      `gorm:"not null" json:"balance_before"`
	BalanceAfter  uint64      `gorm:"not null" json:"balance_after"`
	Activity      string      `gorm:"not null" json:"activity"`
	Type          string      `gorm:"not null" json:"type"`
	Ip            string      `gorm:"not null" json:"ip"`
	UserAgent     string      `gorm:"not null" json:"user_agent"`
	Author        string      `gorm:"not null" json:"author"`
	CreatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *UserBalanceHistory) Prepare() {
	p.UserBalance = UserBalance{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (u *UserBalanceHistory) CreateBalanceHistory(db *gorm.DB) (*UserBalanceHistory, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserBalanceHistory) FindAllHistoryUser(db *gorm.DB, uid uint64) (*[]UserBalanceHistory, error) {
	var err error
	ubh := []UserBalanceHistory{}
	err = db.Debug().Model(&UserBalanceHistory{}).Where("user_balance_id = ?", uid).Limit(100).Order("created_at desc").Find(&ubh).Error
	if err != nil {
		return nil, err
	}

	if len(ubh) > 0 {
		for i, _ := range ubh {
			err := db.Debug().Model(&UserBalance{}).Where("id = ?", ubh[i].UserBalanceId).Take(&ubh[i].UserBalance).Error
			if err != nil {
				return nil, err
			}

			err = db.Debug().Model(&User{}).Where("id = ?", ubh[i].UserBalance.UserId).Take(&ubh[i].UserBalance.User).Error
			if err != nil {
				return nil, err
			}

		}
	}

	return &ubh, nil
}
