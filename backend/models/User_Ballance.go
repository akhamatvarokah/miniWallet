package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserBalance struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	User          User      `json:"user"`
	UserId        uint64    `gorm:"not null" json:"userId"`
	Balance       uint64    `gorm:"not null" json:"balance"`
	BalanceAchive uint64    `gorm:"not null" json:"balance_achieve"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *UserBalance) UpdateBalance(db *gorm.DB) (*UserBalance, error) {
	err := db.Debug().Model(&UserBalance{}).Where("user_id = ?", u.UserId).Updates(u).Error
	if err != nil {
		return &UserBalance{}, err
	}

	return u, nil
}

func (u *UserBalance) GetBalance(db *gorm.DB) (*UserBalance, error) {
	err := db.Debug().Model(User{}).Where("user_id = ?", u.UserId).Take(&u).Error

	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return u, nil
}

func (u *UserBalance) CreateBalance(db *gorm.DB) (*UserBalance, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
