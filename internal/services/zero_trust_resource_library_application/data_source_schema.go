// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustResourceLibraryApplicationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required: true,
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
			"intel_id": schema.Int64Attribute{
				Description: "Returns the Intel API ID for the application.",
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
			"hostnames": schema.ListAttribute{
				Description: "Returns the list of hostnames for the application.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ip_subnets": schema.ListAttribute{
				Description: "Returns the list of IP subnets for the application.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"port_protocols": schema.ListAttribute{
				Description: "Returns the list of port protocols for the application.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"support_domains": schema.ListAttribute{
				Description: "Returns the list of support domains for the application.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
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
