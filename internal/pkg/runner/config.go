package runner

type ConfigGroup struct {
	Runner RunnerConfig `yaml:"runner"`
}

type RunnerConfig struct {
	Folder string `yaml:"folder"`
}

func (v *ConfigGroup) Default() {
	if len(v.Runner.Folder) == 0 {
		v.Runner.Folder = "./.plugins"
	}
}
