package config

import (
	"fmt"
	"os"
)

const EndpointEnvVar = "CSI_ENDPOINT"

type Config struct {
	CSIEndpoint string
}

func NewConfig() (Config, error) {
	csiEndpoint := os.Getenv(EndpointEnvVar)

	if csiEndpoint == "" {
		return Config{}, fmt.Errorf("Environment variable %v not defined", EndpointEnvVar)
	}

	return Config{
		CSIEndpoint: csiEndpoint,
	}, nil
}
