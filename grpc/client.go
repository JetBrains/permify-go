package permifygrpc

import (
	"google.golang.org/grpc"

	pclient "github.com/Permify/permify/pkg/pb/base/v1"
)

// Client - Permify client
type Client struct {
	Permission   pclient.PermissionClient
	Schema       pclient.SchemaClient
	Data         pclient.DataClient
	Bundle       pclient.BundleClient
	Tenancy      pclient.TenancyClient
	Watch        pclient.WatchClient
	SharedSchema pclient.SharedSchemaClient
}

// Config - Permify client configuration
type Config struct {
	Endpoint string
	Cert     byte
}

// NewClient - Creates new Permify client
func NewClient(c Config, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.NewClient(c.Endpoint, opts...)
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	return &Client{
		Permission:   pclient.NewPermissionClient(conn),
		Schema:       pclient.NewSchemaClient(conn),
		Data:         pclient.NewDataClient(conn),
		Bundle:       pclient.NewBundleClient(conn),
		Tenancy:      pclient.NewTenancyClient(conn),
		Watch:        pclient.NewWatchClient(conn),
		SharedSchema: pclient.NewSharedSchemaClient(conn),
	}, nil
}
