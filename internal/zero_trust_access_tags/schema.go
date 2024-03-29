// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tags

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r ZeroTrustAccessTagsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"tag_name": schema.StringAttribute{
				Description: "The name of the tag",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the tag",
				Required:    true,
			},
			"app_count": schema.Int64Attribute{
				Description: "The number of applications that have this tag",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
