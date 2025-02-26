// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PagesDomainResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"project_name": schema.StringAttribute{
				Description:   "Name of the project.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"certificate_authority": schema.StringAttribute{
				Description: "Available values: \"google\", \"lets_encrypt\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("google", "lets_encrypt"),
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"domain_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: "Available values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending",
						"active",
						"deactivated",
						"blocked",
						"error",
					),
				},
			},
			"zone_tag": schema.StringAttribute{
				Computed: true,
			},
			"validation_data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesDomainValidationDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
					},
					"method": schema.StringAttribute{
						Description: "Available values: \"http\", \"txt\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("http", "txt"),
						},
					},
					"status": schema.StringAttribute{
						Description: "Available values: \"initializing\", \"pending\", \"active\", \"deactivated\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"initializing",
								"pending",
								"active",
								"deactivated",
								"error",
							),
						},
					},
					"txt_name": schema.StringAttribute{
						Computed: true,
					},
					"txt_value": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"verification_data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesDomainVerificationDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"error_message": schema.StringAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Description: "Available values: \"pending\", \"active\", \"deactivated\", \"blocked\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"pending",
								"active",
								"deactivated",
								"blocked",
								"error",
							),
						},
					},
				},
			},
		},
	}
}

func (r *PagesDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PagesDomainResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
