package leaked_credential_check_rules

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *LeakedCredentialCheckRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Leaked Credential Check Rules resource for managing user-defined Leaked Credential detection patterns within a specific zone.",
		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"rule": schema.SetNestedBlock{
				Description: "List of user-defined patterns for Leaked Credential Check",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						consts.IDSchemaKey: schema.StringAttribute{
							Description: consts.IDSchemaDescription,
							Computed:    true,
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
				},
			},
		},
	}
}
