// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r AccessKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"key_rotation_interval_days": schema.Float64Attribute{
				Description: "The number of days between key rotations.",
				Required:    true,
			},
		},
	}
}
