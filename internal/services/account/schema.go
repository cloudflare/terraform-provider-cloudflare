// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*AccountResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "Identifier",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "type": schema.StringAttribute{
        Description: "the type of account being created. For self-serve customers, use standard. for enterprise customers, use enterprise.\nAvailable values: \"standard\", \"enterprise\".",
        Required: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("standard", "enterprise"),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "unit": schema.SingleNestedAttribute{
        Description: "information related to the tenant unit, and optionally, an id of the unit to create the account on. see https://developers.cloudflare.com/tenant/how-to/manage-accounts/",
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectType[AccountUnitModel](ctx),
        Attributes: map[string]schema.Attribute{
          "id": schema.StringAttribute{
            Description: "Tenant unit ID",
            Optional: true,
          },
        },
        PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
      },
      "name": schema.StringAttribute{
        Description: "Account name",
        Required: true,
      },
      "settings": schema.SingleNestedAttribute{
        Description: "Account settings",
        Computed: true,
        Optional: true,
        CustomType: customfield.NewNestedObjectType[AccountSettingsModel](ctx),
        Attributes: map[string]schema.Attribute{
          "abuse_contact_email": schema.StringAttribute{
            Description: "Sets an abuse contact email to notify for abuse reports.",
            Optional: true,
          },
          "default_nameservers": schema.StringAttribute{
            Description: "Specifies the default nameservers to be used for new zones added to this account.\n\n- `cloudflare.standard` for Cloudflare-branded nameservers\n- `custom.account` for account custom nameservers\n- `custom.tenant` for tenant custom nameservers\n\nSee [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)\nfor more information.\n\nDeprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).\nAvailable values: \"cloudflare.standard\", \"custom.account\", \"custom.tenant\".",
            Computed: true,
            Optional: true,
            DeprecationMessage: "This attribute is deprecated.",
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "cloudflare.standard",
              "custom.account",
              "custom.tenant",
            ),
            },
            Default: stringdefault.  StaticString("cloudflare.standard"),
          },
          "enforce_twofactor": schema.BoolAttribute{
            Description: "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled",
            Computed: true,
            Optional: true,
            Default: booldefault.  StaticBool(false),
          },
          "use_account_custom_ns_by_default": schema.BoolAttribute{
            Description: "Indicates whether new zones should use the account-level custom\nnameservers by default.\n\nDeprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).",
            Computed: true,
            Optional: true,
            DeprecationMessage: "This attribute is deprecated.",
            Default: booldefault.  StaticBool(false),
          },
        },
      },
      "created_on": schema.StringAttribute{
        Description: "Timestamp for the creation of the account",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (r *AccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *AccountResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
