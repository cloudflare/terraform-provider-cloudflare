// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*RegistrarDomainsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[RegistrarDomainsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Domain identifier.",
							Computed:    true,
						},
						"available": schema.BoolAttribute{
							Description: "Shows if a domain is available for transferring into Cloudflare Registrar.",
							Computed:    true,
						},
						"can_register": schema.BoolAttribute{
							Description: "Indicates if the domain can be registered as a new domain.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Shows time of creation.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"current_registrar": schema.StringAttribute{
							Description: "Shows name of current registrar.",
							Computed:    true,
						},
						"expires_at": schema.StringAttribute{
							Description: "Shows when domain name registration expires.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"locked": schema.BoolAttribute{
							Description: "Shows whether a registrar lock is in place for a domain.",
							Computed:    true,
						},
						"registrant_contact": schema.SingleNestedAttribute{
							Description: "Shows contact information for domain registrant.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RegistrarDomainsRegistrantContactDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description: "Address.",
									Computed:    true,
								},
								"city": schema.StringAttribute{
									Description: "City.",
									Computed:    true,
								},
								"country": schema.StringAttribute{
									Description: "The country in which the user lives.",
									Computed:    true,
								},
								"first_name": schema.StringAttribute{
									Description: "User's first name",
									Computed:    true,
								},
								"last_name": schema.StringAttribute{
									Description: "User's last name",
									Computed:    true,
								},
								"organization": schema.StringAttribute{
									Description: "Name of organization.",
									Computed:    true,
								},
								"phone": schema.StringAttribute{
									Description: "User's telephone number",
									Computed:    true,
								},
								"state": schema.StringAttribute{
									Description: "State.",
									Computed:    true,
								},
								"zip": schema.StringAttribute{
									Description: "The zipcode or postal code where the user lives.",
									Computed:    true,
								},
								"id": schema.StringAttribute{
									Description: "Contact Identifier.",
									Computed:    true,
								},
								"address2": schema.StringAttribute{
									Description: "Optional address line for unit, floor, suite, etc.",
									Computed:    true,
								},
								"email": schema.StringAttribute{
									Description: "The contact email address of the user.",
									Computed:    true,
								},
								"fax": schema.StringAttribute{
									Description: "Contact fax number.",
									Computed:    true,
								},
							},
						},
						"registry_statuses": schema.StringAttribute{
							Description: "A comma-separated list of registry status codes. A full list of status codes can be found at [EPP Status Codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).",
							Computed:    true,
						},
						"supported_tld": schema.BoolAttribute{
							Description: "Whether a particular TLD is currently supported by Cloudflare Registrar. Refer to [TLD Policies](https://www.cloudflare.com/tld-policies/) for a list of supported TLDs.",
							Computed:    true,
						},
						"transfer_in": schema.SingleNestedAttribute{
							Description: "Statuses for domain transfers into Cloudflare Registrar.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[RegistrarDomainsTransferInDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"accept_foa": schema.StringAttribute{
									Description: "Form of authorization has been accepted by the registrant.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("needed", "ok"),
									},
								},
								"approve_transfer": schema.StringAttribute{
									Description: "Shows transfer status with the registry.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"needed",
											"ok",
											"pending",
											"trying",
											"rejected",
											"unknown",
										),
									},
								},
								"can_cancel_transfer": schema.BoolAttribute{
									Description: "Indicates if cancellation is still possible.",
									Computed:    true,
								},
								"disable_privacy": schema.StringAttribute{
									Description: "Privacy guards are disabled at the foreign registrar.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"needed",
											"ok",
											"unknown",
										),
									},
								},
								"enter_auth_code": schema.StringAttribute{
									Description: "Auth code has been entered and verified.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"needed",
											"ok",
											"pending",
											"trying",
											"rejected",
										),
									},
								},
								"unlock_domain": schema.StringAttribute{
									Description: "Domain is unlocked at the foreign registrar.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"needed",
											"ok",
											"pending",
											"trying",
											"unknown",
										),
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Description: "Last updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *RegistrarDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *RegistrarDomainsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
