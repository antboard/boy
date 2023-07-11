package config

type SimpleConfig struct {
	// 一个简单的map 实现kv读写.
	data map[string]string
}

func NewSimpleConfig() *SimpleConfig {
	return &SimpleConfig{data: make(map[string]string)}
}

func (s *SimpleConfig) Get(key string) string {
	return s.data[key]
}

func (s *SimpleConfig) Set(key string, value string) {
	s.data[key] = value
}
