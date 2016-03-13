package config

type Config struct {
	Memcachedbeat MemcachedbeatConfig
}

type MemcachedbeatConfig struct {
	Period string `yaml:"period"`
}
