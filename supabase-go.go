package supabase_go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
)

type supabaseContext struct {
	url string
	key string
	service_role string
}

func NewSupabaseCtx(url string, key string, role_key string) *supabaseContext {
	return &supabaseContext {
		url: url,
		key: key,
		service_role: role_key,
	}
}

func (s *supabaseContext) UploadFile(bucket string, filename string, content []byte, mimetype string) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Define headers da parte com o content-type do arquivo
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

	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.url, bucket, filename)
	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+s.service_role)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Println("Response body:", string(body))
		return fmt.Errorf("error uploading file: %s", resp.Status)
	}

	dump, _ := httputil.DumpRequest(req, true)
	log.Println("Request:", string(dump))

	return nil
}