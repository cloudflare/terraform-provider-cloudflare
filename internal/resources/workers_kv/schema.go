// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r WorkersKVResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Required:    true,
			},
			"key_name": schema.StringAttribute{
				Description: "A key's name. The name may be at most 512 bytes. All printable, non-whitespace characters are valid. Use percent-encoding to define key names as part of a URL.",
				Required:    true,
			},
			"metadata": schema.StringAttribute{
				Description: "Arbitrary JSON to be associated with a key/value pair.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "A byte sequence to be stored, up to 25 MiB in length.",
				Required:    true,
			},
		},
	}
}
