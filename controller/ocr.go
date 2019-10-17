package controller

import (
	"encoding/json"
	"errors"
)

type OcrRequestImg struct {
	ImgBase64  string `json:"img_base64"`
	Engine     string `json:"engine"`
	EngineArgs struct {
		Lang string `json:"lang"`
	} `json:"engine_args"`
}

type OcrRequestUrl struct {
	ImgUrl  string `json:"img_url"`
	Engine     string `json:"engine"`
	EngineArgs struct {
		Lang string `json:"lang"`
	} `json:"engine_args"`
}

func Get(inputType string, image string, lang string) (string) {

	jsonData, err := getRequest(inputType, image, lang)

	if err != nil {
		return err.Error()
	}

	return BytesToString(SendPostRequest(jsonData))
}

func GetOcr(request *OcrRequest) (string, error) {

	var err error
	var jsonData []byte
	if request.InputType == "url" {
		jsonData, err = getUrlRequest(request.FileUrl, request.Lang)
	} else if request.InputType == "img" {
		jsonData, err = getImgRequest(request.File64, request.Lang)
	} else {
		err = errors.New("Unknowen iput type : " + request.InputType)
	}

	if err != nil {
		return "", err
	}

	return BytesToString(SendPostRequest(jsonData)), err
}

func getRequest(inputType string, filePathUrl string, lang string) ([]byte, error) {

	var err error
	var jsonData []byte
	if inputType == "url" {
		 jsonData, err = getUrlRequest(filePathUrl, lang)
	} else if inputType == "img" {
		jsonData, err = getImgRequest(filePathUrl, lang)
	} else {
		err := errors.New("Unknowen iput type : " + inputType)
		return []byte{}, err
	}

	return jsonData, err
}

func getUrlRequest(fileUrl string, lang string) ([]byte, error)  {
	var requestData OcrRequestUrl

	requestData.ImgUrl = fileUrl
	requestData.Engine = "tesseract"
	requestData.EngineArgs.Lang = lang

	jsonData, err := json.MarshalIndent(requestData, "", "	")
	if err != nil {
		return []byte{}, err
	}

	return jsonData, nil
}

func getImgRequest(image string, lang string) ([]byte, error)  {
	var requestData OcrRequestImg

	requestData.ImgBase64 = image
	requestData.Engine = "tesseract"
	requestData.EngineArgs.Lang = lang

	jsonData, err := json.MarshalIndent(requestData, "", "	")
	if err != nil {
		return []byte{}, err
	}

	return jsonData, nil
}

func BytesToString(data []byte) string {
	return string(data[:])
}
