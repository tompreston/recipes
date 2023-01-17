---
title: "{{ .RecipeName }}"
---

## Ingredients
{{ range $i, $ingredientMap := .Ingredients -}}
{{ range $name, $ingredient := $ingredientMap -}}
* {{ printIngredient $name $ingredient }}
{{ end -}}
{{ end }}
## Steps
{{ range $i, $step := .Steps -}}
* {{ $step }}
{{ end }}
## Notes
{{ range $i, $note := .Notes -}}
* {{ $note }}
{{- end }}
