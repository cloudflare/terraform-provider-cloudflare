// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/registrar"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegistrarDomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RegistrarDomainsResultDataSourceModel] `json:"result,computed"`
}

type RegistrarDomainsDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[RegistrarDomainsResultDataSourceModel] `tfsdk:"result"`
}

func (m *RegistrarDomainsDataSourceModel) toListParams(_ context.Context) (params registrar.DomainListParams, diags diag.Diagnostics) {
	params = registrar.DomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type RegistrarDomainsResultDataSourceModel struct {
	ID                types.String                                                               `tfsdk:"id" json:"id,computed"`
	Available         types.Bool                                                                 `tfsdk:"available" json:"available,computed"`
	CanRegister       types.Bool                                                                 `tfsdk:"can_register" json:"can_register,computed"`
	CreatedAt         timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CurrentRegistrar  types.String                                                               `tfsdk:"current_registrar" json:"current_registrar,computed"`
	ExpiresAt         timetypes.RFC3339                                                          `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	Locked            types.Bool                                                                 `tfsdk:"locked" json:"locked,computed"`
	RegistrantContact customfield.NestedObject[RegistrarDomainsRegistrantContactDataSourceModel] `tfsdk:"registrant_contact" json:"registrant_contact,computed"`
	RegistryStatuses  types.String                                                               `tfsdk:"registry_statuses" json:"registry_statuses,computed"`
	SupportedTld      types.Bool                                                                 `tfsdk:"supported_tld" json:"supported_tld,computed"`
	TransferIn        customfield.NestedObject[RegistrarDomainsTransferInDataSourceModel]        `tfsdk:"transfer_in" json:"transfer_in,computed"`
	UpdatedAt         timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type RegistrarDomainsRegistrantContactDataSourceModel struct {
	Address      types.String `tfsdk:"address" json:"address,computed"`
	City         types.String `tfsdk:"city" json:"city,computed"`
	Country      types.String `tfsdk:"country" json:"country,computed"`
	FirstName    types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName     types.String `tfsdk:"last_name" json:"last_name,computed"`
	Organization types.String `tfsdk:"organization" json:"organization,computed"`
	Phone        types.String `tfsdk:"phone" json:"phone,computed"`
	State        types.String `tfsdk:"state" json:"state,computed"`
	Zip          types.String `tfsdk:"zip" json:"zip,computed"`
	ID           types.String `tfsdk:"id" json:"id,computed"`
	Address2     types.String `tfsdk:"address2" json:"address2,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
	Fax          types.String `tfsdk:"fax" json:"fax,computed"`
}

type RegistrarDomainsTransferInDataSourceModel struct {
	AcceptFoa         types.String `tfsdk:"accept_foa" json:"accept_foa,computed"`
	ApproveTransfer   types.String `tfsdk:"approve_transfer" json:"approve_transfer,computed"`
	CanCancelTransfer types.Bool   `tfsdk:"can_cancel_transfer" json:"can_cancel_transfer,computed"`
	DisablePrivacy    types.String `tfsdk:"disable_privacy" json:"disable_privacy,computed"`
	EnterAuthCode     types.String `tfsdk:"enter_auth_code" json:"enter_auth_code,computed"`
	UnlockDomain      types.String `tfsdk:"unlock_domain" json:"unlock_domain,computed"`
}
