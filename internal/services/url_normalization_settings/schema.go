// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r URLNormalizationSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the URL normalization.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of URL normalization performed by Cloudflare.",
				Optional:    true,
			},
		},
	}
}
