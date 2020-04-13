package oss

import (
	"bell/config"
	"bell/library/helper"
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	tim "bell/library/time"
	"time"
)

type AliyunOss struct {
	client *oss.Client
	bucketName string
}

type Option func(*AliyunOss)

func NewAliyunOss(conf *config.Config) (*AliyunOss,error) {
	opt := conf.AliyunOss
	client ,err :=oss.New(opt.Endpoint,opt.AccessId,opt.AccessSecret)
	if err!=nil{
		return nil,err
	}
	return &AliyunOss{client:client},nil
}

func (o *AliyunOss) SetBucketName(s string) Option {
	return func(aliyunOss *AliyunOss) {
		aliyunOss.bucketName = s
	}
}

func (o *AliyunOss) PutObject(key string,data []byte) (string, error)  {
	bucket, err := o.client.Bucket(o.bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(key, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return key,nil
}

func (o *AliyunOss) DeleteObject(objectName string) (string,error) {
	bucket, err := o.client.Bucket(o.bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.DeleteObject(objectName)
	if err!=nil{
		return "",err
	}
	return objectName,nil
}

/*
	上传图片
 */
func (o *AliyunOss)UploadImage(data []byte) (string, error) {
	md5 := helper.MD5Bytes(data)
	key := "images/" + tim.TimeFormat(time.Now(), "2006/01/02/") + md5 + ".jpg"
	return o.PutObject(key, data)
}





