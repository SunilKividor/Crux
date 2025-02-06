package s3_service

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
)

func (service *S3Service) DownloadFromS3(s3Client *s3.S3, bucket, key string) error {
	file, err := os.Create(key)
	if err != nil {
		log.Printf("error creating file : %s", err.Error())
		return err
	}
	defer file.Close()

	decodedKey, err := url.QueryUnescape(key)
	if err != nil {
		log.Printf("error decording key: %s", err.Error())
		return err
	}

	getObjectInput := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &decodedKey,
	}

	objectOutput, err := s3Client.GetObject(getObjectInput)

	if err != nil {
		log.Printf("could not download the object: %s", err.Error())
		return err
	}

	noOfBytes, err := io.Copy(file, objectOutput.Body)

	if err != nil {
		return err
	}

	fmt.Printf("no. of bytes copied: %v", noOfBytes)
	return nil
}
