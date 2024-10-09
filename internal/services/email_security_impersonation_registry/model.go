// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_impersonation_registry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityImpersonationRegistryResultEnvelope struct {
	Result EmailSecurityImpersonationRegistryModel `json:"result"`
}

type EmailSecurityImpersonationRegistryModel struct {
	ID                      types.Int64                                                               `tfsdk:"id" json:"id,computed"`
	AccountID               types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	Body                    customfield.NestedObjectList[EmailSecurityImpersonationRegistryBodyModel] `tfsdk:"body" json:"body,computed_optional"`
	Email                   types.String                                                              `tfsdk:"email" json:"email,optional"`
	IsEmailRegex            types.Bool                                                                `tfsdk:"is_email_regex" json:"is_email_regex,optional"`
	Name                    types.String                                                              `tfsdk:"name" json:"name,optional"`
	Comments                types.String                                                              `tfsdk:"comments" json:"comments,computed"`
	CreatedAt               timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DirectoryID             types.Int64                                                               `tfsdk:"directory_id" json:"directory_id,computed"`
	DirectoryNodeID         types.Int64                                                               `tfsdk:"directory_node_id" json:"directory_node_id,computed"`
	ExternalDirectoryNodeID types.String                                                              `tfsdk:"external_directory_node_id" json:"external_directory_node_id,computed"`
	LastModified            timetypes.RFC3339                                                         `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Provenance              types.String                                                              `tfsdk:"provenance" json:"provenance,computed"`
}

func (m EmailSecurityImpersonationRegistryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSecurityImpersonationRegistryModel) MarshalJSONForUpdate(state EmailSecurityImpersonationRegistryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type EmailSecurityImpersonationRegistryBodyModel struct {
	Email        types.String `tfsdk:"email" json:"email,required"`
	IsEmailRegex types.Bool   `tfsdk:"is_email_regex" json:"is_email_regex,required"`
	Name         types.String `tfsdk:"name" json:"name,required"`
}
