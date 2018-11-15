package service 

import (
	"html/template"
	"net/http"
	"time"
)

type indexBinding struct {
	When string
}

var indexTmpl * template.Template

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req * http.Request){
		w.WriteHeader(200)
		if err := indexTmpl.Execute(w, indexBinding{When: time.Now().Format("2018-01-02 15:15:15")}); err != nil {
			panic(err)
		}
	}
}

type tableBinding struct {
	Name string
	When string
}

var tableTmpl * template.Template

func tableHandler() http.HandlerFunc{
	return func (w http.ResponseWriter, req *http.Request){
		if err := req.ParseForm(); err !=nil {
			panic(err)
		}
		w.Header().Set("Content-type","text/html")
		w.WriteHeader(200)
		err := tableTmpl.Execute(w, tableBinding{
			Name: req.PostForm.Get("name"),
			When: req.PostForm.Get("when"),
		})
		if err != nil{
			panic(err)
		}
	}
}


func init (){
	indexTmpl = template.Must(template.New("index.html").ParseFiles("templates/index.html"))
	tableTmpl =  template.Must(template.New("table.html").ParseFiles("templates/table.html"))
}