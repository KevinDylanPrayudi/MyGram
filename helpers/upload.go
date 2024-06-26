package helpers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
	"os"

	imgBB "github.com/JohnNON/ImgBB"
)

var (
	key = os.Getenv("apiKeyImgBB")
)

func Upload(file *multipart.FileHeader) (string, error) {

	dataFile, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(dataFile)
	img, err := imgBB.NewImageFromFile(hashSum(b), 60, b)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	imgBBClient := imgBB.NewClient(httpClient, key)

	resp, err := imgBBClient.Upload(context.Background(), img)
	if err != nil {
		return "", err
	}

	return string(resp.Data.Image.URL), nil
}

func hashSum(b []byte) string {
	sum := md5.Sum(b)

	return hex.EncodeToString(sum[:])
}
