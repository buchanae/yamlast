package yamlast

import (
	"testing"

	"github.com/stvp/assert"
)

func TestSelectKey(t *testing.T) {
	template := "v: 5"
	doc := Parse([]byte(template))

	selectedNode := SelectNode(doc, "v")
	assert.NotNil(t, selectedNode)
	assert.Equal(t, "5", selectedNode.Value)
}

func TestSelectStringKey(t *testing.T) {
	template := "key: foo"
	doc := Parse([]byte(template))

	selectedNode := SelectNode(doc, "\"key\"")
	assert.NotNil(t, selectedNode)
	assert.Equal(t, "foo", selectedNode.Value)
}

func TestSelectArray(t *testing.T) {
	template := `
    - a
    - test
  `
	doc := Parse([]byte(template))

	selectedNode := SelectNode(doc, "[1]")
	assert.NotNil(t, selectedNode)
	assert.Equal(t, "test", selectedNode.Value)
}

func TestComplex(t *testing.T) {
	template := `
    foo: bar
    baz:
      - this
      - is
      - something
      - cool: key
        other: rad
  `
	doc := Parse([]byte(template))
	selectedNode := SelectNode(doc, "baz[3].other")
	assert.NotNil(t, selectedNode)
	assert.Equal(t, selectedNode.Value, "rad")
}

func TestTrailingChars(t *testing.T) {
	template := "v: foo"
	doc := Parse([]byte(template))

	selectedNode := SelectNode(doc, "v[]")
	assert.Nil(t, selectedNode)
}

func TestTrailingSpace(t *testing.T) {
	template := "v: foo"
	doc := Parse([]byte(template))

	selectedNode := SelectNode(doc, "v ")
	assert.NotNil(t, selectedNode)
	assert.Equal(t, "foo", selectedNode.Value)
}
