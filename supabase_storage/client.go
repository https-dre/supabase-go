package storage

import (
	"github.com/https-dre/supabase-go/internal"
	"github.com/https-dre/supabase-go/models"
)

type supabaseCtx struct {
	url          string
	key          string
	service_role string
}

func (s *supabaseCtx) UploadFile(bucket, filename string, content []byte, mimetype string) models.StorageStatus {
	return internal.UploadFile(s.url, s.service_role, bucket, filename, mimetype, content)
}