package generate

import (
	"github.com/magicsea/behavior3go/core"
	"testing"
)

func TestReflectBaseNode(t *testing.T) {
	actionNode := core.Action{}
	SetBaseNodePrivate(&actionNode.BaseNode, "desc", "category")
}
