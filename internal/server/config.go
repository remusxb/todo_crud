package server

import "time"

// Config to set up the networking of the server.
type Config struct {
	// Host or IP to listen on.  Defaults to localhost to start safe.
	Host string `mapstructure:"host" env:"HOST"`

	// Port to listen on
	Port int `mapstructure:"port" env:"PORT"`

	// TLS enabled
	TLS bool `mapstructure:"tls" env:"TLS"`

	// CertFile to use for TLS
	CertFile string `mapstructure:"tls_cert" env:"CERT_FILE"`

	// KeyFile to use for TLS
	KeyFile string `mapstructure:"tls_key" env:"KEY_FILE"`

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	ReadTimeout time.Duration `mapstructure:"read_timeout" env:"READ_TIMEOUT"`

	// WriteTimeout is the maximum duration before timing out
	// writes of the response.
	WriteTimeout time.Duration `mapstructure:"write_timeout" env:"WRITE_TIMEOUT"`

	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT"`
}
