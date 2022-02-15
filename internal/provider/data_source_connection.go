package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/batchcorp/terraform-provider-plumber/plumber"
)

func dataSourceConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceConnectionRead,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"id": {
				Description: "Connection ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Connection name",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	connection, moreDiags := pc.GetConnection(filters)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	d.SetId(connection["_id"].(string))
	d.Set("name", connection["name"].(string))

	return diags
}
