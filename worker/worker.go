package worker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type AsukaWorker struct {
	client *client.Client
	ctx    context.Context
}

func NewAsukaWorker(c *client.Client, ctx context.Context) *AsukaWorker {
	return &AsukaWorker{
		client: c,
		ctx:    ctx,
	}
}

func (n *AsukaWorker) NodeInit(params swarm.InitRequest) (*string, error) {

	// Get the Swarm info
	swarm, err := n.client.SwarmInit(n.ctx, params)
	if err != nil {
		return nil, err
	}

	return &swarm, nil
}

func (n *AsukaWorker) NodeRemove() error {
	err := n.client.SwarmLeave(n.ctx, true)
	if err != nil {
		return err
	}

	return nil
}

func (n *AsukaWorker) NodeJoin(params swarm.JoinRequest) error {
	err := n.client.SwarmJoin(n.ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (n *AsukaWorker) NodeInfo() (*swarm.Swarm, error) {

	sw, err := n.client.SwarmInspect(n.ctx)
	if err != nil {
		return nil, err
	}

	return &sw, nil
}

func (n *AsukaWorker) GetNodeToken(role string) (*string, error) {

	swarmInfo, err := n.client.SwarmInspect(n.ctx)
	if err != nil {
		return nil, err
	}

	// Get the join token for the specified role
	var token string
	switch role {
	case "manager":
		token = swarmInfo.JoinTokens.Manager
	case "worker":
		token = swarmInfo.JoinTokens.Worker
	default:
		return nil, fmt.Errorf("unknown role: %s", role)
	}

	return &token, nil

}

func (n *AsukaWorker) CreateWorkerServices(spec swarm.ServiceSpec, opt types.ServiceCreateOptions) (*string, error) {

	services, err := n.client.ServiceCreate(n.ctx, spec, opt)
	if err != nil {
		return nil, err
	}

	if len(services.Warnings) > 0 {
		return nil, errors.New(strings.Join(services.Warnings, ","))
	}

	return &services.ID, nil
}

func (n *AsukaWorker) InspectWorkerServices(services string, opt types.ServiceInspectOptions) (*swarm.Service, []byte, error) {

	_services, body, err := n.client.ServiceInspectWithRaw(n.ctx, services, opt)
	if err != nil {
		return nil, nil, err
	}

	return &_services, body, nil
}

func (n *AsukaWorker) ListWorkerServices(opt types.ServiceListOptions) (*[]swarm.Service, error) {

	list, err := n.client.ServiceList(n.ctx, opt)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (n *AsukaWorker) WokrkerLogsServices(service string, opt container.LogsOptions, f func(rd io.Reader) error) error {

	res, err := n.client.ServiceLogs(n.ctx, service, opt)
	if err != nil {
		return err
	}
	defer res.Close()

	// Hook function
	return f(res)
}

func (n *AsukaWorker) RemoveWorkerServices(service string) error {

	err := n.client.ServiceRemove(n.ctx, service)
	if err != nil {
		return err
	}

	return nil
}
