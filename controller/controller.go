package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"ocrClient/view"
)
type OcrRequest struct {
	InputType string		`json:inpuType`
	File64 string			`json:file64`
	FileUrl string			`json:fileUrl`
	Lang string				`json:lang`
}

type OcrResponse struct {
	Text string	`json:"text"`
	Message string	`json:"message"`
}


func GetOcrUi(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "home"
	v.Render(w)
}

func GetOcrApi(w http.ResponseWriter, r *http.Request) {
	request, err := getRequestData(r)
	if err != nil {
			response := OcrResponse{
			Message: err.Error(),
		}
		SendJSON(w, response)
		return
	}

	ocrString, err := GetOcr(request)
	if err != nil {
		response := OcrResponse{
			Message: err.Error(),
		}
		SendJSON(w, response)
		return
	}

	fmt.Println(ocrString)

	response := OcrResponse{
		Text: ocrString,
	}

	SendJSON(w, response)
}

func getRequestData(r *http.Request) (*OcrRequest, error)  {

	var request OcrRequest

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, err
	}

	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

func SendJSON(w http.ResponseWriter, i interface{}) {
	js, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
