package database

import (
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

func NewStorage() *storage_go.Client {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	storageClient := storage_go.NewClient(supabaseUrl, supabaseKey, nil)

	return storageClient
}
