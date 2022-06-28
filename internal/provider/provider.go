package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/batchcorp/terraform-provider-plumber/plumber"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var types = []string{"kafka", "rabbit"}

func init() {
	schema.DescriptionKind = schema.StringMarkdown
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" (Default: `%v`)", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New(version, apiToken string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"plumber_token": {
					Description: "Plumber API token",
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("PLUMBER_TOKEN", apiToken),
				},
				"address": {
					Description: "The address of the Plumber server.",
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("PLUMBER_ADDRESS", "localhost:9090"),
				},
				"connection_timeout": {
					Description: "The connection timeout for the Plumber server.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     schema.EnvDefaultFunc("PLUMBER_CONNECTION_TIMEOUT", 10),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"plumber_connection": resourceConnection(),
				"plumber_relay":      resourceRelay(),
				"plumber_tunnel":     resourceTunnel(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"plumber_connection": dataSourceConnection(),
				"plumber_relay":      dataSourceRelay(),
				"plumber_tunnel":     dataSourceTunnel(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		token := d.Get("plumber_token").(string)
		address := d.Get("address").(string)

		plumber, err := plumber.New(address, token)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return plumber, nil
	}
}

func dataSourceFiltersSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Description: "Field name to filter on",
					Type:        schema.TypeString,
					Required:    true,
				},

				"values": {
					Description: "Value(s) to filter by. Wildcards '*' are supported.",
					Type:        schema.TypeList,
					Required:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func buildFiltersDataSource(set *schema.Set) []*plumber.Filter {
	var filters []*plumber.Filter
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var filterValues []string
		for _, e := range m["values"].([]interface{}) {
			filterValues = append(filterValues, e.(string))
		}
		filters = append(filters, &plumber.Filter{
			Name:   m["name"].(string),
			Values: filterValues,
		})
	}
	return filters
}

// getBusType is a convenience function to return the message bus type we are operating on
func getBusType(d *schema.ResourceData) string {
	for _, bus := range types {
		if opts, ok := d.Get(bus).([]interface{}); ok && len(opts) > 0 {
			return bus
		}
	}

	return ""
}
