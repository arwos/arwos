package runner

type ConfigGroup struct {
	Config Config `yaml:"runner"`
}

type Config struct {
	Folder string `yaml:"folder"`
}
