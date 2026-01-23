// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*APITokenResource)(nil)

func (r *APITokenResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   priorSchemaV0(),
			StateUpgrader: upgradeAPITokenStateV0toV1,
		},
		1:   {},
		130: {},
	}
}

// priorSchemaV0 returns the schema for version 0 (before the resources field change)
func priorSchemaV0() *schema.Schema {
	return &schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of this resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"issued_on": schema.StringAttribute{
				Description: "Timestamp of when the token was issued.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_used_on": schema.StringAttribute{
				Description: "Last time the token was used.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "Timestamp of when the token was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Token name.",
				Required:    true,
			},
			"policies": schema.ListNestedAttribute{
				Description: "List of access policies assigned to the token.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Policy identifier.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"effect": schema.StringAttribute{
							Description: "Allow or deny operations against the resources.",
							Required:    true,
						},
						"permission_groups": schema.ListNestedAttribute{
							Description: "List of permission groups assigned to this policy.",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier of the permission group.",
										Required:    true,
									},
									"meta": schema.SingleNestedAttribute{
										Description: "Attributes associated to the permission group.",
										Optional:    true,
										Computed:    true,
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Optional: true,
											},
											"value": schema.StringAttribute{
												Optional: true,
											},
										},
									},
									"name": schema.StringAttribute{
										Description: "Name of the permission group.",
										Computed:    true,
									},
								},
							},
						},
						// OLD: resources as a map
						"resources": schema.MapAttribute{
							Description: "A list of resource names that the policy applies to.",
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Description: "Status of the token.",
				Optional:    true,
			},
			"value": schema.StringAttribute{
				Description: "The token value.",
				Computed:    true,
				Sensitive:   true,
			},
			"not_before": schema.StringAttribute{
				Description: "The time before which the token MUST NOT be accepted for processing.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"expires_on": schema.StringAttribute{
				Description: "The expiration time on or after which the token MUST NOT be accepted for processing.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"condition": schema.SingleNestedAttribute{
				Description: "Conditions under which the token should be considered valid.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"request_ip": schema.SingleNestedAttribute{
						Description: "Client IP restrictions.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"in": schema.ListAttribute{
								Description: "List of IP addresses or CIDR notation the token can use.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"not_in": schema.ListAttribute{
								Description: "List of IP addresses or CIDR notation the token cannot use.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
		},
	}
}

// upgradeAPITokenStateV0toV1 upgrades the state from version 0 to version 1
// This converts the resources field from a map to a JSON-encoded string and removes obsolete fields
func upgradeAPITokenStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Define the old state structure
	type metaV0 struct {
		Key   types.String `tfsdk:"key"`
		Value types.String `tfsdk:"value"`
	}

	type permissionGroupV0 struct {
		ID   types.String `tfsdk:"id"`
		Meta *metaV0      `tfsdk:"meta"`
		Name types.String `tfsdk:"name"`
	}

	type policyV0 struct {
		ID               types.String            `tfsdk:"id"`
		Effect           types.String            `tfsdk:"effect"`
		PermissionGroups []permissionGroupV0     `tfsdk:"permission_groups"`
		Resources        map[string]types.String `tfsdk:"resources"` // OLD: map format
	}

	type conditionRequestIPV0 struct {
		In    []types.String `tfsdk:"in"`
		NotIn []types.String `tfsdk:"not_in"`
	}

	type conditionV0 struct {
		RequestIP *conditionRequestIPV0 `tfsdk:"request_ip"`
	}

	type modelV0 struct {
		ID         types.String      `tfsdk:"id"`
		IssuedOn   timetypes.RFC3339 `tfsdk:"issued_on"`
		LastUsedOn timetypes.RFC3339 `tfsdk:"last_used_on"`
		ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
		Name       types.String      `tfsdk:"name"`
		Policies   []policyV0        `tfsdk:"policies"`
		Status     types.String      `tfsdk:"status"`
		Value      types.String      `tfsdk:"value"`
		NotBefore  timetypes.RFC3339 `tfsdk:"not_before"`
		ExpiresOn  timetypes.RFC3339 `tfsdk:"expires_on"`
		Condition  *conditionV0      `tfsdk:"condition"`
	}

	// Get the old state
	var oldState modelV0
	diags := req.State.Get(ctx, &oldState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the new state structure
	type permissionGroupV1 struct {
		ID types.String `tfsdk:"id"`
	}

	type policyV1 struct {
		Effect           types.String        `tfsdk:"effect"`
		PermissionGroups []permissionGroupV1 `tfsdk:"permission_groups"`
		Resources        types.String        `tfsdk:"resources"` // NEW: JSON string format
	}

	type modelV1 struct {
		ID         types.String      `tfsdk:"id"`
		IssuedOn   timetypes.RFC3339 `tfsdk:"issued_on"`
		LastUsedOn timetypes.RFC3339 `tfsdk:"last_used_on"`
		ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
		Name       types.String      `tfsdk:"name"`
		Policies   []policyV1        `tfsdk:"policies"`
		Status     types.String      `tfsdk:"status"`
		Value      types.String      `tfsdk:"value"`
		NotBefore  timetypes.RFC3339 `tfsdk:"not_before"`
		ExpiresOn  timetypes.RFC3339 `tfsdk:"expires_on"`
		Condition  *conditionV0      `tfsdk:"condition"`
	}

	// Convert policies
	upgradedPolicies := make([]policyV1, 0, len(oldState.Policies))
	for _, oldPolicy := range oldState.Policies {
		// Convert resources map to JSON string
		resourcesMap := make(map[string]string)
		for k, v := range oldPolicy.Resources {
			if !v.IsNull() && !v.IsUnknown() {
				resourcesMap[k] = v.ValueString()
			}
		}

		resourcesJSON, err := json.Marshal(resourcesMap)
		if err != nil {
			resp.Diagnostics.AddError(
				"State Upgrade Error",
				fmt.Sprintf("Failed to convert resources to JSON: %v", err),
			)
			return
		}

		// Convert permission groups (removing computed fields)
		upgradedPermGroups := make([]permissionGroupV1, 0, len(oldPolicy.PermissionGroups))
		for _, oldPG := range oldPolicy.PermissionGroups {
			upgradedPermGroups = append(upgradedPermGroups, permissionGroupV1{
				ID: oldPG.ID,
			})
		}

		upgradedPolicy := policyV1{
			Effect:           oldPolicy.Effect,
			PermissionGroups: upgradedPermGroups,
			Resources:        types.StringValue(string(resourcesJSON)),
		}
		// Note: policy.ID is removed as it's no longer in the schema

		upgradedPolicies = append(upgradedPolicies, upgradedPolicy)
	}

	// Create the upgraded state
	upgradedState := modelV1{
		ID:         oldState.ID,
		IssuedOn:   oldState.IssuedOn,
		LastUsedOn: oldState.LastUsedOn,
		ModifiedOn: oldState.ModifiedOn,
		Name:       oldState.Name,
		Policies:   upgradedPolicies,
		Status:     oldState.Status,
		Value:      oldState.Value,
		NotBefore:  oldState.NotBefore,
		ExpiresOn:  oldState.ExpiresOn,
		Condition:  oldState.Condition,
	}

	// Set the upgraded state
	diags = resp.State.Set(ctx, upgradedState)
	resp.Diagnostics.Append(diags...)
}
