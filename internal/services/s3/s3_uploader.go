package s3_service

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (service *S3Service) UploadToS3(s3Client *s3.S3) error {
	trancodedVideosBucket := os.Getenv("AWS_S3_TRANSCODED_VIDEOS_BUCKET")
	file, err := os.Open("13000.mp4") //name from ffmpeg command
	if err != nil {
		log.Fatalf("could not open the transcoded file to upload")
	}
	putObjectInput := &s3.PutObjectInput{
		Bucket: &trancodedVideosBucket,
		Key:    aws.String("transcoded-video.mp4"),
		Body:   file,
	}
	_, err = s3Client.PutObject(putObjectInput)
	return err
}
