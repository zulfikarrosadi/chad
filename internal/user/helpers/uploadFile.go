package helpers

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func Upload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	formatFile, b := isImage(src)
	if !b {
		return "", errors.New("uploaded file is not supported")
	}
	defer src.Close()

	fileName := generateFileName(formatFile)
	dst, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}

func isImage(file multipart.File) (string, bool) {
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", false
	}
	return format, format == "jpg" || format == "jpeg" || format == "png"
}

func generateFileName(fileType string) string {
	timestamp := strconv.Itoa(time.Now().Minute())
	randomNumber := strconv.Itoa(rand.Intn(1000))
	return timestamp + randomNumber + "." + fileType
}
