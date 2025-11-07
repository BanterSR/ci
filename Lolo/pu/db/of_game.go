package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type OFGame struct {
	UserId    uint32    `gorm:"primaryKey;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	BinData   []byte
}

// GetOFGameByUserId 使用UserId拉取数据 如果不存在就添加
func GetOFGameByUserId(userId uint32) (*OFGame, error) {
	user := &OFGame{}
	tx := db.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		}
	}()
	err := tx.Where("user_id = ?", userId).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &OFGame{
				UserId: userId,
			}
			err = tx.Create(user).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	tx.Commit()

	return user, nil
}

// 更新账号数据
func SaveOFGame(userId uint32, fx func(user *OFGame) bool) error {
	tx := db.Begin()
	info := new(OFGame)
	if tx.Where("user_id = ?", userId).First(info); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	if !fx(info) {
		tx.Rollback()
		return nil
	}
	if tx.Save(info).Error != nil {
		tx.Rollback()
		return tx.Error
	}

	return tx.Commit().Error
}
