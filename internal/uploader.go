package internal

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

func UploadFile(url, service_role, bucket, filename, mimetype string, content []byte) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	partHeader.Set("Content-Type", mimetype)

	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return err
	}

	_, err = part.Write(content)
	if err != nil {
		return err
	}

	writer.Close()

	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", url, bucket, filename)
	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+service_role)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Println("Response body:", string(body))
		return fmt.Errorf("error uploading file: %s", resp.Status)
	}

	return nil
}