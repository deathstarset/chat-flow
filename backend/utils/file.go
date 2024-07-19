package utils

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GenRandString(length int) string {
	const charset = "abcdefjhijklmnopqrstuvwxyzABCDEFJHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func SaveFile(file *multipart.FileHeader, c *gin.Context) (string, error) {
	ext := filepath.Ext(file.Filename)
	imageName := fmt.Sprintf("%s%s", GenRandString(10), ext)
	dst := filepath.Join("uploads", "profile", imageName)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		return imageName, err
	}
	return imageName, nil
}
