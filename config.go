package behavior3gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	OutPath       string
	OutFile       string
	Version       string
	EnableNameGen bool
}

func (cfg *Config) Revise() (err error) {
	if cfg.OutPath == "" {
		cfg.Version = "0.3.0"
	}
	if cfg.OutPath == "" {
		cfg.OutPath = fmt.Sprintf(".%sconfig%s", string(os.PathSeparator), string(os.PathSeparator))
	}
	if cfg.OutFile == "" {
		cfg.OutFile = filepath.Join(cfg.OutPath, "node.b3")
	} else if !strings.Contains(cfg.OutFile, string(os.PathSeparator)) {
		cfg.OutFile = filepath.Join(cfg.OutPath, cfg.OutFile)
	}
	return nil
}
