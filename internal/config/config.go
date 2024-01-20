package config

import (
	"os"
	"strconv"

	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(Parse),
)

type Config struct {
	GRPCPort int

	ProxmoxHost     string
	ProxmoxUsername string
	ProxmoxPassword string
	ProxmoxRealm    string
}

func Parse() (Config, error) {
	config := Config{}

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return config, err
	}

	config.GRPCPort = grpcPort
	config.ProxmoxHost = os.Getenv("PROXMOX_HOST")
	config.ProxmoxUsername = os.Getenv("PROXMOX_USERNAME")
	config.ProxmoxPassword = os.Getenv("PROXMOX_PASSWORD")
	config.ProxmoxRealm = os.Getenv("PROXMOX_REALM")

	return config, nil
}
