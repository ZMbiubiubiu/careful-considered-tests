package main

import (
	"context"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

var (
	TableGrayscaleSetModel = new(GrayscaleSetModel)
)

type GrayscaleSetModel struct {
	Id                int64     `orm:"auto;pk"`
	Scene             int       // 业务场景
	CollectionId      int64     // 业务id
	Status            int8      // 是否进行灰度：1进行 0不进行
	Percent           string    //
	WaitTime          int       //
	FailContinueFault int8      //
	Modifier          string    // 更新人
	ModifyTime        time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(TableGrayscaleSetModel)
}
func (r *GrayscaleSetModel) TableName() string {
	return "grayscale_set"
}

func NewGrayscaleSet(ctx context.Context, setting *GrayscaleSetModel) (id int64, err error) {
	id, err = orm.NewOrm().Insert(setting)
	if err != nil {
		log.Fatalf("insert service_plan fail||grayscaleSetting=%v||err=%v", setting, err)
	}
	return
}

func GetGrayscaleSet(ctx context.Context, scene int, collectionId int64) (setting *GrayscaleSetModel, err error) {
	setting = &GrayscaleSetModel{}
	err = orm.NewOrm().QueryTable(TableGrayscaleSetModel).
		Filter("scene", scene).Filter("collection_id", collectionId).One(setting)
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}
	return setting, nil
}
