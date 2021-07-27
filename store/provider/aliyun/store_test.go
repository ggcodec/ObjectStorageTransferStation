package aliyun_test

import (
	"testing"

	"gitee.com/infraboard/go-course/day8/cloudstation/store/provider/aliyun"

	"github.com/stretchr/testify/assert"
)

var (
	bucketName    = "cloud-station"
	objectKey     = "store.go"
	localFilePath = "store.go"

	endpoint = "http://oss-cn-chengdu.aliyuncs.com"
	ak       = "LTAI5tMvG5NA51eiH3ENZDaa"
	sk       = "xxx"
)

func TestUploadFile(t *testing.T) {
	should := assert.New(t)

	uploader, err := aliyun.NewUploader(endpoint, ak, sk)
	if should.NoError(err) {
		err = uploader.UploadFile(bucketName, objectKey, localFilePath)
		should.NoError(err)
	}
}
