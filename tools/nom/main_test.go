package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderRecipeMarkdown(t *testing.T) {
	r, err := NewRecipeFromFilename("./testdata/spaghetti-carbonara.yaml")
	require.NoError(t, err)

	buf := &bytes.Buffer{}
	err = renderRecipeMarkdown(buf, r)
	require.NoError(t, err)

	expMarkdown, err := os.ReadFile("./testdata/spaghetti-carbonara.md")
	require.NoError(t, err)

	require.Equal(t, string(expMarkdown), buf.String())
}
