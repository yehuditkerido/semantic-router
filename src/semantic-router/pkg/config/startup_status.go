package config

// StartupStatusConfig controls how the router publishes its startup
// progress so the dashboard can display model download and initialization state.
type StartupStatusConfig struct {
	Backend string                    `yaml:"backend" json:"backend"`
	Redis   *StartupStatusRedisConfig `yaml:"redis,omitempty" json:"redis,omitempty"`
}

// StartupStatusRedisConfig holds the Redis connection details for startup status.
type StartupStatusRedisConfig struct {
	Address  string `yaml:"address" json:"address"`
	Password string `yaml:"password,omitempty" json:"password,omitempty"`
	DB       int    `yaml:"db,omitempty" json:"db,omitempty"`
}
