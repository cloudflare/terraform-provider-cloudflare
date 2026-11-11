// This file is intentionally NOT code-generated.

package zero_trust_access_service_token

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ action.ActionWithConfigure = (*RefreshAction)(nil)

// RefreshAction extends a service token's expiry by its configured duration,
// using POST /accounts/{account_id}/access/service_tokens/{id}/refresh. The
// client ID and secret don't change, so the resource's state stays valid.
//
// There is no rotate action on purpose: an action can't write to resource
// state, so the new client_secret would be lost. Rotation is done on the
// resource by incrementing client_secret_version.
type RefreshAction struct {
	client *cloudflare.Client
}

type RefreshActionModel struct {
	AccountID      types.String `tfsdk:"account_id"`
	ServiceTokenID types.String `tfsdk:"service_token_id"`
}

func NewRefreshAction() action.Action {
	return &RefreshAction{}
}

func (a *RefreshAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_access_service_token_refresh"
}

func (a *RefreshAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Refreshes an Access service token: its expiry moves forward by the token's duration. The client ID and secret stay the same. To rotate the secret, increment client_secret_version on the cloudflare_zero_trust_access_service_token resource instead.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID that owns the service token.",
				Required:    true,
			},
			"service_token_id": schema.StringAttribute{
				Description: "The ID of the service token to refresh, such as cloudflare_zero_trust_access_service_token.example.id.",
				Required:    true,
			},
		},
	}
}

func (a *RefreshAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unexpected action configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	a.client = client
}

func (a *RefreshAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	if a.client == nil {
		resp.Diagnostics.AddError(
			"provider not configured",
			"The Cloudflare client is missing, so the service token can't be refreshed. Please report this issue to the provider developers.",
		)
		return
	}

	var data RefreshActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	token, err := a.client.ZeroTrust.Access.ServiceTokens.Refresh(
		ctx,
		data.ServiceTokenID.ValueString(),
		zero_trust.AccessServiceTokenRefreshParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to refresh service token", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Refreshed service token %s; it now expires at %s.", token.ID, token.ExpiresAt.Format(time.RFC3339)),
		})
	}
}
