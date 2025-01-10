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

type RegistrarDomainResultDataSourceEnvelope struct {
	Result RegistrarDomainDataSourceModel `json:"result,computed"`
}

type RegistrarDomainResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RegistrarDomainDataSourceModel] `json:"result,computed"`
}

type RegistrarDomainDataSourceModel struct {
	AccountID         types.String                                                              `tfsdk:"account_id" path:"account_id,optional"`
	DomainName        types.String                                                              `tfsdk:"domain_name" path:"domain_name,optional"`
	Available         types.Bool                                                                `tfsdk:"available" json:"available,computed"`
	CanRegister       types.Bool                                                                `tfsdk:"can_register" json:"can_register,computed"`
	CreatedAt         timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CurrentRegistrar  types.String                                                              `tfsdk:"current_registrar" json:"current_registrar,computed"`
	ExpiresAt         timetypes.RFC3339                                                         `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	ID                types.String                                                              `tfsdk:"id" json:"id,computed"`
	Locked            types.Bool                                                                `tfsdk:"locked" json:"locked,computed"`
	RegistryStatuses  types.String                                                              `tfsdk:"registry_statuses" json:"registry_statuses,computed"`
	SupportedTld      types.Bool                                                                `tfsdk:"supported_tld" json:"supported_tld,computed"`
	UpdatedAt         timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	RegistrantContact customfield.NestedObject[RegistrarDomainRegistrantContactDataSourceModel] `tfsdk:"registrant_contact" json:"registrant_contact,computed"`
	TransferIn        customfield.NestedObject[RegistrarDomainTransferInDataSourceModel]        `tfsdk:"transfer_in" json:"transfer_in,computed"`
	Filter            *RegistrarDomainFindOneByDataSourceModel                                  `tfsdk:"filter"`
}

func (m *RegistrarDomainDataSourceModel) toReadParams(_ context.Context) (params registrar.DomainGetParams, diags diag.Diagnostics) {
	params = registrar.DomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *RegistrarDomainDataSourceModel) toListParams(_ context.Context) (params registrar.DomainListParams, diags diag.Diagnostics) {
	params = registrar.DomainListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type RegistrarDomainRegistrantContactDataSourceModel struct {
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

type RegistrarDomainTransferInDataSourceModel struct {
	AcceptFoa         types.String `tfsdk:"accept_foa" json:"accept_foa,computed"`
	ApproveTransfer   types.String `tfsdk:"approve_transfer" json:"approve_transfer,computed"`
	CanCancelTransfer types.Bool   `tfsdk:"can_cancel_transfer" json:"can_cancel_transfer,computed"`
	DisablePrivacy    types.String `tfsdk:"disable_privacy" json:"disable_privacy,computed"`
	EnterAuthCode     types.String `tfsdk:"enter_auth_code" json:"enter_auth_code,computed"`
	UnlockDomain      types.String `tfsdk:"unlock_domain" json:"unlock_domain,computed"`
}

type RegistrarDomainFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
