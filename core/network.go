package core

import (
	"errors"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

func (a *Asuka) NetworkList(opt network.ListOptions) (*[]network.Summary, error) {

	res, err := a.client.NetworkList(a.ctx, opt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *Asuka) NetworkCreate(name string, opt network.CreateOptions) (*string, error) {

	res, err := a.client.NetworkCreate(a.ctx, name, opt)
	if err != nil {
		return nil, err
	}

	if res.Warning != "" {
		return nil, errors.New(res.Warning)
	}

	return &res.ID, nil
}

func (a *Asuka) NetworkConnect(network, container string, opt *network.EndpointSettings) error {

	if err := a.client.NetworkConnect(a.ctx, network, container, opt); err != nil {
		return err
	}

	return nil
}

func (a *Asuka) NetworkDisconnect(network, container string, force bool) error {

	if err := a.client.NetworkDisconnect(a.ctx, network, container, force); err != nil {
		return err
	}

	return nil
}

func (a *Asuka) NetworkInspect(network string, opt network.InspectOptions) (*network.Inspect, []byte, error) {

	inspect, body, err := a.client.NetworkInspectWithRaw(a.ctx, network, opt)
	if err != nil {
		return nil, nil, err
	}

	return &inspect, body, nil
}

func (a *Asuka) NetworkRemove(network string) error {
	if err := a.client.NetworkRemove(a.ctx, network); err != nil {
		return err
	}

	return nil
}

func (a *Asuka) NetworksPrune(filter filters.Args) (*[]string, error) {

	deleted, err := a.client.NetworksPrune(a.ctx, filter)
	if err != nil {
		return nil, err
	}

	return &deleted.NetworksDeleted, nil
}
