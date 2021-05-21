package parser

import (
	"github.com/loxygenK/viece/model"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"strings"
)

func LoadPluginConfig(filename string) (model.DeinPluginConfigurationFile, error) {
	rawcontent, readerr := ioutil.ReadFile(filename)
	if readerr != nil {
		return model.DeinPluginConfigurationFile{}, readerr
	}
	espaced := strings.Replace(string(rawcontent), "\\", "\\\\", -1)
	tomlcontent := []byte(espaced)

	config := model.DeinPluginConfiguration{}
	tomlerr := toml.Unmarshal(tomlcontent, &config)
	if tomlerr != nil {
		return model.DeinPluginConfigurationFile{}, tomlerr
	}

	return model.DeinPluginConfigurationFile{Filename: filename, Configuration: config}, nil
}

func LoadPluginConfigDirectory(directory string) ([]model.DeinPluginConfigurationFile, error) {
	files, err := findTomlFile(directory)
	if err != nil {
		return []model.DeinPluginConfigurationFile{}, err
	}

	var configs []model.DeinPluginConfigurationFile
	for _, file := range files {
		config, configerr := LoadPluginConfig(file)
		if configerr != nil {
			return []model.DeinPluginConfigurationFile{}, FileParseError{
				Filename: file, Additional: configerr,
			}
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func findTomlFile(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return []string{}, err
	}

	var tomlfiles []string
	for _, file := range files {
		fullpath := directory + "/" + file.Name()
		if file.IsDir() {
			children, childerr := findTomlFile(fullpath)
			if childerr != nil {
				return []string{}, childerr
			}
			tomlfiles = append(tomlfiles, children...)
			continue
		}
		if strings.HasSuffix(file.Name(), ".toml") {
			tomlfiles = append(tomlfiles, fullpath)
		}
	}

	return tomlfiles, nil
}
