// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomHostnamesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "direction": schema.StringAttribute{
        Description: "Direction to order hostnames.\nAvailable values: \"asc\", \"desc\".",
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("asc", "desc"),
        },
      },
      "hostname": schema.StringAttribute{
        Description: "Fully qualified domain name to match against. This parameter cannot be used with the 'id' parameter.",
        Optional: true,
      },
      "id": schema.StringAttribute{
        Description: "Hostname ID to match against. This ID was generated and returned during the initial custom_hostname creation. This parameter cannot be used with the 'hostname' parameter.",
        Optional: true,
      },
      "ssl": schema.Float64Attribute{
        Description: "Whether to filter hostnames based on if they have SSL enabled.\nAvailable values: 0, 1.",
        Optional: true,
        Validators: []validator.Float64{
        float64validator.OneOf(0, 1),
        },
      },
      "order": schema.StringAttribute{
        Description: "Field to order hostnames by.\nAvailable values: \"ssl\", \"ssl_status\".",
        Computed: true,
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("ssl", "ssl_status"),
        },
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "result": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[CustomHostnamesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Identifier",
              Computed: true,
            },
            "hostname": schema.StringAttribute{
              Description: "The custom hostname that will point to your hostname via CNAME.",
              Computed: true,
            },
            "ssl": schema.SingleNestedAttribute{
              Computed: true,
              CustomType: customfield.NewNestedObjectType[CustomHostnamesSSLDataSourceModel](ctx),
              Attributes: map[string]schema.Attribute{
                "id": schema.StringAttribute{
                  Description: "Custom hostname SSL identifier tag.",
                  Computed: true,
                },
                "bundle_method": schema.StringAttribute{
                  Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.\nAvailable values: \"ubiquitous\", \"optimal\", \"force\".",
                  Computed: true,
                  Validators: []validator.String{
                  stringvalidator.OneOfCaseInsensitive(
                    "ubiquitous",
                    "optimal",
                    "force",
                  ),
                  },
                },
                "certificate_authority": schema.StringAttribute{
                  Description: "The Certificate Authority that will issue the certificate\nAvailable values: \"digicert\", \"google\", \"lets_encrypt\", \"ssl_com\".",
                  Computed: true,
                  Validators: []validator.String{
                  stringvalidator.OneOfCaseInsensitive(
                    "digicert",
                    "google",
                    "lets_encrypt",
                    "ssl_com",
                  ),
                  },
                },
                "custom_certificate": schema.StringAttribute{
                  Description: "If a custom uploaded certificate is used.",
                  Computed: true,
                },
                "custom_csr_id": schema.StringAttribute{
                  Description: "The identifier for the Custom CSR that was used.",
                  Computed: true,
                },
                "custom_key": schema.StringAttribute{
                  Description: "The key for a custom uploaded certificate.",
                  Computed: true,
                  Sensitive: true,
                },
                "expires_on": schema.StringAttribute{
                  Description: "The time the custom certificate expires on.",
                  Computed: true,
                  CustomType: timetypes.RFC3339Type{

                  },
                },
                "hosts": schema.ListAttribute{
                  Description: "A list of Hostnames on a custom uploaded certificate.",
                  Computed: true,
                  CustomType: customfield.NewListType[types.String](ctx),
                  ElementType: types.StringType,
                },
                "issuer": schema.StringAttribute{
                  Description: "The issuer on a custom uploaded certificate.",
                  Computed: true,
                },
                "method": schema.StringAttribute{
                  Description: "Domain control validation (DCV) method used for this hostname.\nAvailable values: \"http\", \"txt\", \"email\".",
                  Computed: true,
                  Validators: []validator.String{
                  stringvalidator.OneOfCaseInsensitive(
                    "http",
                    "txt",
                    "email",
                  ),
                  },
                },
                "serial_number": schema.StringAttribute{
                  Description: "The serial number on a custom uploaded certificate.",
                  Computed: true,
                },
                "settings": schema.SingleNestedAttribute{
                  Computed: true,
                  CustomType: customfield.NewNestedObjectType[CustomHostnamesSSLSettingsDataSourceModel](ctx),
                  Attributes: map[string]schema.Attribute{
                    "ciphers": schema.ListAttribute{
                      Description: "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.",
                      Computed: true,
                      CustomType: customfield.NewListType[types.String](ctx),
                      ElementType: types.StringType,
                    },
                    "early_hints": schema.StringAttribute{
                      Description: "Whether or not Early Hints is enabled.\nAvailable values: \"on\", \"off\".",
                      Computed: true,
                      Validators: []validator.String{
                      stringvalidator.OneOfCaseInsensitive("on", "off"),
                      },
                    },
                    "http2": schema.StringAttribute{
                      Description: "Whether or not HTTP2 is enabled.\nAvailable values: \"on\", \"off\".",
                      Computed: true,
                      Validators: []validator.String{
                      stringvalidator.OneOfCaseInsensitive("on", "off"),
                      },
                    },
                    "min_tls_version": schema.StringAttribute{
                      Description: "The minimum TLS version supported.\nAvailable values: \"1.0\", \"1.1\", \"1.2\", \"1.3\".",
                      Computed: true,
                      Validators: []validator.String{
                      stringvalidator.OneOfCaseInsensitive(
                        "1.0",
                        "1.1",
                        "1.2",
                        "1.3",
                      ),
                      },
                    },
                    "tls_1_3": schema.StringAttribute{
                      Description: "Whether or not TLS 1.3 is enabled.\nAvailable values: \"on\", \"off\".",
                      Computed: true,
                      Validators: []validator.String{
                      stringvalidator.OneOfCaseInsensitive("on", "off"),
                      },
                    },
                  },
                },
                "signature": schema.StringAttribute{
                  Description: "The signature on a custom uploaded certificate.",
                  Computed: true,
                },
                "status": schema.StringAttribute{
                  Description: "Status of the hostname's SSL certificates.\nAvailable values: \"initializing\", \"pending_validation\", \"deleted\", \"pending_issuance\", \"pending_deployment\", \"pending_deletion\", \"pending_expiration\", \"expired\", \"active\", \"initializing_timed_out\", \"validation_timed_out\", \"issuance_timed_out\", \"deployment_timed_out\", \"deletion_timed_out\", \"pending_cleanup\", \"staging_deployment\", \"staging_active\", \"deactivating\", \"inactive\", \"backup_issued\", \"holding_deployment\".",
                  Computed: true,
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
                  Description: "Level of validation to be used for this hostname. Domain validation (dv) must be used.\nAvailable values: \"dv\".",
                  Computed: true,
                  Validators: []validator.String{
                  stringvalidator.OneOfCaseInsensitive("dv"),
                  },
                },
                "uploaded_on": schema.StringAttribute{
                  Description: "The time the custom certificate was uploaded.",
                  Computed: true,
                  CustomType: timetypes.RFC3339Type{

                  },
                },
                "validation_errors": schema.ListNestedAttribute{
                  Description: "Domain validation errors that have been received by the certificate authority (CA).",
                  Computed: true,
                  CustomType: customfield.NewNestedObjectListType[CustomHostnamesSSLValidationErrorsDataSourceModel](ctx),
                  NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                      "message": schema.StringAttribute{
                        Description: "A domain validation error.",
                        Computed: true,
                      },
                    },
                  },
                },
                "validation_records": schema.ListNestedAttribute{
                  Computed: true,
                  CustomType: customfield.NewNestedObjectListType[CustomHostnamesSSLValidationRecordsDataSourceModel](ctx),
                  NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                      "emails": schema.ListAttribute{
                        Description: "The set of email addresses that the certificate authority (CA) will use to complete domain validation.",
                        Computed: true,
                        CustomType: customfield.NewListType[types.String](ctx),
                        ElementType: types.StringType,
                      },
                      "http_body": schema.StringAttribute{
                        Description: "The content that the certificate authority (CA) will expect to find at the http_url during the domain validation.",
                        Computed: true,
                      },
                      "http_url": schema.StringAttribute{
                        Description: "The url that will be checked during domain validation.",
                        Computed: true,
                      },
                      "txt_name": schema.StringAttribute{
                        Description: "The hostname that the certificate authority (CA) will check for a TXT record during domain validation .",
                        Computed: true,
                      },
                      "txt_value": schema.StringAttribute{
                        Description: "The TXT record that the certificate authority (CA) will check during domain validation.",
                        Computed: true,
                      },
                    },
                  },
                },
                "wildcard": schema.BoolAttribute{
                  Description: "Indicates whether the certificate covers a wildcard.",
                  Computed: true,
                },
              },
            },
            "created_at": schema.StringAttribute{
              Description: "This is the time the hostname was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "custom_metadata": schema.MapAttribute{
              Description: "Unique key/value metadata for this hostname. These are per-hostname (customer) settings.",
              Computed: true,
              CustomType: customfield.NewMapType[types.String](ctx),
              ElementType: types.StringType,
            },
            "custom_origin_server": schema.StringAttribute{
              Description: "a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME record.",
              Computed: true,
            },
            "custom_origin_sni": schema.StringAttribute{
              Description: "A hostname that will be sent to your custom origin server as SNI for TLS handshake. This can be a valid subdomain of the zone or custom origin server name or the string ':request_host_header:' which will cause the host header in the request to be used as SNI. Not configurable with default/fallback origin server.",
              Computed: true,
            },
            "ownership_verification": schema.SingleNestedAttribute{
              Description: "This is a record which can be placed to activate a hostname.",
              Computed: true,
              CustomType: customfield.NewNestedObjectType[CustomHostnamesOwnershipVerificationDataSourceModel](ctx),
              Attributes: map[string]schema.Attribute{
                "name": schema.StringAttribute{
                  Description: "DNS Name for record.",
                  Computed: true,
                },
                "type": schema.StringAttribute{
                  Description: "DNS Record type.\nAvailable values: \"txt\".",
                  Computed: true,
                  Validators: []validator.String{
                  stringvalidator.OneOfCaseInsensitive("txt"),
                  },
                },
                "value": schema.StringAttribute{
                  Description: "Content for the record.",
                  Computed: true,
                },
              },
            },
            "ownership_verification_http": schema.SingleNestedAttribute{
              Description: "This presents the token to be served by the given http url to activate a hostname.",
              Computed: true,
              CustomType: customfield.NewNestedObjectType[CustomHostnamesOwnershipVerificationHTTPDataSourceModel](ctx),
              Attributes: map[string]schema.Attribute{
                "http_body": schema.StringAttribute{
                  Description: "Token to be served.",
                  Computed: true,
                },
                "http_url": schema.StringAttribute{
                  Description: "The HTTP URL that will be checked during custom hostname verification and where the customer should host the token.",
                  Computed: true,
                },
              },
            },
            "status": schema.StringAttribute{
              Description: "Status of the hostname's activation.\nAvailable values: \"active\", \"pending\", \"active_redeploying\", \"moved\", \"pending_deletion\", \"deleted\", \"pending_blocked\", \"pending_migration\", \"pending_provisioned\", \"test_pending\", \"test_active\", \"test_active_apex\", \"test_blocked\", \"test_failed\", \"provisioned\", \"blocked\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "active",
                "pending",
                "active_redeploying",
                "moved",
                "pending_deletion",
                "deleted",
                "pending_blocked",
                "pending_migration",
                "pending_provisioned",
                "test_pending",
                "test_active",
                "test_active_apex",
                "test_blocked",
                "test_failed",
                "provisioned",
                "blocked",
              ),
              },
            },
            "verification_errors": schema.ListAttribute{
              Description: "These are errors that were encountered while trying to activate a hostname.",
              Computed: true,
              CustomType: customfield.NewListType[types.String](ctx),
              ElementType: types.StringType,
            },
          },
        },
      },
    },
  }
}

func (d *CustomHostnamesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CustomHostnamesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
