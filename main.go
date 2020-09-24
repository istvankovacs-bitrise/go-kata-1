package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"path"
)

var tempDir = "/temp"

func main() {
	url := "https://bitrise-prod-build-storage.s3.amazonaws.com/builds/df7d2a57d4272143/artifacts/28806521/app-armeabi-v7a-release.apk?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIV2YZWMVCNWNR2HA%2F20200923%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20200923T134448Z&X-Amz-Expires=43200&X-Amz-SignedHeaders=host&X-Amz-Signature=3460bcf69613d79187eebd1dfe73da3b888b58242dd08265bd8b03efc5fa3962"
	fileName := "app-armeabi-v7a-release.apk"

	err := downloadFile(fileName, url)

	if err != nil {
		panic(err)
	}
}

func downloadFile(fileName string, url string) error {
	fmt.Println("Download started...")

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempDir := "temp"
	os.Mkdir(tempDir, 0777)
	
	filePath := path.Join(tempDir, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	pwd, _ := os.Getwd()
	fmt.Println("Download finished: " + path.Join(pwd, filePath))

	return err
}