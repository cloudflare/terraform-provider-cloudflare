// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_hostname_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r AccessMutualTLSHostnameSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"settings": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"china_network": schema.BoolAttribute{
							Description: "Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.",
							Required:    true,
						},
						"client_certificate_forwarding": schema.BoolAttribute{
							Description: "Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.",
							Required:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname that these settings apply to.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}
