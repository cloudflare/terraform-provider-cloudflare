// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r APIShieldSchemaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"schema_id": schema.StringAttribute{
				Optional: true,
			},
			"file": schema.StringAttribute{
				Description: "Schema file bytes",
				Required:    true,
			},
			"kind": schema.StringAttribute{
				Description: "Kind of schema",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("openapi_v3"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the schema",
				Optional:    true,
			},
			"validation_enabled": schema.StringAttribute{
				Description: "Flag whether schema is enabled for validation.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("true", "false"),
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"source": schema.StringAttribute{
				Description: "Source of the schema",
				Computed:    true,
			},
		},
	}
}
