package controller

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
)

func ToBase64 (fileName string) (string) {
	byteArray, _ := ioutil.ReadFile(fileName)
	return base64.StdEncoding.EncodeToString(byteArray)
}

func ToBase64FromFile (file multipart.File) (string) {

	byteArray := bytes.NewBuffer(nil)
	if _, err := io.Copy(byteArray, file); err != nil {
		log.Println(err)
	}

	return base64.StdEncoding.EncodeToString(byteArray.Bytes())
}
