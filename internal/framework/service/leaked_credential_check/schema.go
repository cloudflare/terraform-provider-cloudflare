package leaked_credential_check

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *LeakedCredentialCheckResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Leaked Credential Check resource to be used for managing the status of the Cloudflare Leaked Credential detection within a specific zone.",
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "State of the Leaked Credential Check detection",
				Required:    true,
			},
		},
	}
}
