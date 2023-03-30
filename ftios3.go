package ftios3

// Modified from source code -->  https://aws.github.io/aws-sdk-go-v2/docs/code-examples/s3/putobject/

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3PutObjectAPI := defines interface for the PutObject function & tests it.
type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// PutFile := uploads a file to an Amazon S3 bucket
func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

// PutFileS3 := uses directory, fileName, and region to verify the user's S3 files for S3
func PutFileS3(dir, filename, bucket, reg string) error {

	var cfg aws.Config

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(reg))

	if err != nil {
		panic("\n\t Configuration Error: " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	file, err2 := os.Open(dir + filename)
	defer file.Close()

	if err2 != nil {
		fmt.Println("\n\t Unable to open file " + filename)
		return err2
	}

	tmp := "backup" + dir + filename

	uri := strings.Replace(tmp, " ", "##,##", -1)

	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    aws.String(uri),
		Body:   file,
	}

	ctx := context.Background()

	_, err2 = PutFile(ctx, client, input)
	if err2 != nil {
		fmt.Println("\n\t Error uploading file: ", dir+filename)
		fmt.Println(err2)
		return err2
	}

	return nil
}
