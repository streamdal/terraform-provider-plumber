package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/batchcorp/terraform-provider-plumber/plumber"
)

func dataSourceTunnel() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceTunnelRead,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"id": {
				Description: "Tunnel ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_id": {
				Description: "Connection ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceTunnelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var filters []*plumber.Filter

	pc := m.(*plumber.Plumber)

	if v, ok := d.GetOk("filter"); ok {
		filters = buildFiltersDataSource(v.(*schema.Set))
	} else {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No filters defined",
			Detail:   "At least one filter must be defined",
		})
	}

	dynamic, moreDiags := pc.GetTunnel(filters)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	d.SetId(dynamic["_dynamic_id"].(string))
	d.Set("name", dynamic["name"].(string))

	return diags
}
