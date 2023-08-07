package sqlstore

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client interface {
	PutObject(bucket, key string, file io.Reader) error
	GetObject(bucket, key string) (io.Reader, error)
	RemoveObject(bucket, key string) error
}

type S3Service struct {
	client S3Client
}

func (s3s *S3Service) UploadFile(bucket, key string, file io.Reader) error {
	return s3s.client.PutObject(bucket, key, file)
}

func (s3s *S3Service) DownloadFile(bucket, key string) (io.Reader, error) {
	return s3s.client.GetObject(bucket, key)
}

func (s3s *S3Service) DeleteFile(bucket, key string) error {
	return s3s.client.RemoveObject(bucket, key)
}

type AWSS3Client struct {
	svc *s3.S3
}

func NewAWSS3Client() *AWSS3Client {
	sess := session.Must(session.NewSession())
	return &AWSS3Client{
		svc: s3.New(sess),
	}
}

func (a *AWSS3Client) PutObject(bucket, key string, file io.Reader) error {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(content)

	_, err = a.svc.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   reader,
	})
	return err
}

func (a *AWSS3Client) GetObject(bucket, key string) (io.Reader, error) {
	result, err := a.svc.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func (a *AWSS3Client) RemoveObject(bucket, key string) error {
	_, err := a.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	return err
}
