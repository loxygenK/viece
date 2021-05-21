package model

type DeinPlugin struct {
	Frozen           int      `toml:"frozen"`
	Local            int      `toml:"local"`
	Depends          []string `toml:"depends"`
	On_i             int      `toml:"on_i"`
	On_idle          int      `toml:"on_idle"`
	On_ft            []string `toml:"on_ft"`
	On_cmd           []string `toml:"on_cmd"`
	On_lua           []string `toml:"on_lua"`
	On_map           []string `toml:"on_map"`
	On_path          []string `toml:"on_path"`
	On_source        []string `toml:"on_source"`
	Build            string   `toml:"build"`
	Hook_add         string   `toml:"hook_add"`
	Hook_source      string   `toml:"hook_source"`
	Hook_post_source string   `toml:"hook_post_source"`
	Hook_post_update string   `toml:"hook_post_update"`
	Repo             string   `toml:"repo"`
}

type DeinPluginConfiguration struct {
	Plugins []DeinPlugin `toml:"plugins"`
}

type DeinPluginConfigurationFile struct {
	Filename      string
	Configuration DeinPluginConfiguration
}
