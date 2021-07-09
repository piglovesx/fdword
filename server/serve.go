package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/piglovesx/fdword/client"
)

var IsAlpha = regexp.MustCompile(`^[a-z0-9]+$`).MatchString

func main() {
	http.HandleFunc("/w", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/json", getJson)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func getJson(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query()["word"]
	if IsAlpha(word[0]) {
		words := client.Permute(strings.Split(word[0], ""))
		existedwords := client.FindAllLenWords(words)
		w.Header().Add("Content-Type", "application/json")
		if d, err := json.Marshal(existedwords); err == nil {
			io.WriteString(w, string(d))
		}

	} else {
		io.WriteString(w, "")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[1:]
	word := r.URL.Query()["word"]
	t, _ := template.ParseFiles("wordescape.html")
	if word != nil {
		if IsAlpha(word[0]) {
			words := client.Permute(strings.Split(word[0], ""))
			existedwords := client.FindExistedWord(words)
			t.Execute(w, &existedwords)
		}
	} else {
		t.Execute(w, "Try again")
	}
}
