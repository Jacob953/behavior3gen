package behavior3gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Jacob953/behavior3gen/internal/generate"
	"github.com/Jacob953/behavior3gen/internal/model"
	"github.com/magicsea/behavior3go/core"
	"io"
	"os"
	"path/filepath"
	"sort"
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
		model:  make(map[string]*generate.NodeMeta),
	}
}

type Generator struct {
	Config
	Data  map[string]*genInfo
	model map[string]*generate.NodeMeta
}

func (g *Generator) GenerateNodeAs(name string, node core.IBaseNode) *generate.NodeMeta {
	s, err := generate.GetMetaFromObject(node, g.genModelConfig())
	if err != nil {
		panic(err)
	}
	if !g.EnableNameGen {
		s.Name = name
	}
	_, ok := g.model[s.Name]
	if ok {
		panic("model already exists")
	}
	g.model[s.Name] = s
	return s
}

func (g *Generator) GenerateNodeFrom(node map[string]core.IBaseNode) {
	for k, v := range node {
		g.GenerateNodeAs(k, v)
	}
}

func (g *Generator) Execute() {
	err := g.generateNodeFile()
	if err != nil {
		panic(err)
	}
}

func (g *Generator) genModelConfig() *model.Config {
	return &model.Config{
		Version: g.Version,
	}
}

func (g *Generator) generateNodeFile() error {
	if len(g.model) == 0 {
		return nil
	}
	err := encode(os.Stdout, g.getNodeList())
	if err != nil {
		return err
	}
	nodeOutPath, err := g.genNodeOutputPath()
	if err != nil {
		panic(err)
	}
	if err = os.MkdirAll(nodeOutPath, os.ModePerm); err != nil {
		return fmt.Errorf("create model pkg path(%s) fail: %s", nodeOutPath, err)
	}
	var buf bytes.Buffer
	err = encode(&buf, g.getNodeList())
	if err != nil {
		return err
	}
	err = os.WriteFile(g.OutFile, buf.Bytes(), 0640)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) genNodeOutputPath() (outputPath string, err error) {
	outPath := filepath.Join(filepath.Dir(g.OutPath), "")
	return outPath + string(os.PathSeparator), nil
}

func (g *Generator) getNodeList() []*generate.NodeMeta {
	res := make([]*generate.NodeMeta, 0, len(g.model))
	for _, node := range g.model {
		res = append(res, node)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Category <= res[j].Category && res[i].Name < res[j].Name
	})
	return res
}

func encode(w io.Writer, data any) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.SetEscapeHTML(false)
	return e.Encode(data)
}
