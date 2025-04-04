package storage

type StorageClient interface {
	UploadFile(bucket, filename string, content []byte, mimetype string) error
}

func NewClient(url, key, service_role string) StorageClient {
	return &supabaseCtx{
		url:          url,
		key:          key,
		service_role: service_role,
	}
}