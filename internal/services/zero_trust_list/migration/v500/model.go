package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceTeamsListModel represents the v4 cloudflare_teams_list state (schema_version=0).
//
// Key differences from v5:
// - "items" is a list of strings (not a set of objects)
// - "items_with_description" exists as a separate list of objects
// - In v5, both are merged into unified "items" set of {value, description} objects
type SourceTeamsListModel struct {
	ID                   types.String                        `tfsdk:"id"`
	AccountID            types.String                        `tfsdk:"account_id"`
	Type                 types.String                        `tfsdk:"type"`
	Name                 types.String                        `tfsdk:"name"`
	Description          types.String                        `tfsdk:"description"`
	Items                []types.String                      `tfsdk:"items"`
	ItemsWithDescription []SourceItemsWithDescriptionModel   `tfsdk:"items_with_description"`
}

type SourceItemsWithDescriptionModel struct {
	Value       types.String `tfsdk:"value"`
	Description types.String `tfsdk:"description"`
}

// TargetZeroTrustListModel represents the v5 cloudflare_zero_trust_list state.
// Must include ALL fields from the current schema for resp.State.Set() to work.
type TargetZeroTrustListModel struct {
	ID          types.String              `tfsdk:"id"`
	AccountID   types.String              `tfsdk:"account_id"`
	Type        types.String              `tfsdk:"type"`
	Name        types.String              `tfsdk:"name"`
	Items       *[]*TargetItemModel       `tfsdk:"items"`
	Description types.String              `tfsdk:"description"`
	CreatedAt   timetypes.RFC3339         `tfsdk:"created_at"`
	ListCount   types.Float64             `tfsdk:"list_count"`
	UpdatedAt   timetypes.RFC3339         `tfsdk:"updated_at"`
}

type TargetItemModel struct {
	Value       types.String `tfsdk:"value"`
	Description types.String `tfsdk:"description"`
}
