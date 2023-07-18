package ascart

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func Control(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Error! Not a Post method!")
		Error(w, http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
		log.Println("Error! Incorrect Form Parse!")
	}
	if !(r.Form.Has("text")) {
		Error(w, http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if len(text) > 400 {
		Error(w, http.StatusBadRequest)
		return
	}
	if CheckSymbol(text) {
		log.Println("Error! Incorrect symbol!")
		Error(w, http.StatusBadRequest)
		return
	}
	format := r.FormValue("format")
	if format == "" {
		format = "standard"
	}
	if checkFormat(format) {
		Error(w, http.StatusInternalServerError)
		return
	}
	format = "./arts/" + format + ".txt"
	arrByte, err := ioutil.ReadFile(format)
	if err != nil {
		log.Println("Error! Incorrect read file from /arts/.txt")
		Error(w, http.StatusBadRequest)
		return
	}
	hash := HashMD5(string(arrByte))
	if !(CheckHash(hash)) {
		log.Println("Error! Incorrect hash! Don`t change .txt files!")
		Error(w, http.StatusInternalServerError)
		return
	}
	Result := ReadFile(text, format)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, Result)
}
