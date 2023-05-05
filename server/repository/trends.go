package repository

import (
	"context"
	"next-social/server/model"
)

var TrendsRepository = new(trendsRepository)

type trendsRepository struct {
	baseRepository
}

func (r trendsRepository) FindTrends(c context.Context, userId []string, pageIndex, pageSize int) (o []model.Trends, err error) {
	err = r.GetDB(c).Select("id, user_id, created, content").
		Where("user_id in (?)", userId).Order("id desc").
		Find(&o).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Error
	return
}

func (r trendsRepository) FindTrendsById(c context.Context, trends_id int) (exist bool, err error) {
	trends := model.Trends{}
	var count uint64
	err = r.GetDB(c).Table(trends.TableName()).Select("count(*)").
		Where("id = ?", trends_id).Find(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r trendsRepository) FindComment(c context.Context, trend_ids []int) (o []model.TrendsComment, err error) {
	err = r.GetDB(c).Where("trends_id in (?)", trend_ids).Order("id desc").Find(&o).Error
	return
}

func (r trendsRepository) FindLikes(c context.Context, trend_ids []int) (o []model.TrendsLikes, err error) {
	err = r.GetDB(c).Where("trends_id in (?)", trend_ids).Order("id desc").Find(&o).Error
	return
}

func (r trendsRepository) Create(c context.Context, o *model.Trends) (err error) {
	err = r.GetDB(c).Create(&o).Error
	return
}

func (r trendsRepository) GetFrid(c context.Context, trends_id int) (trends model.Trends, err error) {
	err = r.GetDB(c).Select("user_id").Where("id = ?", trends_id).First(&trends).Error
	return
}

func (r trendsRepository) CreateComment(c context.Context, o *model.TrendsComment) (err error) {
	err = r.GetDB(c).Create(&o).Error
	return
}
func (r trendsRepository) Linkes(c context.Context, o *model.TrendsLikes, like_type string) (err error) {
	if like_type == "insert" {
		err = r.GetDB(c).Create(&o).Error
	} else {
		err = r.GetDB(c).Where(&o).Delete(&model.TrendsLikes{}).Error
	}
	return
}

func (r trendsRepository) DeleteById(c context.Context, del_type, userId string, id int) (err error) {
	if del_type == "trend" {
		err = r.GetDB(c).Where("id = ? and user_id = ?", id, userId).Delete(&model.Trends{}).Error
	} else {
		err = r.GetDB(c).Where("id = ? and user_id = ?", id, userId).Delete(&model.TrendsComment{}).Error
	}
	return
}
