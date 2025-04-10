// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsScriptSecretResultDataSourceEnvelope struct {
	Result WorkersForPlatformsScriptSecretDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsScriptSecretDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" path:"secret_name,computed"`
	SecretName        types.String                   `tfsdk:"secret_name" path:"secret_name,optional"`
	AccountID         types.String                   `tfsdk:"account_id" path:"account_id,required"`
	DispatchNamespace types.String                   `tfsdk:"dispatch_namespace" path:"dispatch_namespace,required"`
	ScriptName        types.String                   `tfsdk:"script_name" path:"script_name,required"`
	Format            types.String                   `tfsdk:"format" json:"format,computed"`
	KeyBase64         types.String                   `tfsdk:"key_base64" json:"key_base64,computed"`
	Name              types.String                   `tfsdk:"name" json:"name,computed"`
	Text              types.String                   `tfsdk:"text" json:"text,computed"`
	Type              types.String                   `tfsdk:"type" json:"type,computed"`
	Usages            customfield.List[types.String] `tfsdk:"usages" json:"usages,computed"`
	Algorithm         jsontypes.Normalized           `tfsdk:"algorithm" json:"algorithm,computed"`
	KeyJwk            jsontypes.Normalized           `tfsdk:"key_jwk" json:"key_jwk,computed"`
}

func (m *WorkersForPlatformsScriptSecretDataSourceModel) toReadParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceScriptSecretGetParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceScriptSecretGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
