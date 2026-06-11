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

type ShareResultDataSourceEnvelope struct {
	Result ShareDataSourceModel `json:"result,computed"`
}

type ShareDataSourceModel struct {
	ID                           types.String                                                `tfsdk:"id" path:"share_id,computed"`
	ShareID                      types.String                                                `tfsdk:"share_id" path:"share_id,optional"`
	AccountID                    types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	IncludeRecipientCounts       types.Bool                                                  `tfsdk:"include_recipient_counts" query:"include_recipient_counts,optional"`
	IncludeResources             types.Bool                                                  `tfsdk:"include_resources" query:"include_resources,optional"`
	AccountName                  types.String                                                `tfsdk:"account_name" json:"account_name,computed"`
	AssociatedRecipientCount     types.Int64                                                 `tfsdk:"associated_recipient_count" json:"associated_recipient_count,computed"`
	AssociatingRecipientCount    types.Int64                                                 `tfsdk:"associating_recipient_count" json:"associating_recipient_count,computed"`
	Created                      timetypes.RFC3339                                           `tfsdk:"created" json:"created,computed" format:"date-time"`
	DisassociatedRecipientCount  types.Int64                                                 `tfsdk:"disassociated_recipient_count" json:"disassociated_recipient_count,computed"`
	DisassociatingRecipientCount types.Int64                                                 `tfsdk:"disassociating_recipient_count" json:"disassociating_recipient_count,computed"`
	Kind                         types.String                                                `tfsdk:"kind" json:"kind,computed"`
	Modified                     timetypes.RFC3339                                           `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name                         types.String                                                `tfsdk:"name" json:"name,computed"`
	OrganizationID               types.String                                                `tfsdk:"organization_id" json:"organization_id,computed"`
	Status                       types.String                                                `tfsdk:"status" json:"status,computed"`
	TargetType                   types.String                                                `tfsdk:"target_type" json:"target_type,computed"`
	Resources                    customfield.NestedObjectList[ShareResourcesDataSourceModel] `tfsdk:"resources" json:"resources,computed"`
	Filter                       *ShareFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *ShareDataSourceModel) toReadParams(_ context.Context) (params resource_sharing.ResourceSharingGetParams, diags diag.Diagnostics) {
	params = resource_sharing.ResourceSharingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.IncludeRecipientCounts.IsNull() {
		params.IncludeRecipientCounts = cloudflare.F(m.IncludeRecipientCounts.ValueBool())
	}
	if !m.IncludeResources.IsNull() {
		params.IncludeResources = cloudflare.F(m.IncludeResources.ValueBool())
	}

	return
}

func (m *ShareDataSourceModel) toListParams(_ context.Context) (params resource_sharing.ResourceSharingListParams, diags diag.Diagnostics) {
	mFilterResourceTypes := []resource_sharing.ResourceSharingListParamsResourceType{}
	if m.Filter.ResourceTypes != nil {
		for _, item := range *m.Filter.ResourceTypes {
			mFilterResourceTypes = append(mFilterResourceTypes, resource_sharing.ResourceSharingListParamsResourceType(item.ValueString()))
		}
	}
	mFilterTag := []string{}
	if m.Filter.Tag != nil {
		for _, item := range *m.Filter.Tag {
			mFilterTag = append(mFilterTag, item.ValueString())
		}
	}

	params = resource_sharing.ResourceSharingListParams{
		AccountID:     cloudflare.F(m.AccountID.ValueString()),
		ResourceTypes: cloudflare.F(mFilterResourceTypes),
		Tag:           cloudflare.F(mFilterTag),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(resource_sharing.ResourceSharingListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.IncludeRecipientCounts.IsNull() {
		params.IncludeRecipientCounts = cloudflare.F(m.IncludeRecipientCounts.ValueBool())
	}
	if !m.IncludeResources.IsNull() {
		params.IncludeResources = cloudflare.F(m.IncludeResources.ValueBool())
	}
	if !m.Filter.Kind.IsNull() {
		params.Kind = cloudflare.F(resource_sharing.ResourceSharingListParamsKind(m.Filter.Kind.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(resource_sharing.ResourceSharingListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(resource_sharing.ResourceSharingListParamsStatus(m.Filter.Status.ValueString()))
	}
	if !m.Filter.TargetType.IsNull() {
		params.TargetType = cloudflare.F(resource_sharing.ResourceSharingListParamsTargetType(m.Filter.TargetType.ValueString()))
	}

	return
}

type ShareResourcesDataSourceModel struct {
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

type ShareFindOneByDataSourceModel struct {
	Direction     types.String    `tfsdk:"direction" query:"direction,computed_optional"`
	Kind          types.String    `tfsdk:"kind" query:"kind,optional"`
	Order         types.String    `tfsdk:"order" query:"order,computed_optional"`
	ResourceTypes *[]types.String `tfsdk:"resource_types" query:"resource_types,optional"`
	Status        types.String    `tfsdk:"status" query:"status,optional"`
	Tag           *[]types.String `tfsdk:"tag" query:"tag,optional"`
	TargetType    types.String    `tfsdk:"target_type" query:"target_type,optional"`
}
