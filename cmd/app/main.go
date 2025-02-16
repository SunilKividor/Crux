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
	utils.FailOnError(err, "Error uploading file to bucket")

	//cleanup (removing the local downloaded video file) //getting error check this
	err = utils.DeleteLocalVidoeFile(key)
	utils.FailOnErrorWithoutPanic(err, "Error deleting local file")
}
