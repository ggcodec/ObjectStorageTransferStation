package store

type Uploader interface {
	// 上传文件到后端存储
	UploadFile(bucketName, objectKey, localFilePath string) error
}
