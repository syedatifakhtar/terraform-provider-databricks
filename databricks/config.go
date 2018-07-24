package databricks

import (
	"github.com/betabandido/databricks-sdk-go/api/clusters"
	"github.com/betabandido/databricks-sdk-go/api/workspace"
	apiClient "github.com/betabandido/databricks-sdk-go/client"
)

type Config struct {
	Domain *string
	Token  *string
}

type Client struct {
	clusters  *clusters.Endpoint
	workspace *workspace.Endpoint
}

func (c *Config) Client() (interface{}, error) {
	var client Client

	opts := apiClient.Options{Domain: c.Domain, Token: c.Token}
	cl, err := apiClient.NewClient(opts)
	if err != nil {
		return nil, err
	}

	client.clusters = &clusters.Endpoint{Client: cl}
	client.workspace = &workspace.Endpoint{Client: cl}

	return &client, nil
}
