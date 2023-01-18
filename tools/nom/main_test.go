package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCaseRenderRecipeMarkdown struct {
	desc          string
	inputFilename string
	expMarkdown   string
}

func TestCaseRenderRecipeMarkdown(t *testing.T) {
	testCases := []testCaseRenderRecipeMarkdown{
		{
			desc:          "spaghetti carbonara",
			inputFilename: "./testdata/spaghetti-carbonara.yaml",
			expMarkdown:   "./testdata/spaghetti-carbonara.md",
		},
		{
			desc:          "spaghetti carbonara",
			inputFilename: "./testdata/nonotes.yaml",
			expMarkdown:   "./testdata/nonotes.md",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			testRenderRecipeMarkdown(t, tC)
		})
	}
}

func testRenderRecipeMarkdown(t *testing.T, tC testCaseRenderRecipeMarkdown) {
	r, err := NewRecipeFromFilename(tC.inputFilename)
	require.NoError(t, err)

	buf := &bytes.Buffer{}
	err = renderRecipeMarkdown(buf, r)
	require.NoError(t, err)

	expMarkdown, err := os.ReadFile(tC.expMarkdown)
	require.NoError(t, err)

	require.Equal(t, string(expMarkdown), buf.String())
}
