package main

import (
	"log"
	"os"

	"github.com/SunilKividor/Cruc/internal/config"
	"github.com/SunilKividor/Cruc/internal/services/ffmpeg"
	s3_service "github.com/SunilKividor/Cruc/internal/services/s3"
	"github.com/SunilKividor/Cruc/pkg/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Could not load .env")
	}
}

func main() {
	//aws config
	awsConfigModel := config.NewAwsConfigModel()
	awsConfig := awsConfigModel.NewAwsConfig()
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			Config: awsConfig,
		},
	))

	//download video from s3-1
	s3Client := s3.New(sess)
	bucket := os.Getenv("bucket") //get this from env
	key := os.Getenv("key")       //get this from env
	s3Service := s3_service.NewS3Service(sess)
	err := s3Service.DownloadFromS3(s3Client, bucket, key)
	utils.FailOnError(err, "Error downloading from s3")

	// ffmpeg video transcoding
	err = ffmpeg.VideoTranscoder(key)
	utils.FailOnError(err, "Error transcoding video")

	// upload file to s3-2
	err = s3Service.UploadToS3(s3Client)
	utils.FailOnError(err, "error uploading file to bucket")

	//sqs
	// sqsClient := sqs.New(sess)
	// sqsDeleteMsgInput := sqs.DeleteMessageInput{
	// 	QueueUrl:      &sqsQueryUrl,
	// 	ReceiptHandle: aws.String(""),
	// }
	// // delete file from s3-1
	// decodedKey, err := url.QueryUnescape(key)
	// if err != nil {
	// 	log.Fatalf("error decording key: %s", err.Error())
	// }
	// deleteObjectInput := &s3.DeleteObjectInput{
	// 	Bucket: &bucket,
	// 	Key:    &decodedKey,
	// }
	// _, err = s3Client.DeleteObject(deleteObjectInput)
	// if err != nil {
	// 	log.Println("Could not delete the object from s3-2")
	// }
}
