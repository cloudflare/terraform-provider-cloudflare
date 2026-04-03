// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketResource)(nil)

// locationIgnorePlanModifier normalizes location values to lowercase
func locationIgnorePlanModifier() planmodifier.String {
	return &locationNormalizer{}
}

type locationNormalizer struct{}

func (m *locationNormalizer) Description(ctx context.Context) string {
	return "Ignores changes to location"
}

func (m *locationNormalizer) MarkdownDescription(ctx context.Context) string {
	return "Ignores changes to location"
}

func (m *locationNormalizer) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// If state exists, preserve state value
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
		return
	}

	// Otherwise, use the config value as-is (no normalization)
	resp.PlanValue = req.ConfigValue
}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Name of the bucket.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the bucket.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"location": schema.StringAttribute{
				Description: "Location of the bucket.\nAvailable values: \"apac\", \"eeur\", \"enam\", \"weur\", \"wnam\", \"oc\".  Note: `location` is only honored the first time a bucket with a given name is created. If you delete and recreate a bucket with the same name, the original bucket location will be used. It is also a best-effort, not a guarantee, of bucket location.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"apac",
						"eeur",
						"enam",
						"weur",
						"wnam",
						"oc",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					locationIgnorePlanModifier(),
				},
			},
			"jurisdiction": schema.StringAttribute{
				Description: "Jurisdiction where objects in this bucket are guaranteed to be stored.\nAvailable values: \"default\", \"eu\", \"fedramp\".",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"default",
						"eu",
						"fedramp",
					),
				},
			},
			"storage_class": schema.StringAttribute{
				Description: "Storage class for newly uploaded objects, unless specified otherwise.\nAvailable values: \"Standard\", \"InfrequentAccess\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Standard", "InfrequentAccess"),
				},
				Default: stringdefault.StaticString("Standard"),
			},
			"creation_date": schema.StringAttribute{
				Description:   "Creation timestamp.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *R2BucketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
