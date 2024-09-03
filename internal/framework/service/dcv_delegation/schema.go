package dcv_delegation

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func (r *DCVDelegationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to retrieve the DCV Delegation unique identifier for a zone.",
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Required:            true,
			},
			"id": schema.StringAttribute{
				Description: "The DCV Delegation unique identifier",
				Computed:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The DCV Delegation hostname",
				Computed:    true,
			},
		},
	}
}
