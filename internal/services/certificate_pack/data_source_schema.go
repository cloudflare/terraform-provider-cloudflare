// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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

var _ datasource.DataSourceWithConfigValidators = (*CertificatePackDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"certificate_pack_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"certificate_authority": schema.StringAttribute{
				Description: "Certificate Authority selected for the order.  For information on any certificate authority specific details or restrictions [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)\nAvailable values: \"google\", \"lets_encrypt\", \"ssl_com\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"google",
						"lets_encrypt",
						"ssl_com",
					),
				},
			},
			"cloudflare_branding": schema.BoolAttribute{
				Description: "Whether or not to add Cloudflare Branding for the order.  This will add a subdomain of sni.cloudflaressl.com as the Common Name if set to true.",
				Computed:    true,
			},
			"primary_certificate": schema.StringAttribute{
				Description: "Identifier of the primary certificate in a pack.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of certificate pack.\nAvailable values: \"initializing\", \"pending_validation\", \"deleted\", \"pending_issuance\", \"pending_deployment\", \"pending_deletion\", \"pending_expiration\", \"expired\", \"active\", \"initializing_timed_out\", \"validation_timed_out\", \"issuance_timed_out\", \"deployment_timed_out\", \"deletion_timed_out\", \"pending_cleanup\", \"staging_deployment\", \"staging_active\", \"deactivating\", \"inactive\", \"backup_issued\", \"holding_deployment\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending_validation",
						"deleted",
						"pending_issuance",
						"pending_deployment",
						"pending_deletion",
						"pending_expiration",
						"expired",
						"active",
						"initializing_timed_out",
						"validation_timed_out",
						"issuance_timed_out",
						"deployment_timed_out",
						"deletion_timed_out",
						"pending_cleanup",
						"staging_deployment",
						"staging_active",
						"deactivating",
						"inactive",
						"backup_issued",
						"holding_deployment",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "Type of certificate pack.\nAvailable values: \"mh_custom\", \"managed_hostname\", \"sni_custom\", \"universal\", \"advanced\", \"total_tls\", \"keyless\", \"legacy_custom\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"mh_custom",
						"managed_hostname",
						"sni_custom",
						"universal",
						"advanced",
						"total_tls",
						"keyless",
						"legacy_custom",
					),
				},
			},
			"validation_method": schema.StringAttribute{
				Description: "Validation Method selected for the order.\nAvailable values: \"txt\", \"http\", \"email\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"txt",
						"http",
						"email",
					),
				},
			},
			"validity_days": schema.Int64Attribute{
				Description: "Validity Days selected for the order.\nAvailable values: 14, 30, 90, 365.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.OneOf(
						14,
						30,
						90,
						365,
					),
				},
			},
			"hosts": schema.ListAttribute{
				Description: "Comma separated list of valid host names for the certificate packs. Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"certificates": schema.ListNestedAttribute{
				Description: "Array of certificates in this pack.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CertificatePackCertificatesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Certificate identifier.",
							Computed:    true,
						},
						"hosts": schema.ListAttribute{
							Description: "Hostnames covered by this certificate.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"status": schema.StringAttribute{
							Description: "Certificate status.",
							Computed:    true,
						},
						"bundle_method": schema.StringAttribute{
							Description: "Certificate bundle method.",
							Computed:    true,
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate from the authority expires.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"geo_restrictions": schema.SingleNestedAttribute{
							Description: "Specify the region where your private key can be held locally.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[CertificatePackCertificatesGeoRestrictionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"label": schema.StringAttribute{
									Description: `Available values: "us", "eu", "highest_security".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"us",
											"eu",
											"highest_security",
										),
									},
								},
							},
						},
						"issuer": schema.StringAttribute{
							Description: "The certificate authority that issued the certificate.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the certificate was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"priority": schema.Float64Attribute{
							Description: "The order/priority in which the certificate will be used.",
							Computed:    true,
						},
						"signature": schema.StringAttribute{
							Description: "The type of hash used for the certificate.",
							Computed:    true,
						},
						"uploaded_on": schema.StringAttribute{
							Description: "When the certificate was uploaded to Cloudflare.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"zone_id": schema.StringAttribute{
							Description: "Identifier.",
							Computed:    true,
						},
					},
				},
			},
			"validation_errors": schema.ListNestedAttribute{
				Description: "Domain validation errors that have been received by the certificate authority (CA).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CertificatePackValidationErrorsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"message": schema.StringAttribute{
							Description: "A domain validation error.",
							Computed:    true,
						},
					},
				},
			},
			"validation_records": schema.ListNestedAttribute{
				Description: "Certificates' validation records.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CertificatePackValidationRecordsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"emails": schema.ListAttribute{
							Description: "The set of email addresses that the certificate authority (CA) will use to complete domain validation.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"http_body": schema.StringAttribute{
							Description: "The content that the certificate authority (CA) will expect to find at the http_url during the domain validation.",
							Computed:    true,
						},
						"http_url": schema.StringAttribute{
							Description: "The url that will be checked during domain validation.",
							Computed:    true,
						},
						"txt_name": schema.StringAttribute{
							Description: "The hostname that the certificate authority (CA) will check for a TXT record during domain validation .",
							Computed:    true,
						},
						"txt_value": schema.StringAttribute{
							Description: "The TXT record that the certificate authority (CA) will check during domain validation.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Description: "Include Certificate Packs of all statuses, not just active ones.\nAvailable values: \"all\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("all"),
						},
					},
				},
			},
		},
	}
}

func (d *CertificatePackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CertificatePackDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("certificate_pack_id"), path.MatchRoot("filter")),
	}
}
