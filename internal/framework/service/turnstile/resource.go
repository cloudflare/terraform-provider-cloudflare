package turnstile

import (
	"context"
	"fmt"
	"strings"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/expanders"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TurnstileWidgetResource{}
var _ resource.ResourceWithImportState = &TurnstileWidgetResource{}

func NewResource() resource.Resource {
	return &TurnstileWidgetResource{}
}

// TurnstileWidgetResource defines the resource implementation for challenge widgets.
type TurnstileWidgetResource struct {
	client *muxclient.Client
}

func (r *TurnstileWidgetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_turnstile_widget"
}

func (r *TurnstileWidgetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *TurnstileWidgetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	widget := buildChallengeWidgetFromModel(ctx, data)

	createWidget, err := r.client.V1.CreateTurnstileWidget(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()),
		cfv1.CreateTurnstileWidgetParams{
			OffLabel:     widget.OffLabel,
			Name:         widget.Name,
			Domains:      widget.Domains,
			Mode:         widget.Mode,
			BotFightMode: widget.BotFightMode,
			Region:       widget.Region,
		})
	if err != nil {
		resp.Diagnostics.AddError("Error creating challenge widget", err.Error())
		return
	}

	data = buildChallengeModelFromWidget(
		data.AccountID,
		createWidget,
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	widget, err := r.client.V1.GetTurnstileWidget(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Error reading challenge widget", err.Error())
		return
	}

	data = buildChallengeModelFromWidget(
		data.AccountID,
		widget,
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	widget := buildChallengeWidgetFromModel(ctx, data)

	updatedWidget, err := r.client.V1.UpdateTurnstileWidget(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), cfv1.UpdateTurnstileWidgetParams{
		SiteKey:      widget.SiteKey,
		OffLabel:     widget.OffLabel,
		Name:         widget.Name,
		Domains:      widget.Domains,
		Mode:         widget.Mode,
		BotFightMode: widget.BotFightMode,
		Region:       widget.Region,
	})

	if err != nil {
		resp.Diagnostics.AddError("Error reading challenge widget", err.Error())
		return
	}

	data = buildChallengeModelFromWidget(
		data.AccountID,
		updatedWidget,
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TurnstileWidgetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TurnstileWidgetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.V1.DeleteTurnstileWidget(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting challenge widget", err.Error())
		return
	}
}

func (r *TurnstileWidgetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError("Error importing challenge widget", "Invalid ID specified. Please specify the ID as \"accounts_id/sitekey\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func buildChallengeWidgetFromModel(ctx context.Context, widget *TurnstileWidgetModel) cfv1.TurnstileWidget {
	built := cfv1.TurnstileWidget{
		SiteKey:      widget.ID.ValueString(),
		Name:         widget.Name.ValueString(),
		BotFightMode: widget.BotFightMode.ValueBool(),
		Mode:         widget.Mode.ValueString(),
		Region:       widget.Region.ValueString(),
		Domains:      expanders.StringSet(ctx, widget.Domains),
		OffLabel:     widget.OffLabel.ValueBool(),
	}

	return built
}

func buildChallengeModelFromWidget(accountID types.String, widget cfv1.TurnstileWidget) *TurnstileWidgetModel {
	built := TurnstileWidgetModel{
		AccountID:    accountID,
		ID:           flatteners.String(widget.SiteKey),
		Secret:       flatteners.String(widget.Secret),
		BotFightMode: types.BoolValue(widget.BotFightMode),
		Name:         flatteners.String(widget.Name),
		Mode:         flatteners.String(widget.Mode),
		Region:       flatteners.String(widget.Region),
		OffLabel:     types.BoolValue(widget.OffLabel),
	}

	var domains []attr.Value
	for _, s := range widget.Domains {
		domains = append(domains, types.StringValue(s))
	}
	built.Domains = types.SetValueMust(types.StringType, domains)

	return &built
}
