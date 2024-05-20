package internal

import "github.com/spf13/viper"

// Server is a struct that represents the server
type Server struct {
	Host string
	Port int
}

// NewServer is a function to create a new server
func NewServerConfig(v *viper.Viper) *Server {
	return &Server{
		Host: v.GetString("app.host"),
		Port: v.GetInt("app.port"),
	}
}

