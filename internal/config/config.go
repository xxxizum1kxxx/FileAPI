package config

type Config struct {
	Server struct {
		Port int    `json:"port"`
		Mode string `json:"mode"`
	} `json:"server"`
	Filesystem struct {
		Path string `json:"path"`
	} `json:"filesystem"`
}
