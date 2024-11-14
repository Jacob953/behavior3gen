package behavior3gen

import (
	"github.com/Jacob953/behavior3gen/internal/generate"
	"github.com/Jacob953/behavior3gen/internal/model"
	"github.com/magicsea/behavior3go/core"
)

type genInfo struct {
	*generate.NodeMeta
}

func NewGenerator(cfg Config) *Generator {
	err := cfg.Revise()
	if err != nil {
		panic(err)
	}
	return &Generator{
		Config: cfg,
		Data:   make(map[string]*genInfo),
	}
}

type Generator struct {
	Config
	Data  map[string]*genInfo
	Model []*generate.NodeMeta
}

func (g *Generator) GenerateModelFrom(node map[string]core.IBaseNode) {
	for k, v := range node {
		s, err := generate.GetMetaFromObject(k, v, g.genModelConfig())
		if err != nil {
			panic(err)
		}
		g.Model = append(g.Model, s)
	}
}

func (g *Generator) genModelConfig() *model.Config {
	return &model.Config{
		Version: g.Version,
	}
}
