package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Step struct {
	Step  string   `yaml:"step"`
	Notes []string `yaml:"notes"`
}

func (s *Step) String() string {
	return appendNotes(s.Step, s.Notes)
}

type Amount struct {
	Amount float64 `yaml:"amount"`
	Unit   string  `yaml:"unit"`
}

func (a *Amount) String() string {
	return fmt.Sprintf("%v %v", a.Amount, a.Unit)
}

type Ingredient struct {
	Amounts []Amount `yaml:"amounts"`
	Notes   []string `yaml:"notes"`
}

type Recipe struct {
	RecipeName  string                  `yaml:"recipe_name"`
	Yields      []Amount                `yaml:"yields"`
	Tags        []string                `yaml:"tags"`
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

func appendNotes(prefix string, notes []string) string {
	if len(notes) <= 0 {
		return prefix
	}
	notesJoined := strings.Join(notes, " ")
	return fmt.Sprintf("%v (%v)", prefix, notesJoined)
}

// printIngredient prints OpenRecipeFormat Ingredient structs because the name isn't stored
func printIngredient(name string, i Ingredient) string {
	amounts := []string{}
	for _, a := range i.Amounts {
		amounts = append(amounts, a.String())
	}
	amountsJoined := strings.Join(amounts, " / ")
	output := fmt.Sprintf("%v %v", amountsJoined, name)
	output = appendNotes(output, i.Notes)
	return output
}

//go:embed templates/recipe.md
var recipeTemplate string

func renderRecipeMarkdown(w io.Writer, r *Recipe) error {
	tmpl := template.New("recipe.md")

	tmpl = tmpl.Funcs(template.FuncMap{
		"printIngredient": printIngredient,
	})

	tmpl, err := tmpl.Parse(recipeTemplate)
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
