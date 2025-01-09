package local_fs

type Config struct {
	Enable       bool   `yaml:"enable"`
	HttpfsEnable bool   `yaml:"httpfs-enable"`
	SavePath     string `yaml:"save-path"`
}

type LocalFS struct {
	IsCheckSave bool
	Config      *Config
}

func NewClient(conf map[string]any) (*LocalFS, error) {

	config := &Config{
		Enable:       conf["Enable"].(bool),
		HttpfsEnable: conf["HttpfsEnable"].(bool),
		SavePath:     conf["SavePath"].(string),
	}
	return &LocalFS{
		Config: config,
	}, nil
}
