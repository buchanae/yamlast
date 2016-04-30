package yamlast

import (
	"testing"

	"github.com/stvp/assert"
)

func TestDecode(t *testing.T) {
	template := "v: hi"
	node := Parse([]byte(template))

	if node != nil {
		assert.Equal(t, node.kind, documentNode)
	}
}
