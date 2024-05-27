// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r LogpullRetentionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"flag": schema.BoolAttribute{
				Description: "The log retention flag for Logpull API.",
				Required:    true,
			},
		},
	}
}
