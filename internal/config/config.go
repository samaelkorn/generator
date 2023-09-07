package config

type Config struct {
	Dir *Dir
}

func getConfig() *Config {
	return &Config{
		Dir: getDir(),
	}
}
