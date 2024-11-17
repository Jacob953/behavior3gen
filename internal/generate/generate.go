package generate

import (
	"fmt"
	"github.com/Jacob953/behavior3gen/internal/model"
	"github.com/magicsea/behavior3go/core"
	"reflect"
	"strings"
)

type NodeMeta struct {
	Version     string         `json:"version"`
	Scope       string         `json:"scope"`
	Name        string         `json:"name"`
	Category    string         `json:"category"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Properties  map[string]any `json:"properties"`
}

func GetMetaFromObject(node core.IBaseNode, cfg *model.Config) (res *NodeMeta, err error) {
	res = &NodeMeta{
		Version:    cfg.Version,
		Scope:      "node",
		Properties: make(map[string]any),
	}
	v := reflect.ValueOf(node).Elem()
	tags := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}
		tags = append(tags, fmt.Sprintf("<%s>", tag))
		res.Properties[tag] = v.Field(i).Interface()
	}
	res.Name = v.Type().Name()
	res.Title = fmt.Sprintf("%s(%s)", v.Type().Name(), strings.Join(tags, ", "))
	// reflect to set private members of BaseNode
	baseNodeField := v.FieldByName("BaseNode")
	if !baseNodeField.IsValid() {
		return nil, fmt.Errorf("no BaseNode")
	}
	descField := baseNodeField.FieldByName("description")
	if !descField.IsValid() {
		return nil, fmt.Errorf("no description")
	}
	res.Description = descField.String()
	cateField := baseNodeField.FieldByName("category")
	if !descField.IsValid() {
		return nil, fmt.Errorf("no category")
	}
	res.Category = cateField.String()
	return
}
