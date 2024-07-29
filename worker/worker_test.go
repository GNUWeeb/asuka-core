package worker

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

func TestSwarmStop(t *testing.T) {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	list, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Fatalf("Failed get list swarm: %v", err)
	}
	if len(list) > 0 {
		for _, sw := range list {
			_ = cli.ServiceRemove(context.Background(), sw.ID)
		}
	}

	// cli.ServiceLogs(context.Background(), )
}

func TestSwarm(t *testing.T) {

	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	// Define the service
	serviceName := "my-service"
	image := "nginx:latest" // Replace with your desired image

	serviceSpec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: serviceName,
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image: image,
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: getUint64Pointer(3), // Number of replicas
			},
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: []swarm.PortConfig{
				{
					TargetPort:    80,
					PublishedPort: 8080,
					PublishMode:   "Ingress",
				},
			},
		},
	}

	// Create the service
	serviceCreateResp, err := cli.ServiceCreate(context.Background(), serviceSpec, types.ServiceCreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	fmt.Printf("Service created with ID: %s\n", serviceCreateResp.ID)
}

func getUint64Pointer(v uint64) *uint64 {
	return &v
}
