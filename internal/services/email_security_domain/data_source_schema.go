// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailSecurityDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloud Email Security: Read",
				"Cloud Email Security: Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Domain identifier.",
				Computed:    true,
			},
			"domain_id": schema.StringAttribute{
				Description: "Domain identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"dmarc_status": schema.StringAttribute{
				Description: `Available values: "none", "good", "invalid".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"good",
						"invalid",
					),
				},
			},
			"domain": schema.StringAttribute{
				Computed: true,
			},
			"folder": schema.StringAttribute{
				Description: `Available values: "AllItems", "Inbox".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("AllItems", "Inbox"),
				},
			},
			"inbox_provider": schema.StringAttribute{
				Description: `Available values: "Microsoft", "Google".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Microsoft", "Google"),
				},
			},
			"integration_id": schema.StringAttribute{
				Computed: true,
			},
			"last_modified": schema.StringAttribute{
				Description:        "Deprecated, use `modified_at` instead. End of life: November 1, 2026.",
				Computed:           true,
				DeprecationMessage: "Use `modified_at` instead.",
				CustomType:         timetypes.RFC3339Type{},
			},
			"lookback_hops": schema.Int64Attribute{
				Computed: true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"o365_tenant_id": schema.StringAttribute{
				Computed: true,
			},
			"require_tls_inbound": schema.BoolAttribute{
				Computed: true,
			},
			"require_tls_outbound": schema.BoolAttribute{
				Computed: true,
			},
			"spf_status": schema.StringAttribute{
				Description: `Available values: "none", "good", "neutral", "open", "invalid".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"good",
						"neutral",
						"open",
						"invalid",
					),
				},
			},
			"status": schema.StringAttribute{
				Description: `Available values: "PENDING", "ACTIVE", "FAILED", "TIMEOUT".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"ACTIVE",
						"FAILED",
						"TIMEOUT",
					),
				},
			},
			"transport": schema.StringAttribute{
				Computed: true,
			},
			"allowed_delivery_modes": schema.SetAttribute{
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"drop_dispositions": schema.SetAttribute{
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ip_restrictions": schema.SetAttribute{
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"regions": schema.SetAttribute{
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"authorization": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[EmailSecurityDomainAuthorizationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authorized": schema.BoolAttribute{
						Computed: true,
					},
					"timestamp": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"status_message": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"emails_processed": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[EmailSecurityDomainEmailsProcessedDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"timestamp": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"total_emails_processed": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"total_emails_processed_previous": schema.Int64Attribute{
						Computed: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active_delivery_mode": schema.StringAttribute{
						Description: "Currently active delivery mode to filter by.\nAvailable values: \"DIRECT\", \"BCC\", \"JOURNAL\", \"API\", \"RETRO_SCAN\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"DIRECT",
								"BCC",
								"JOURNAL",
								"API",
								"RETRO_SCAN",
							),
						},
					},
					"allowed_delivery_mode": schema.StringAttribute{
						Description: "Delivery mode to filter by.\nAvailable values: \"DIRECT\", \"BCC\", \"JOURNAL\", \"API\", \"RETRO_SCAN\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"DIRECT",
								"BCC",
								"JOURNAL",
								"API",
								"RETRO_SCAN",
							),
						},
					},
					"direction": schema.StringAttribute{
						Description: "The sorting direction.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"domain": schema.ListAttribute{
						Description: "Domain names to filter by.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"integration_id": schema.StringAttribute{
						Description: "Integration ID to filter by.",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to sort by.\nAvailable values: \"domain\", \"created_at\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("domain", "created_at"),
						},
					},
					"search": schema.StringAttribute{
						Description: "Search term for filtering records. Behavior may change.",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Filters response to domains with the provided status.\nAvailable values: \"PENDING\", \"ACTIVE\", \"FAILED\", \"TIMEOUT\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"PENDING",
								"ACTIVE",
								"FAILED",
								"TIMEOUT",
							),
						},
					},
				},
			},
		},
	}
}

func (d *EmailSecurityDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailSecurityDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("domain_id"), path.MatchRoot("filter")),
	}
}
