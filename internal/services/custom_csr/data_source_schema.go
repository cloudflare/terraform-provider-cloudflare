// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_csr

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomCsrDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account: SSL and Certificates Read",
				"Account: SSL and Certificates Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Custom CSR identifier tag.",
				Computed:    true,
			},
			"custom_csr_id": schema.StringAttribute{
				Description: "Custom CSR identifier tag.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"account_tag": schema.StringAttribute{
				Description: "Account identifier associated with this CSR.",
				Computed:    true,
			},
			"common_name": schema.StringAttribute{
				Description: "The common name (domain) for the CSR.",
				Computed:    true,
			},
			"country": schema.StringAttribute{
				Description: "Two-letter ISO 3166-1 alpha-2 country code.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the CSR was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"csr": schema.StringAttribute{
				Description: "The PEM-encoded Certificate Signing Request.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description for the CSR.",
				Computed:    true,
			},
			"key_type": schema.StringAttribute{
				Description: "The key algorithm used to generate the CSR.\nAvailable values: \"rsa2048\", \"p256v1\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("rsa2048", "p256v1"),
				},
			},
			"locality": schema.StringAttribute{
				Description: "City or locality name.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Human-readable name for the CSR.",
				Computed:    true,
			},
			"organization": schema.StringAttribute{
				Description: "Organization name.",
				Computed:    true,
			},
			"organizational_unit": schema.StringAttribute{
				Description: "Organizational unit name.",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "State or province name.",
				Computed:    true,
			},
			"sans": schema.ListAttribute{
				Description: "Subject Alternative Names included in the CSR.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *CustomCsrDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomCsrDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
