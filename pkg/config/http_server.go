package config

import "fmt"

type HTTPServerConfig struct {
	host             string
	port             string
	readTimoutInSec  int
	writeTimoutInSec int
}

func newHTTPServerConfig() HTTPServerConfig {
	return HTTPServerConfig{
		host:             getString("HTTP_SERVER_HOST", "localhost"),
		port:             getString("PORT", "8080"),
		readTimoutInSec:  getInt("HTTP_SERVER_READ_TIMEOUT_IN_SEC"),
		writeTimoutInSec: getInt("HTTP_SERVER_WRITE_TIMEOUT_IN_SEC"),
	}
}

func (sc HTTPServerConfig) GetAddress() string {
	return fmt.Sprintf(":%s", sc.port)
}

func (sc HTTPServerConfig) GetReadTimeout() int {
	return sc.readTimoutInSec
}

func (sc HTTPServerConfig) GetWriteTimeout() int {
	return sc.readTimoutInSec
}
