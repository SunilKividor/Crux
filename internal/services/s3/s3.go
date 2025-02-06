package s3_service

import "github.com/aws/aws-sdk-go/aws/session"

type S3Service struct {
	Session *session.Session
}

func NewS3Service(session *session.Session) S3Service {
	return S3Service{
		Session: session,
	}
}
