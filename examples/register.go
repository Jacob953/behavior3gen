package main

import (
	"github.com/Jacob953/behavior3gen/field"
	"github.com/magicsea/behavior3go/core"
)

var Register = map[string]core.IBaseNode{
	"ActionNode":     NewActionNode(),    // recommended the same name for debug
	"CONDITION_NODE": NewConditionNode(), // need to stay unique names among all nodes
	// composite node won't be registered if defined only
}

func NewActionNode() *ActionNode {
	return &ActionNode{
		Action:   field.NewAction(descActionNode),
		Prefix:   "prefix-action",
		Duration: "duration-action",
	}
}

const descActionNode = "this is description of action"

type ActionNode struct {
	core.Action
	Prefix   string `json:"prefix"`
	Duration string `json:"duration"`
}

func NewConditionNode() *ConditionNode {
	return &ConditionNode{
		Condition: field.NewCondition(descConditionNode),
		Prefix:    "prefix-condition",
		Limit:     "limit-condition",
	}
}

const descConditionNode = "this is description of condition\nmore information of condition"

type ConditionNode struct {
	core.Condition
	Prefix string `json:"prefix"`
	Limit  string `json:"limit"`
}

func NewCompositeNode() *CompositeNode {
	return &CompositeNode{
		Composite:   field.NewComposite(""),
		Description: "description-composite",
	}
}

type CompositeNode struct {
	core.Composite
	Description string `json:"description"`
}
