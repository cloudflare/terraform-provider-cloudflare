// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketResource)(nil)

// locationNormalizePlanModifier normalizes location values to lowercase
func locationNormalizePlanModifier() planmodifier.String {
	return &locationNormalizer{}
}

type locationNormalizer struct{}

func (m *locationNormalizer) Description(ctx context.Context) string {
	return "Normalizes location to lowercase"
}

func (m *locationNormalizer) MarkdownDescription(ctx context.Context) string {
	return "Normalizes location to lowercase"
}

func (m *locationNormalizer) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	
	// If state exists and values are case-insensitively equal, preserve state value
	if !req.StateValue.IsNull() {
		configVal := req.ConfigValue.ValueString()
		stateVal := req.StateValue.ValueString()
		
		// If they're case-insensitively equal but different case, preserve state
		if strings.EqualFold(configVal, stateVal) && configVal != stateVal {
			resp.PlanValue = req.StateValue
			return
		}
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
				Description: "Location of the bucket.\nAvailable values: \"apac\", \"eeur\", \"enam\", \"weur\", \"wnam\", \"oc\".",
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
					locationNormalizePlanModifier(),
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
