package model

import (
	"io/ioutil"
	"strings"
	"github.com/pelletier/go-toml"
)

type DeinPlugin struct {
  Frozen int              `toml:"frozen"`
  Local int               `toml:"local"`
  Depends []string        `toml:"depends"`
  On_i int                `toml:"on_i"`
  On_idle int             `toml:"on_idle"`
  On_ft []string          `toml:"on_ft"`
  On_cmd []string         `toml:"on_cmd"`
  On_lua []string         `toml:"on_lua"`
  On_map []string         `toml:"on_map"`
  On_path []string        `toml:"on_path"`
  On_source []string      `toml:"on_source"`
  Build string            `toml:"build"`
  Hook_add string         `toml:"hook_add"`
  Hook_source string      `toml:"hook_source"`
  Hook_post_source string `toml:"hook_post_source"`
  Hook_post_update string `toml:"hook_post_update"`
  Repo string             `toml:"repo"`
}

type DeinPluginConfiguration struct {
  Plugins []DeinPlugin `toml:"plugins"`
}

type DeinPluginConfigurationFile struct {
  Filename string
  Configuration DeinPluginConfiguration
}

func LoadPluginConfig(filename string) (DeinPluginConfigurationFile, error) {
  rawfile, err := ioutil.ReadFile(filename)
  if err != nil {
    return getEmptyConfig(filename), err
  }
  escapedfile := strings.Replace(string(rawfile), "\\", "\\\\", -1)
  file := []byte(escapedfile);

  config := DeinPluginConfiguration{}
  tomlerr := toml.Unmarshal(file, &config)
  if tomlerr != nil {
    return getEmptyConfig(filename), tomlerr
  }
  
  return DeinPluginConfigurationFile{ Filename: filename, Configuration: config }, nil;
}

func getEmptyConfig(filename string) DeinPluginConfigurationFile {
  return DeinPluginConfigurationFile{
    Filename: filename,
    Configuration: DeinPluginConfiguration{ Plugins: []DeinPlugin{} },
  }
}

