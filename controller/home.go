package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/josephspurrier/csrfbanana"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"ocrClient/session"
	"ocrClient/view"
	"strings"
)

var languages = []string{"eng", "ara", "bel", "ben", "bul", "ces", "dan", "deu", "ell", "fin", "fra", "heb", "hin", "ind", "isl", "ita", "jpn", "kor", "nld", "nor", "pol", "por", "ron", "rus", "spa", "swe", "tha", "tur", "ukr", "vie", "chi-sim", "chi-tra"}

// IndexGET displays the home page
func HomeGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "home"
	v.Vars["Languages"] = languages
	v.Vars["Language"] = "eng"
	v.Vars["token"] = csrfbanana.Token(w, r, session)
	v.Render(w)
	view.Repopulate([]string{"ImgURL"}, r.Form, v.Vars)
	return
}

func HomePOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	ocrFile := false

	// Form values
	imgUrl := r.FormValue("ImgURL")
	language := r.FormValue("Language")

	var file multipart.File
	var header * multipart.FileHeader
	var err error
	image := ""

	if imgUrl == "" {
		file, header, err = r.FormFile("ImgPath")
		if err != nil {
			sess.AddFlash(view.Flash{"Field missing: Image URL or Image path", view.FlashError})
			sess.Save(r, w)
			HomeGET(w, r)
			return
		}
		defer file.Close()
		log.Println(header.Filename)
		image = ToBase64FromFile(file)
		ocrFile = true
	} else {
		resp, _ := http.Get(imgUrl)
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil.ReadAll -> %v", err)
		} else {
			reqImage := base64.StdEncoding.EncodeToString(data)
			contentType := http.DetectContentType(data)
			r.Form.Add("ImgSrc", reqImage)
			r.Form.Add("ImgType", contentType)
		}
	}

	var result string

	if ocrFile {
		result = Get("img", image, language)
	} else {
		result = Get("url", imgUrl, language)
	}

	fmt.Println(result)

	r.Form.Add("OcrResult", result)
	r.Form.Add("ImgURL", imgUrl)
	if ocrFile {
		r.Form.Add("ImgPath", header.Filename)
		r.Form.Add("ImgSrc", image)
		r.Form.Add("ImgType", header.Header.Get("Content-Type"))
	}

	r.Form.Add("Language", language)

	HomeGETResult(w, r)
}

// IndexGET displays the home page
func HomeGETResult(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "home"
	v.Vars["Languages"] = languages
	v.Vars["Language"] = r.FormValue("Language")
	v.Vars["ImgURL"] = ""
	v.Vars["ImgPath"] = ""
	imgSrc := r.FormValue("ImgSrc")
	if imgSrc != "" {
		v.Vars["ImgSrc"] = template.URL("data:"+ r.FormValue("ImgType") +";base64," + imgSrc)
	}
	v.Vars["OcrResult"] = (template.HTML(strings.Replace(r.FormValue("OcrResult"), "\n", "<br>", -1)))
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)
}
