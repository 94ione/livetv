package service

import (
	"github.com/94ione/livetv/global"
	"github.com/94ione/livetv/model"
)

func GetAllChannel() (channels []model.Channel, err error) {
	err = global.DB.Find(&channels).Error
	return
}

func SaveChannel(channel model.Channel) error {
	return global.DB.Save(&channel).Error
}

func DeleteChannel(id uint) error {
	return global.DB.Delete(model.Channel{}, "id = ?", id).Error
}

func GetChannel(channelNumber uint) (channel model.Channel, err error) {
	err = global.DB.Where("id = ?", channelNumber).First(&channel, channelNumber).Error
	return
}
