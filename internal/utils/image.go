package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

var supabaseClient *storage_go.Client

func NewStorage() {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	supabaseClient = storage_go.NewClient(supabaseUrl, supabaseKey, nil)

}

func UploadImage(image multipart.File, path string) error {

	var response storage_go.FileUploadResponse
	fileBuffer := new(bytes.Buffer)
	_, err := io.Copy(fileBuffer, image)

	if err != nil {
		return fmt.Errorf("failed to read image data: %v", err)
	}

	response, err = supabaseClient.UploadFile("Connect Verse", path, image)

	fmt.Println(err)

	if err != nil {
		return fmt.Errorf(err.Error() + "tes")
	}

	fmt.Println(response)

	return nil
}

func GenerateFilePath(file multipart.FileHeader, folder string) string {
	date := time.Now().Unix()

	fileExtension := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	return fmt.Sprintf("%s/%d%d.%s", folder, date, file.Size, fileExtension)
}
