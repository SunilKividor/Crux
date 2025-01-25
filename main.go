package main

func main() {

	//download video from s3-1
	// s3Client := s3.New(sess)
	// key := event.Records[0].S3Events.Object.Key
	// bucket := event.Records[0].S3Events.Bucket.Name
	// err = downloadFromS3(s3Client, bucket, key)
	// if err != nil {
	// 	log.Printf("Error downloading from s3: %s", err.Error())
	// 	continue
	// }

	//ffmpeg video transcoding
	// cmd := exec.Command("ffmpeg", "-i", key, "-b:v", "13000k", "13000-3.mp4")
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }

	//delete file from s3-1
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

	//upload file to s3-2
	// file, err := os.Open("13000-3.mp4") //name from ffmpeg command
	// if err != nil {
	// 	log.Fatalf("could not open the transcoded file to upload")
	// }
	// putObjectInput := &s3.PutObjectInput{
	// 	Bucket: &trancodedVideosBucket,
	// 	Key:    aws.String("transcoded-video.mp4"),
	// 	Body:   file,
	// }
	// _, err = s3Client.PutObject(putObjectInput)
	// if err != nil {
	// 	log.Println("error uploading file to bucket")
	// }

	// func downloadFromS3(s3Client *s3.S3, bucket, key string) error {
	// 	file, err := os.Create(key)
	// 	if err != nil {
	// 		log.Printf("error creating file : %s", err.Error())
	// 		return err
	// 	}
	// 	defer file.Close()

	// 	decodedKey, err := url.QueryUnescape(key)
	// 	if err != nil {
	// 		log.Printf("error decording key: %s", err.Error())
	// 		return err
	// 	}

	//	getObjectInput := &s3.GetObjectInput{
	//		Bucket: &bucket,
	//		Key:    &decodedKey,
	//	}
	//
	// objectOutput, err := s3Client.GetObject(getObjectInput)
	//
	//	if err != nil {
	//		log.Printf("could not download the object: %s", err.Error())
	//		return err
	//	}
	//
	// noOfBytes, err := io.Copy(file, objectOutput.Body)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	// fmt.Printf("no. of bytes copied: %v", noOfBytes)
	// return nil
}
