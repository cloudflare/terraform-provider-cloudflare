// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/resource_sharing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SharesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SharesResultDataSourceModel] `json:"result,computed"`
}

type SharesDataSourceModel struct {
	AccountID              types.String                                              `tfsdk:"account_id" path:"account_id,required"`
	IncludeRecipientCounts types.Bool                                                `tfsdk:"include_recipient_counts" query:"include_recipient_counts,optional"`
	IncludeResources       types.Bool                                                `tfsdk:"include_resources" query:"include_resources,optional"`
	Kind                   types.String                                              `tfsdk:"kind" query:"kind,optional"`
	Status                 types.String                                              `tfsdk:"status" query:"status,optional"`
	TargetType             types.String                                              `tfsdk:"target_type" query:"target_type,optional"`
	ResourceTypes          *[]types.String                                           `tfsdk:"resource_types" query:"resource_types,optional"`
	Tag                    *[]types.String                                           `tfsdk:"tag" query:"tag,optional"`
	Direction              types.String                                              `tfsdk:"direction" query:"direction,computed_optional"`
	Order                  types.String                                              `tfsdk:"order" query:"order,computed_optional"`
	MaxItems               types.Int64                                               `tfsdk:"max_items"`
	Result                 customfield.NestedObjectList[SharesResultDataSourceModel] `tfsdk:"result"`
}

func (m *SharesDataSourceModel) toListParams(_ context.Context) (params resource_sharing.ResourceSharingListParams, diags diag.Diagnostics) {
	mResourceTypes := []resource_sharing.ResourceSharingListParamsResourceType{}
	if m.ResourceTypes != nil {
		for _, item := range *m.ResourceTypes {
			mResourceTypes = append(mResourceTypes, resource_sharing.ResourceSharingListParamsResourceType(item.ValueString()))
		}
	}
	mTag := []string{}
	if m.Tag != nil {
		for _, item := range *m.Tag {
			mTag = append(mTag, item.ValueString())
		}
	}

	params = resource_sharing.ResourceSharingListParams{
		AccountID:     cloudflare.F(m.AccountID.ValueString()),
		ResourceTypes: cloudflare.F(mResourceTypes),
		Tag:           cloudflare.F(mTag),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(resource_sharing.ResourceSharingListParamsDirection(m.Direction.ValueString()))
	}
	if !m.IncludeRecipientCounts.IsNull() {
		params.IncludeRecipientCounts = cloudflare.F(m.IncludeRecipientCounts.ValueBool())
	}
	if !m.IncludeResources.IsNull() {
		params.IncludeResources = cloudflare.F(m.IncludeResources.ValueBool())
	}
	if !m.Kind.IsNull() {
		params.Kind = cloudflare.F(resource_sharing.ResourceSharingListParamsKind(m.Kind.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(resource_sharing.ResourceSharingListParamsOrder(m.Order.ValueString()))
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(resource_sharing.ResourceSharingListParamsStatus(m.Status.ValueString()))
	}
	if !m.TargetType.IsNull() {
		params.TargetType = cloudflare.F(resource_sharing.ResourceSharingListParamsTargetType(m.TargetType.ValueString()))
	}

	return
}

type SharesResultDataSourceModel struct {
	ID                           types.String                                                 `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String                                                 `tfsdk:"account_id" json:"account_id,computed"`
	AccountName                  types.String                                                 `tfsdk:"account_name" json:"account_name,computed"`
	Created                      timetypes.RFC3339                                            `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified                     timetypes.RFC3339                                            `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name                         types.String                                                 `tfsdk:"name" json:"name,computed"`
	OrganizationID               types.String                                                 `tfsdk:"organization_id" json:"organization_id,computed"`
	Status                       types.String                                                 `tfsdk:"status" json:"status,computed"`
	TargetType                   types.String                                                 `tfsdk:"target_type" json:"target_type,computed"`
	AssociatedRecipientCount     types.Int64                                                  `tfsdk:"associated_recipient_count" json:"associated_recipient_count,computed"`
	AssociatingRecipientCount    types.Int64                                                  `tfsdk:"associating_recipient_count" json:"associating_recipient_count,computed"`
	DisassociatedRecipientCount  types.Int64                                                  `tfsdk:"disassociated_recipient_count" json:"disassociated_recipient_count,computed"`
	DisassociatingRecipientCount types.Int64                                                  `tfsdk:"disassociating_recipient_count" json:"disassociating_recipient_count,computed"`
	Kind                         types.String                                                 `tfsdk:"kind" json:"kind,computed"`
	Resources                    customfield.NestedObjectList[SharesResourcesDataSourceModel] `tfsdk:"resources" json:"resources,computed"`
}

type SharesResourcesDataSourceModel struct {
	ID                types.String         `tfsdk:"id" json:"id,computed"`
	Created           timetypes.RFC3339    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Meta              jsontypes.Normalized `tfsdk:"meta" json:"meta,computed"`
	Modified          timetypes.RFC3339    `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	ResourceAccountID types.String         `tfsdk:"resource_account_id" json:"resource_account_id,computed"`
	ResourceID        types.String         `tfsdk:"resource_id" json:"resource_id,computed"`
	ResourceType      types.String         `tfsdk:"resource_type" json:"resource_type,computed"`
	ResourceVersion   types.Int64          `tfsdk:"resource_version" json:"resource_version,computed"`
	Status            types.String         `tfsdk:"status" json:"status,computed"`
}
