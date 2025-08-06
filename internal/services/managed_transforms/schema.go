// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ManagedTransformsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"managed_request_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
						"has_conflict": schema.BoolAttribute{
							Description: "Whether the Managed Transform conflicts with the currently-enabled Managed Transforms.",
							Computed:    true,
						},
						"conflicts_with": schema.ListAttribute{
							Description: "The Managed Transforms that this Managed Transform conflicts with.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
				PlanModifiers: []planmodifier.Set{newMergeDefaultStateValues()},
			},
			"managed_response_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
						"has_conflict": schema.BoolAttribute{
							Description: "Whether the Managed Transform conflicts with the currently-enabled Managed Transforms.",
							Computed:    true,
						},
						"conflicts_with": schema.ListAttribute{
							Description: "The Managed Transforms that this Managed Transform conflicts with.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
				PlanModifiers: []planmodifier.Set{newMergeDefaultStateValues()},
			},
		},
	}
}

func (r *ManagedTransformsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ManagedTransformsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}


type mergeDefaultStateValues struct{}

// WIP! explain
func newMergeDefaultStateValues() planmodifier.Set {
	return mergeDefaultStateValues{}
}

func (m mergeDefaultStateValues) Description(_ context.Context) string {
	return "WIP!"
}

func (m mergeDefaultStateValues) MarkdownDescription(_ context.Context) string {
	return "WIP!"
}

func (m mergeDefaultStateValues) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	getId := func (v attr.Value) types.String {
		objValue, ok := v.(types.Object)

		if !ok {
			panic("unexpected value type")
		}

		id, ok := objValue.Attributes()["id"].(types.String)

		if !ok {
			panic("unexpected value type")
		}

		return id
	}
	getEnabled := func (v attr.Value) types.Bool {
		objValue, ok := v.(types.Object)

		if !ok {
			panic("unexpected value type")
		}

		id, ok := objValue.Attributes()["enabled"].(types.Bool)

		if !ok {
			panic("unexpected value type")
		}

		return id
	}

	if req.StateValue.IsNull() {
		return
	}

	fmt.Printf("\nstate\n")
	for _, r := range req.StateValue.Elements() {
		fmt.Printf("    %v\n", r)
	}

	fmt.Printf("\n\nplan\n")
	for _, r := range req.PlanValue.Elements() {
		fmt.Printf("    %v\n", r)
	}

	fmt.Printf("\nmerging...\n")

	planTransformationIds := make(map[string]bool)

	for _, transformation := range req.PlanValue.Elements() {
		id := getId(transformation)

		if id.IsNull() || id.IsUnknown() {
			continue
		}

		planTransformationIds[id.ValueString()] = true
	}

	newPlanElements := req.PlanValue.Elements()

	for _, transformation := range req.StateValue.Elements() {
		id := getId(transformation)

		if id.IsNull() || id.IsUnknown() {
			continue
		}

		enabled := getEnabled(transformation)

		if enabled.IsNull() || enabled.IsUnknown() {
			continue
		}

		transformationInDefaultState := !enabled.ValueBool()

		if !planTransformationIds[id.ValueString()] && transformationInDefaultState {
			newPlanElements = append(newPlanElements, transformation)

			fmt.Printf("    adding %v\n", transformation)
		}
	}

	newPlanValue, _ := types.SetValue(req.PlanValue.ElementType(ctx), newPlanElements)

	resp.PlanValue = newPlanValue

	fmt.Printf("\n\nnew plan\n")
	for _, r := range resp.PlanValue.Elements() {
		fmt.Printf("    %v\n", r)
	}

	fmt.Printf("\n================================================================\n\n")
}
