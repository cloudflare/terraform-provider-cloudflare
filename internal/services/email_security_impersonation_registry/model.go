// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityImpersonationRegistryResultEnvelope struct {
	Result EmailSecurityImpersonationRegistryModel `json:"result"`
}

type EmailSecurityImpersonationRegistryModel struct {
	ID                      types.String      `tfsdk:"id" json:"id,computed"`
	AccountID               types.String      `tfsdk:"account_id" path:"account_id,required"`
	Email                   types.String      `tfsdk:"email" json:"email,required"`
	IsEmailRegex            types.Bool        `tfsdk:"is_email_regex" json:"is_email_regex,required"`
	Name                    types.String      `tfsdk:"name" json:"name,required"`
	Comments                types.String      `tfsdk:"comments" json:"comments,optional"`
	DirectoryID             types.Int64       `tfsdk:"directory_id" json:"directory_id,optional"`
	DirectoryNodeID         types.Int64       `tfsdk:"directory_node_id" json:"directory_node_id,optional"`
	ExternalDirectoryNodeID types.String      `tfsdk:"external_directory_node_id" json:"external_directory_node_id,optional"`
	Provenance              types.String      `tfsdk:"provenance" json:"provenance,optional"`
	CreatedAt               timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastModified            timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	ModifiedAt              timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m EmailSecurityImpersonationRegistryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSecurityImpersonationRegistryModel) MarshalJSONForUpdate(state EmailSecurityImpersonationRegistryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
