package internal

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/https-dre/supabase-go/models"
)

func UploadFile(url, service_role, bucket, filename, mimetype string, content []byte) models.StorageStatus {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	partHeader.Set("Content-Type", mimetype)

	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return models.StorageStatus{
			Err: err,
		}
	}

	_, err = part.Write(content)
	if err != nil {
		return models.StorageStatus{
			Err: err,
		}
	}

	writer.Close()

	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", url, bucket, filename)
	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		log.Println("Error creating request:", err)
		return models.StorageStatus{
			Err: err,
		}
	}

	req.Header.Set("Authorization", "Bearer "+service_role)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return models.StorageStatus{
			Status: resp.Status,
			Err: err,
		}
	}
	
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		msg := fmt.Errorf("error uploading file: %s", resp.Status)

		return models.StorageStatus{
			Status: resp.Status,
			Err: msg,
			ResponseBody: string(body),
			BucketName: bucket,
		}
	}

	return models.StorageStatus{
		BucketName: bucket,
		Status:     resp.Status,
		Err:        nil,
		ResponseBody: string(body),
	}
}