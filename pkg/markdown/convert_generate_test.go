package markdown

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/avelino/awesome-go/pkg/slug"
)

type IDGenerator struct {
	used map[string]bool
}

func (g *IDGenerator) Generate(value []byte, _ int) []byte {
	return []byte(slug.Generate(string(value)))
}

// TestGenerate is a test function for the Generate function.
func TestGenerate(t *testing.T) {
	// define test cases
	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Valid Slug Generation Test",
			input:          "Hello World",
			expectedOutput: "hello-world",
		},
		{
			name:           "Empty String Test",
			input:          "",
			expectedOutput: "",
		},
		{
			name:           "Non-English Characters Test",
			input:          "こんにちは世界",
			expectedOutput: "こんにちは世界",
		},
		{
			name:           "Long String Test",
			input:          "This is a long string used for testing the Generate function in the IDGenerator struct",
			expectedOutput: "this-is-a-long-string-used-for-testing-the-generate-function-in-the-idgenerator-struct",
		},
	}

	// create an instance of IDGenerator
	g := IDGenerator{
		used: make(map[string]bool),
	}

	// run tests
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Running:", tc.name)

			// generate slug
			slug := g.Generate([]byte(tc.input), 0)

			// check if the output is as expected
			assert.Equal(t, tc.expectedOutput, string(slug), "The output slug does not match the expected output")

			t.Log("Success:", tc.name)
		})
	}
}
