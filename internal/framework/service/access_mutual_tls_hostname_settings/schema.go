package access_mutual_tls_hostname_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *AccessMutualTLSHostnameSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare Access Mutual TLS Certificate Settings resource.",
		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				Description: consts.AccountIDSchemaDescription,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot(consts.ZoneIDSchemaKey)),
					),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.Expression(path.MatchRoot(consts.AccountIDSchemaKey)),
					),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"settings": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							Required:    true,
							Description: "The hostname that these settings apply to.",
						},
						"china_network": schema.BoolAttribute{
							Optional:    true,
							Description: "Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.",
						},
						"client_certificate_forwarding": schema.BoolAttribute{
							Optional:    true,
							Description: "Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
		},
	}
}
