package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// 画像のPath
var filePath = "image/sakura.jpeg"

// S3のバケット名
var bucket = "test-bucket-0814"

// key S3に保存するオブジェクトの名前になります
var key = "image/sakura"

// awsのリージョン名
var awsRegion = "ap-northeast-1"

func main() {
	putS3Object()
}

func putS3Object() {

	// sessionを作成します
	s := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// 画像を読み込みます
	imageFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// 最後に画像ファイルを閉じます
	defer imageFile.Close()

	// Uploaderを作成し、S3にアップロードします
	uploader := s3manager.NewUploader(s)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   imageFile,
	})
	// エラーハンドリング
	if err != nil {
		log.Fatal(err)
	}
	log.Println("S3へアップロードが完了しました。")
}
