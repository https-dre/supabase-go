package models

type StorageStatus struct {
	BucketName string `json:"bucket_name"`
	Status string `json:"status"`
	Err error `json:"err"`
	ResponseBody string `json:"response_body"`
}