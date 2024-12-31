package content_scanning_expression

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *ContentScanningExpressionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Content Scanning Expression resource for managing custom scan expression within a specific zone.",
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Description: consts.IDSchemaDescription,
				Computed:    true,
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
			"payload": schema.StringAttribute{
				Description: "Custom scan expression to tell the content scanner where to find the content objects.",
				Required:    true,
			},
		},
	}
}
