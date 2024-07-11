package web

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Elment, Value string
	Font          []string
}

var output Data

func Print(w http.ResponseWriter, r *http.Request) { // w for send data from server to user and r for take data from user
	output.NameFont()
	if output.Font == nil {
		http.Error(w, "server down", http.StatusInternalServerError) // hundul if was file html not
		return
	}
	tmp, err := template.ParseFiles("./templet/index.html") // pionter in file html
	if r.URL.Path != "/" {                                  // handel if url was not valide
		http.NotFound(w, r) // enter to func for print not found
		return
	}

	if err != nil {
		http.Error(w, "server down", http.StatusInternalServerError) // hundul if was file html not
		return
	}
	tmp = template.Must(tmp, err)
	err = (tmp.Execute(w, output)) // f send data to showing in page in form respons and creat ascii art
	if err != nil {
		http.Error(w, "server down", http.StatusAccepted)
	}
}

func Handel_input(w http.ResponseWriter, r *http.Request) {
	font := r.FormValue("select")
	user_input := r.FormValue("user_input")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	if len(font) == 0 || len(user_input) == 0 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if len(user_input) > 1000 {
		user_input = user_input[:1001]
	}
	mapDraw, ErrMsg, ErrStatus := Font(font)
	if mapDraw == nil {
		http.Error(w, ErrMsg, ErrStatus)
		return
	}
	output.Stock(user_input, mapDraw)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (r *Data) Stock(s string, mapDraw map[int][]string) {
	r.Value = s
	r.Elment = SplitAndPrint(s, mapDraw)
}

func (r *Data) NameFont() {
	dir, _ := os.Open("func/draw")
	tr, _ := dir.Readdirnames(-1)
	output.Font = nil
	for _, t := range tr {
		if strings.HasSuffix(t, ".txt") {
			r.Font = append(r.Font, t[:len(t)-4])
		}
	}
}
