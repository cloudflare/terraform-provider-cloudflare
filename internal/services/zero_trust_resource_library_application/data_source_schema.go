// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustResourceLibraryApplicationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"application_confidence_score": schema.Float64Attribute{
				Description: "Confidence score for the application. Returns -1 when no score is available.",
				Computed:    true,
			},
			"application_source": schema.StringAttribute{
				Description: "Returns the application source.",
				Computed:    true,
			},
			"application_type": schema.StringAttribute{
				Description: "Returns the application type.",
				Computed:    true,
			},
			"application_type_description": schema.StringAttribute{
				Description: "Returns the application type description.",
				Computed:    true,
			},
			"category_id": schema.Int64Attribute{
				Description: "Returns the category ID.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "Returns the application creation time.",
				Computed:    true,
			},
			"gen_ai_score": schema.Float64Attribute{
				Description: "GenAI score for the application. Returns -1 when no score is available.",
				Computed:    true,
			},
			"human_id": schema.StringAttribute{
				Description: "Returns the human readable ID.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Returns the application name.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Returns the application update time.",
				Computed:    true,
			},
			"version": schema.StringAttribute{
				Description: "Returns the application version.",
				Computed:    true,
			},
			"hostnames": schema.SetAttribute{
				Description: "Hostnames matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ip_subnets": schema.SetAttribute{
				Description: "IP subnets matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"port_protocols": schema.SetAttribute{
				Description: "Port and protocol pairs matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"support_domains": schema.SetAttribute{
				Description: "Support domains matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"supported": schema.SetAttribute{
				Description: "Cloudflare products that support this application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"application_score_composition": schema.StringAttribute{
				Description: "Returns the score composition breakdown for the application.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *ZeroTrustResourceLibraryApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustResourceLibraryApplicationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
