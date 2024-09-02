package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GenerateFileName(file *multipart.FileHeader) string {
	date := time.Now().Unix()

	fileExtension := strings.TrimPrefix(filepath.Ext(file.Filename), ".")

	return fmt.Sprintf("%d%d.%s", date, file.Size, fileExtension)
}

func UploadFileToS3(s3Client *s3.Client, file multipart.File, fileName string, bucketName string, folderName string) (string, error) {

	fullPath := fmt.Sprintf("%s/%s", folderName, fileName)

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fullPath,
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.ap-southeast-1.amazonaws.com/%s", bucketName, fullPath), nil

}

func UpdateFileInS3(s3Client *s3.Client, file multipart.File, path string, bucketName string) (string, error) {

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &path,
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return path, nil

}
