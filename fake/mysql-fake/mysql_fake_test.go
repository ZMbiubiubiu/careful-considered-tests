package main

import (
	"context"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// 启动内存mysql server
	go ServerRun()
	time.Sleep(2 * time.Second)

	// 创建client
	var mysql interface{} = "root:@tcp(localhost:3307)/chaos?charset=utf8mb4&parseTime=false&loc=Local"
	_ = orm.RegisterDataBase("default", "mysql", mysql.(string), 30, 30)
	orm.Debug = true
	//数据库中的datetime使用local时间
	orm.DefaultTimeLoc = time.Local

	os.Exit(m.Run())
}

func TestInsertAndGet(t *testing.T) {
	var (
		err   error
		expId = int64(666)
		scene = 2
		ctx   = context.TODO()
	)

	grayscaleModel := &GrayscaleSetModel{
		Scene:             scene,
		CollectionId:      expId,
		Status:            1,
		WaitTime:          70,
		FailContinueFault: 1,
	}

	// 保存&获取批量编排的灰度配置
	_, err = NewGrayscaleSet(ctx, grayscaleModel)
	assert.Nil(t, err)

	got, err := GetGrayscaleSet(ctx, scene, expId)
	assert.Nil(t, err)
	assert.Equal(t, grayscaleModel.Scene, got.Scene)
	assert.Equal(t, grayscaleModel.CollectionId, got.CollectionId)
	assert.Equal(t, grayscaleModel.Status, got.Status)
	assert.Equal(t, grayscaleModel.WaitTime, got.WaitTime)
	assert.Equal(t, grayscaleModel.FailContinueFault, got.FailContinueFault)
}
