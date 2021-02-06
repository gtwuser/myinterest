package main

import (
	"crypto/md5"
	"encoding/base64"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"

	"fmt"
	"os"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	getwd, err := os.Getwd()
	log.Info().Msgf("pwd %v", getwd)
	// Create S3 service client
	svc := s3.New(sess)

	filename := `temp.txt`
	bucket := `cxdl-snb-landing-usw2`
	//tmpFile, err := ioutil.ReadFile(`temp.txt`)
	//tmpBytes := md5.Sum(tmpFile)
	//encSum := base64.StdEncoding.EncodeToString(tmpBytes[:])
	//log.Info().Msgf(`checksum %v`, encSum)

	file, err := os.Open(`temp.txt`)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}
	file.Seek(0, io.SeekStart) // Fix
	hash := md5.New()
	io.Copy(hash, file)
	sum := hash.Sum(nil)
	encSum := base64.StdEncoding.EncodeToString(sum)
	log.Info().Msgf(`checksum %v`, encSum)

	defer file.Close()
	uploader := s3manager.NewUploader(sess)
	v := uploader.PartSize / 1024
	v = v / 1024 // In MBs
	log.Info().Msgf("Default part size %v MB", v)
	metadata := map[string]*string{
		`md5sum`: aws.String(encSum),
		`demo`:   aws.String(`sample`),
	}
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:       file,
		Bucket:     aws.String(bucket),
		Key:        aws.String(filename),
		ContentMD5: aws.String(encSum),
		Metadata:   metadata,
	})
	//,func(u *s3manager.Uploader) {
	//	u.PartSize = 5 * 1024 * 1024 * 1024 // 5GB part size
	//	u.LeavePartsOnError = true    // Don't delete the parts if the upload fails.
	//}
	//)
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	// Get metadata
	results, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	tag := results.ETag
	log.Info().Msgf("Etag %v", base64.StdEncoding.EncodeToString([]byte(*tag)))
	ssum := results.Metadata[`Md5sum`]
	//ssum := metadata[`md5sum`]
	log.Info().Msgf("md5sum %v", *ssum)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
