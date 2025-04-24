package dao

import "github.com/cheng1005/XxCh-Common/common/global"

// SearchById 根据id搜索
func SearchById[T any](id uint, data *T) bool {
	if err := global.DB.Debug().Where("id = ?", id).Find(&data).Error; err != nil {
		return false
	}
	return true
}

// Create 添加
func Create[T any](data *T) bool {
	if err := global.DB.Debug().Create(&data).Error; err != nil {
		return false
	}
	return true
}

// Update 更新
func Update[T any](data *T) bool {
	if err := global.DB.Debug().Updates(&data).Error; err != nil {
		return false
	}
	return true
}
