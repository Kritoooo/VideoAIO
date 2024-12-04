package config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Video struct {
		UploadDir string `yaml:"upload_dir"`
		MaxSize   int64  `yaml:"max_size"` // 单位：MB
	} `yaml:"video"`
}

var GlobalConfig Config
