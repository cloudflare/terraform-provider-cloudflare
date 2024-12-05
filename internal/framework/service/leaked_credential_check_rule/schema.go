package leaked_credential_check_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *LeakedCredentialCheckRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Leaked Credential Check Rule resource for managing user-defined Leaked Credential detection patterns within a specific zone.",
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Description: consts.IDSchemaDescription,
				Computed:    true,
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
			"username": schema.StringAttribute{
				Description: "The ruleset expression to use in matching the username in a request.",
				Required:    true,
			},
			"password": schema.StringAttribute{
				Description: "The ruleset expression to use in matching the password in a request",
				Required:    true,
			},
		},
	}
}
