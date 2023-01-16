---
title: "{{ .RecipeName }}"
---

## Ingredients
{{ range $i, $ingredientMap := .Ingredients -}}
{{ range $ingredientName, $ingredient := $ingredientMap -}}
* {{ $ingredientName }}
{{ end }}
{{- end }}

TODO
## Method
1. Do a foo
1. Bish a bosh
