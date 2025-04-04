package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/https-dre/supabase-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANO_KEY")
	supabaseServiceRole := os.Getenv("SUPABASE_SERVICE_ROLE")

	supabaseClient := supabase_go.NewSupabaseCtx(supabaseURL, supabaseKey, supabaseServiceRole)
	bucket := "docs-system"
	filename := "example2.txt"
	file := []byte("This is an example file.")
	mimetype := "text/plain"
	err = supabaseClient.UploadFile(bucket, filename, file, mimetype)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("File uploaded successfully")
	}
}

