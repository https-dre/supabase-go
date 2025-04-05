package storage

import "github.com/https-dre/supabase-go/models"

type StorageClient interface {
	UploadFile(bucket, filename string, content []byte, mimetype string) models.StorageStatus
}

func NewClient(url, key, service_role string) StorageClient {
	return &supabaseCtx{
		url:          url,
		key:          key,
		service_role: service_role,
	}
}