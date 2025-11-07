package db

import (
	"errors"

	"gorm.io/gorm"
)

type OFUser struct {
	UserId    uint32 `gorm:"primarykey;autoIncrement"`
	SdkUid    uint32 `gorm:"unique"`
	Token     string
	DeviceId  string // 设备码
	ChannelId string
	Game      *OFGame `gorm:"foreignKey:UserId"`
}

// GetOFUserByUserId 使用UserId拉取数据
func GetOFUserByUserId(userId uint32) (*OFUser, error) {
	user := &OFUser{UserId: userId}
	tx := db.Where("user_id = ?", userId).First(user)
	return user, tx.Error
}

// GetOFUserBySdkUid 使用SdkUid拉取数据
func GetOFUserBySdkUid(sdkUid uint32) (*OFUser, error) {
	user := &OFUser{}
	tx := db.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		}
	}()
	err := tx.Where("sdk_uid = ?", sdkUid).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &OFUser{
				SdkUid: sdkUid,
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
	return user, tx.Error
}

func SaveOFUser(sdkUid uint32, fx func(user *OFUser) bool) error {
	tx := db.Begin()
	info := new(OFUser)
	if tx.Where("sdk_uid = ?", sdkUid).First(info); tx.Error != nil {
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
