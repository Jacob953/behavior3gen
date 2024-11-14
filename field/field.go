package field

import (
	"github.com/Jacob953/behavior3gen/internal/generate"
	b3 "github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
)

func NewAction(desc string) (res core.Action) {
	err := generate.SetBaseNodePrivate(&res.BaseNode, desc, b3.ACTION)
	if err != nil {
		panic(err)
	}
	return
}

func NewComposite(desc string) (res core.Composite) {
	err := generate.SetBaseNodePrivate(&res.BaseNode, desc, b3.COMPOSITE)
	if err != nil {
		panic(err)
	}
	return
}

func NewCondition(desc string) (res core.Condition) {
	err := generate.SetBaseNodePrivate(&res.BaseNode, desc, b3.CONDITION)
	if err != nil {
		panic(err)
	}
	return
}

func NewDecorator(desc string) (res core.Decorator) {
	err := generate.SetBaseNodePrivate(&res.BaseNode, desc, b3.DECORATOR)
	if err != nil {
		panic(err)
	}
	return
}
