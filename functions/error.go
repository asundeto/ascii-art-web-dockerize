package ascart

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, code int) {
	response := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: http.StatusText(code),
	}
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, response)
}

func CheckSymbol(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 31 || s[i] > 127 {
			if s[i] != 10 && s[i] != 13 {
				return true
			}
		}
	}
	return false
}

func checkFormat(s string) bool {
	if s != "standard" && s != "shadow" && s != "thinkertoy" {
		return true
	}
	return false
}
