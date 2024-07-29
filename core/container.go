package core

import (
	"errors"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

func (a *Asuka) ContainerCreate(cfg *container.Config, host *container.HostConfig,
	network *network.NetworkingConfig, platform *ocispec.Platform, name string) (*string, error) {

	res, err := a.client.ContainerCreate(a.ctx, cfg, host, network, platform, name)
	if err != nil {
		return nil, err
	}

	if len(res.Warnings) > 0 {
		return nil, errors.New(strings.Join(res.Warnings, ", "))
	}

	return &res.ID, nil
}
