package plumber

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"google.golang.org/grpc"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
)

type Plumber struct {
	Token    string
	Client   protos.PlumberServerClient
	grpcConn *grpc.ClientConn
}

func New(address, token string) (*Plumber, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	// TODO: configure
	timeout := time.Second * 10

	dialContext, _ := context.WithTimeout(context.Background(), timeout)

	conn, err := grpc.DialContext(dialContext, address, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to grpc address '%s': %s", address, err)
	}

	client := protos.NewPlumberServerClient(conn)

	return &Plumber{
		Token:    token,
		Client:   client,
		grpcConn: conn,
	}, nil
}

func (p *Plumber) Close() error {
	return p.grpcConn.Close()
}

// GetConnection obtains a connection for a data source
func (p *Plumber) GetConnection(filters []*Filter) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	resp, err := p.Client.GetAllConnections(context.Background(), &protos.GetAllConnectionsRequest{
		Auth: &common.Auth{Token: p.Token},
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}

	respBytes, err := json.Marshal(resp.GetOptions())
	if err != nil {
		return nil, diag.FromErr(err)
	}

	raw := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(respBytes, &raw); err != nil {
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to parse response",
			Detail:   err.Error(),
		})
	}

	connections, moreDiags := filterJSON(raw, filters)
	if moreDiags.HasError() {
		return nil, moreDiags
	}

	if len(connections) < 1 {
		// No connection found using filter
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to find connection",
			Detail:   "Filters: " + filterString(filters),
		})
	} else if len(connections) > 1 {
		// Filter must find only one connection
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Filter returned more than one connection",
		})
	}

	return connections[0], diags
}

// GetTunnel obtains a tunnel config for a data source
func (p *Plumber) GetTunnel(filters []*Filter) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	resp, err := p.Client.GetAllTunnels(context.Background(), &protos.GetAllTunnelsRequest{
		Auth: &common.Auth{Token: p.Token},
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}

	respBytes, err := json.Marshal(resp.Opts)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	raw := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(respBytes, &raw); err != nil {
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to parse response",
			Detail:   err.Error(),
		})
	}

	tunnels, moreDiags := filterJSON(raw, filters)
	if moreDiags.HasError() {
		return nil, moreDiags
	}

	if len(tunnels) < 1 {
		// No connection found using filter
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to find tunnel",
			Detail:   "Filters: " + filterString(filters),
		})
	} else if len(tunnels) > 1 {
		// Filter must find only one tunnel
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Filter returned more than one tunnel",
		})
	}

	return tunnels[0], diags
}

// GetRelay obtains a relay for a data source
func (p *Plumber) GetRelay(filters []*Filter) (map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	resp, err := p.Client.GetAllRelays(context.Background(), &protos.GetAllRelaysRequest{
		Auth: &common.Auth{Token: p.Token},
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}

	respBytes, err := json.Marshal(resp.Opts)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	raw := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(respBytes, &raw); err != nil {
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to parse response",
			Detail:   err.Error(),
		})
	}

	relay, moreDiags := filterJSON(raw, filters)
	if moreDiags.HasError() {
		return nil, moreDiags
	}

	if len(relay) < 1 {
		// No connection found using filter
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to find relay",
			Detail:   "Filters: " + filterString(filters),
		})
	} else if len(relay) > 1 {
		// Filter must find only one tunnel
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Filter returned more than one relay",
		})
	}

	return relay[0], diags
}
