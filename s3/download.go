package s3

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader interface {
	Download() (string, error)
}

type downloaderImpl struct {
	manager s3manager.Downloader
	bucket  string
	key     string
	dst     string
}

func NewDownloader(s *session.Session, bucket, key, dst string) Downloader {
	manager := *s3manager.NewDownloader(s)
	return &downloaderImpl{manager, bucket, key, dst}
}

func (d *downloaderImpl) Download() (string, error) {
	file, err := os.Create(d.dst)
	if err != nil {
		return "", fmt.Errorf("failed to create file; %s; %w", d.dst, err)
	}
	defer file.Close()

	numBytes, err := d.manager.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(d.key),
	})
	if err != nil {
		return "", fmt.Errorf("failed to download; %w", err)
	}

	log.Printf("Downloaded %s %d bytes", file.Name(), numBytes)
	return file.Name(), nil
}
