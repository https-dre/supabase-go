package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/https-dre/supabase-go/supabase_storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANO_KEY")
	supabaseServiceRole := os.Getenv("SUPABASE_SERVICE_ROLE")

	supabaseClient := storage.NewClient(supabaseURL, supabaseKey, supabaseServiceRole)

	bucket := "docs-system"
	filename := "example5.txt"
	file := []byte("This is an example file.")
	mimetype := "text/plain"

	resp := supabaseClient.UploadFile(bucket, filename, file, mimetype)

	log.Println("Supabase Status:", resp.Status)
	log.Println("Supabase Error:", resp.Err)
	log.Println("Supabase Bucket Name:", resp.BucketName)
	log.Println("Supabase Response Body:", resp.ResponseBody)
}

