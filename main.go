package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

// type Data struct {
// 	Function string `json:"function"`
// 	Request  string `json:"request"`
// }
type Data struct {
	DefinitionRepository    string `json:"DefinitionRepository"`
	DefinitionServiceserver string `json:"DefinitionServiceserver"`
	DefinitionRPC           string `json:"DefinitionRPC"`
	UsecaseFunction         string `json:"UsecaseFunction"`
}

func main() {

	inputDir := "inputs/"
	inputFile := "sample_server.json"
	inputPath := inputDir + inputFile

	jsonBody, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	if err := json.Unmarshal(jsonBody, &data); err != nil {
		log.Fatal(err)
	}

	templateDir := "templates/"
	templateFile := "server.tmpl"
	templatePath := templateDir + templateFile

	t, err := template.New(templateFile).ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	outputDir := "outputs/hoge/"
	outputFile := "hoge.go"
	outputPath := outputDir + outputFile
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = t.Execute(f, data); err != nil {
		log.Fatal(err)
	}
}
