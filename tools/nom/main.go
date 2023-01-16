package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Step struct {
	Step  string   `yaml:"step"`
	Notes []string `yaml:"notes"`
}

type Amount struct {
	Amount float64 `yaml:"amount"`
	Unit   string  `yaml:"unit"`
}

type Ingredient struct {
	Amounts []map[string]Amount `yaml:"amount"`
	Notes   []string            `yaml:"notes"`
}

type Recipe struct {
	RecipeName  string                  `yaml:"recipe_name"`
	Ingredients []map[string]Ingredient `yaml:"ingredients"`
	Steps       []Step                  `yaml:"steps"`
	Notes       []string                `yaml:"notes"`
}

func NewRecipeFromFilename(filename string) (*Recipe, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	o := &Recipe{}
	err = yaml.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func renderRecipeMarkdown(w io.Writer, r *Recipe) error {
	tmpl := template.New("recipe_template.md")

	tmpl, err := tmpl.ParseFiles("./recipe_template.md")
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, r)
	if err != nil {
		return err
	}

	return nil
}

func getFilenameFromArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("not enough args, expected filename as first arg")
	}
	return os.Args[1], nil
}

func main() {
	filename, err := getFilenameFromArgs()
	if err != nil {
		log.Fatalln(err)
	}
	r, err := NewRecipeFromFilename(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if err := renderRecipeMarkdown(os.Stdout, r); err != nil {
		log.Fatalln(err)
	}
}
