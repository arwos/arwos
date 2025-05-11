package webcache

type ConfigGroup struct {
	Static StaticConfig `yaml:"webui"`
}

type StaticConfig struct {
	Build  string `yaml:"build"`
	Source string `yaml:"source"`
}

func (v *ConfigGroup) Default() {
	if len(v.Static.Build) == 0 {
		v.Static.Build = "./web/dist/web/browser"
	}
	if len(v.Static.Source) == 0 {
		v.Static.Source = "./web"
	}
}
