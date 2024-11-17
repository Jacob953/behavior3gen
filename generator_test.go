package behavior3gen

import (
	"github.com/Jacob953/behavior3gen/field"
	"github.com/magicsea/behavior3go/core"
	"testing"
)

type TestNode struct {
	core.Action
	Prefix string `json:"prefix"`
}

func TestGenerator(t *testing.T) {
	g := NewGenerator(Config{})
	//exp := []*generate.NodeMeta{
	//	{
	//		Version:     "0.3.0",
	//		Scope:       "node",
	//		Name:        "TEST_NODE",
	//		Category:    "action",
	//		Title:       "TestNode(<prefix>)",
	//		Description: "description",
	//		Properties: map[string]interface{}{
	//			"prefix": "prefix",
	//		},
	//	},
	//}
	node := map[string]core.IBaseNode{
		"TEST_NODE": &TestNode{
			Action: field.NewAction("description\n"),
			Prefix: "prefix",
		},
	}
	g.GenerateNodeFrom(node)
	g.Execute()
}
