// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*UserGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Settings Read",
				"Account Settings Write",
				"SCIM Provisioning",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "User Group identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the User group.",
				Required:    true,
			},
			"policies": schema.ListNestedAttribute{
				Description: "Policies attached to the User group",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access": schema.StringAttribute{
							Description: "Allow or deny operations against the resources.\nAvailable values: \"allow\", \"deny\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("allow", "deny"),
							},
						},
						"permission_groups": schema.ListNestedAttribute{
							Description: "A set of permission groups that are specified to the policy.",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Permission Group identifier tag.",
										Required:    true,
									},
								},
							},
						},
						"resource_groups": schema.ListNestedAttribute{
							Description: "A set of resource groups that are specified to the policy.",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Resource Group identifier tag.",
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "Timestamp for the creation of the user group",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "Last time the user group was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *UserGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
