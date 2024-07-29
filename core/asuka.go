package core

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

type Asuka struct {
	client *client.Client
	ctx    context.Context

	registryAuthString string
}

func NewAsuka(c *client.Client, ctx context.Context) *Asuka {
	return &Asuka{
		client: c,
		ctx:    ctx,
	}
}

func (a *Asuka) SetAuth(username, password, server string) *Asuka {
	authConfig := registry.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: server,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		log.Fatalf("Failed Marshal Auth Config.")
	}

	a.registryAuthString = base64.URLEncoding.EncodeToString(encodedJSON)
	return a
}

func (a *Asuka) Ping() error {

	_, err := a.client.Ping(a.ctx)
	if err != nil {
		return err
	}

	return nil
}
