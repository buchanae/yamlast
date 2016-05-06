package yamlast

import (
	"testing"

	"github.com/stvp/assert"
)

func TestDecode(t *testing.T) {
	template := "v: 5"
	node, _ := Parse([]byte(template))

	if node != nil {
		assert.Equal(t, node.Kind, DocumentNode)
		assert.Equal(t, len(node.Children), 1)

		child := node.Children[0]
		assert.Equal(t, len(child.Children), 2)
		assert.Equal(t, child.Children[1].Implicit, true)
	}
}
