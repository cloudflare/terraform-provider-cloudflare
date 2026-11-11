// This file is intentionally NOT code-generated.

package worker_version

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceSchemaV0Assets is the assets attribute as version 0 stored it:
// run_worker_first was a list of strings and base_path did not exist. It
// must not reuse the latest schema's custom types, which carry the latest
// assets.config model, or decoding into resourceModelV0 fails.
func resourceSchemaV0Assets(ctx context.Context) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[resourceModelV0AssetsConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"html_handling":      schema.StringAttribute{Optional: true},
					"not_found_handling": schema.StringAttribute{Optional: true},
					"run_worker_first": schema.ListAttribute{
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
			"jwt":                   schema.StringAttribute{Optional: true, Sensitive: true},
			"directory":             schema.StringAttribute{Optional: true},
			"asset_manifest_sha256": schema.StringAttribute{Computed: true},
		},
	}
}
