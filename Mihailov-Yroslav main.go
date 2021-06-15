package main


import (
	'encoding/json'
	'fmt'
	"net/http"
)

const content = "Content-Type"

func main() {
	http.HandleFunc("/first", homeHandle)

	http.HandleFunc("/ff", ffHandle)

	http.HandleFunc("/finish", finishHandle)

	http.ListenAndServe(":80", nil)
}
func homeHandle(w http.ResponseWriter, r *http.Request) {
	r.Header.Set(content, "application/json")

	switch r.Method {
	case "GET":
		p := struct {
			Value string `json:"value"`
			Age int 	`json:"age"`
			Name string `json:"name"`
		}{
			"That is a get method",
			19,
			"Mihailov Yroslav",
		}

		json.NewEncoder(w).Encode(p)
	case "POST":
		bb, _ := json.Marshal(`
			"hello": "that is a post method"
			`)
		fmt.Fprintf(w, string(bb))
	}}

func ffHandle(w http.ResponseWriter, r *http.Request) {
	r.Header.Set(content, "application/json")

	switch r.Method {
	case "GET":
		p := struct {
			Message  string `json:"message"`
			Value  string `json:"value"`
		}{
			"ff handler message",
			"MediaSoft",
		}

		json.NewEncoder(w).Encode(p)
	case "POST":
		bb, _ := json.Marshal(`
			"Hello": "that is a post method (=)"
			`)
		fmt.Fprintf(w, string(bb))
	}}


func finishHandle(w http.ResponseWriter, r *http.Request) {
	r.Header.Set(content, "application/json")

	switch r.Method {
	case "GET":
		p := struct {
			Message  string `json:"message"`
			Value  string `json:"value"`
		}{
			"Finish handler message",
			"PDO 34-18",
		}

		json.NewEncoder(w).Encode(p)
	case "POST":
		bb, _ := json.Marshal(`
			"Hello": "that is a post method=)"
			`)
		fmt.Fprintf(w, string(bb))
	}}
