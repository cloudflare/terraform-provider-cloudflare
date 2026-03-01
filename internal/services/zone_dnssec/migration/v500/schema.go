// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareZoneDNSSECSchema returns the schema for the v4 (SDKv2) resource
// This schema is used to read the old state format before migration
// Version 0 corresponds to the v4 SDKv2 schema
func SourceCloudflareZoneDNSSECSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 SDKv2 resources start at version 0
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed: true, // Computed-only in v4
			},
			"flags": schema.Int64Attribute{
				Computed:      true, // Int64 in v4
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"algorithm": schema.StringAttribute{
				Computed: true,
			},
			"key_type": schema.StringAttribute{
				Computed: true,
			},
			"digest_type": schema.StringAttribute{
				Computed: true,
			},
			"digest_algorithm": schema.StringAttribute{
				Computed: true,
			},
			"digest": schema.StringAttribute{
				Computed: true,
			},
			"ds": schema.StringAttribute{
				Computed: true,
			},
			"key_tag": schema.Int64Attribute{
				Computed:      true, // Int64 in v4
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"public_key": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Optional: true, // Optional+Computed in v4
				Computed: true,
			},
		},
	}
}

// GetAttributeTypes returns the attribute types for the source model
// This is used internally by the framework for state parsing
func GetAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":               types.StringType,
		"zone_id":          types.StringType,
		"status":           types.StringType,
		"flags":            types.Int64Type,
		"algorithm":        types.StringType,
		"key_type":         types.StringType,
		"digest_type":      types.StringType,
		"digest_algorithm": types.StringType,
		"digest":           types.StringType,
		"ds":               types.StringType,
		"key_tag":          types.Int64Type,
		"public_key":       types.StringType,
		"modified_on":      types.StringType,
	}
}
