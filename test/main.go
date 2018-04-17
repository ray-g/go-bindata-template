package main

import (
	"log"
	"net/http"

	template "github.com/ray-g/go-bindata-template"
	bin "github.com/ray-g/go-bindata-template/test/bindata"
)

type indexData struct {
	Target  string
	Content string
}

func main() {
	http.HandleFunc("/files", filesHandler)
	http.HandleFunc("/dir", dirHandler)
	http.HandleFunc("/all", allHandler)
	http.HandleFunc("/", singleHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

func singleHandler(w http.ResponseWriter, r *http.Request) {
	if err := singleRender(w); err != nil {
		log.Println(err)
	}
}

func singleRender(w http.ResponseWriter) error {
	idxData := indexData{
		Target:  "Universe Single File",
		Content: "Earth Single File",
	}

	data := &template.BinData{
		Asset:      bin.Asset,
		AssetDir:   bin.AssetDir,
		AssetNames: bin.AssetNames,
	}

	index, err := template.New("single", data).Parse("templates/single.tmpl")

	if err == nil {
		err = index.Execute(w, idxData)
	}

	return err
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	if err := filesRender(w); err != nil {
		log.Println(err)
	}
}

func filesRender(w http.ResponseWriter) error {
	idxData := indexData{
		Target:  "Universe Files",
		Content: "Earth Files",
	}

	indexFileSet := []string{
		"templates/index.tmpl",
		"templates/content.tmpl",
	}

	data := &template.BinData{
		Asset:      bin.Asset,
		AssetDir:   bin.AssetDir,
		AssetNames: bin.AssetNames,
	}

	index, err := template.New("index", data).ParseFiles(indexFileSet...)

	if err == nil {
		err = index.Execute(w, idxData)
	}

	return err
}

func dirHandler(w http.ResponseWriter, r *http.Request) {
	if err := dirRender(w); err != nil {
		log.Println(err)
	}
}

func dirRender(w http.ResponseWriter) error {
	idxData := indexData{
		Target:  "Universe Dir",
		Content: "Earth Dir",
	}

	data := &template.BinData{
		Asset:      bin.Asset,
		AssetDir:   bin.AssetDir,
		AssetNames: bin.AssetNames,
	}

	index, err := template.New("index", data).ParseDir("templates")

	if err == nil {
		err = index.Execute(w, idxData)
	}

	return err
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	if err := allRender(w); err != nil {
		log.Println(err)
	}
}

func allRender(w http.ResponseWriter) error {
	idxData := indexData{
		Target:  "Universe All",
		Content: "Earth All",
	}

	data := &template.BinData{
		Asset:      bin.Asset,
		AssetDir:   bin.AssetDir,
		AssetNames: bin.AssetNames,
	}

	index, err := template.New("index", data).ParseAll()

	if err == nil {
		err = index.Execute(w, idxData)
	}

	return err
}
