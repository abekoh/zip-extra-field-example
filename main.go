package main

import (
	"archive/zip"
	"log"
	"os"
	"time"
)

var JST *time.Location

func init() {
	os.Setenv("TZ", "UTC")
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	JST = jst
}

func main() {
	saveAsZip("/tmp/default.zip", "default.txt", "default", false)
	saveAsZip("/tmp/in_jst.zip", "in_jst.txt", "in_jst", true)
}

func saveAsZip(path, txtFileName, txtContent string, inJST bool) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zipWriter := zip.NewWriter(f)
	defer zipWriter.Close()

	modified := func() time.Time {
		if inJST {
			return time.Now().In(JST)
		} else {
			return time.Now()
		}
	}()
	zipFile, err := zipWriter.CreateHeader(&zip.FileHeader{
		Name:     txtFileName,
		Modified: modified,
	})
	_, err = zipFile.Write([]byte(txtContent))
	if err != nil {
		log.Fatal(err)
	}
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
