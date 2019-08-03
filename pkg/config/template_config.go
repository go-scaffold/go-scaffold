package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type TemplateConfig struct {
	koanf *koanf.Koanf
}

func NewTemplateConfig(configFilePath string) (*TemplateConfig, error) {
	v := &TemplateConfig{
		koanf: koanf.New("."),
	}
	err := v.koanf.Load(file.Provider(configFilePath), yaml.Parser())
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (t *TemplateConfig) Config() map[string]interface{} {
	return t.koanf.Raw()
}
