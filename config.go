package behavior3gen

type Config struct {
	OutPath string
	OutFile string
	Version string
}

func (cfg *Config) Revise() (err error) {
	if cfg.OutPath == "" {
		cfg.Version = "0.3.0"
	}
	return nil
}
