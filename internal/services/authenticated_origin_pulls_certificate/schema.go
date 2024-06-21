// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r AuthenticatedOriginPullsCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The zone's leaf certificate.",
				Required:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "The zone's private key.",
				Required:    true,
			},
		},
	}
}
