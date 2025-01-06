package list

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ListResource{}
var _ resource.ResourceWithImportState = &ListResource{}

func NewResource() resource.Resource {
	return &ListResource{}
}

type ListResource struct {
	client *muxclient.Client
}

func (r *ListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list"
}

func (r *ListResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*muxclient.Client)
}

func (r *ListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ListModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := plan.AccountID.ValueString()
	createParams := cloudflare.ListCreateParams{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		Kind:        plan.Kind.ValueString(),
	}

	list, err := r.client.V1.CreateList(ctx, cloudflare.AccountIdentifier(accountID), createParams)
	if err != nil {
		resp.Diagnostics.AddError("Error creating List", err.Error())
		return
	}

	plan.ID = types.StringValue(list.ID)

	if len(plan.Items) > 0 {
		items := buildListItemsCreateRequest(plan.Items)
		_, err = r.client.V1.CreateListItems(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListCreateItemsParams{
			ID:    list.ID,
			Items: items,
		})
		if err != nil {
			resp.Diagnostics.AddError("Error creating List Items", err.Error())
			return
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ListModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	list, err := r.client.V1.GetList(ctx, cloudflare.AccountIdentifier(state.AccountID.ValueString()), state.ID.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "could not find list") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading List", err.Error())
		return
	}

	state.Name = types.StringValue(list.Name)
	state.Description = types.StringValue(list.Description)
	state.Kind = types.StringValue(list.Kind)

	if len(state.Items) > 0 {
		items, err := r.client.V1.ListListItems(ctx, cloudflare.AccountIdentifier(state.AccountID.ValueString()), cloudflare.ListListItemsParams{
			ID: state.ID.ValueString(),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error reading List Items", err.Error())
			return
		}

		state.Items = buildListItemModels(items)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ListModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.UpdateList(ctx, cloudflare.AccountIdentifier(plan.AccountID.ValueString()), cloudflare.ListUpdateParams{
		ID:          plan.ID.ValueString(),
		Description: plan.Description.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating List", err.Error())
		return
	}

	items := buildListItemsCreateRequest(plan.Items)
	_, err = r.client.V1.ReplaceListItems(ctx, cloudflare.AccountIdentifier(plan.AccountID.ValueString()), cloudflare.ListReplaceItemsParams{
		ID:    plan.ID.ValueString(),
		Items: items,
	})

	if err != nil {
		resp.Diagnostics.AddError("Error updating List Items", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ListModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.DeleteList(ctx, cloudflare.AccountIdentifier(state.AccountID.ValueString()), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting List", err.Error())
		return
	}
}

func (r *ListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Invalid ID specified. Please provide ID in the format 'accountID/listID'. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func buildListItemModels(items []cloudflare.ListItem) []ListItemModel {
	var models []ListItemModel

	for _, item := range items {
		model := ListItemModel{
			Value:   make([]ItemValueModel, 1),
			Comment: flatteners.String(item.Comment),
		}

		value := &model.Value[0]
		if item.ASN != nil {
			value.ASN = types.Int64Value(int64(*item.ASN))
		}
		if item.IP != nil {
			value.IP = types.StringValue(cloudflare.String(item.IP))
		}
		if item.Hostname != nil {
			value.Hostname = []ListItemHostnameModel{
				{
					URLHostname: types.StringValue(item.Hostname.UrlHostname),
				},
			}
		}
		if item.Redirect != nil {
			value.Redirect = []ListItemRedirectModel{
				{
					SourceURL:           types.StringValue(item.Redirect.SourceUrl),
					TargetURL:           types.StringValue(item.Redirect.TargetUrl),
					IncludeSubdomains:   optBoolToString(item.Redirect.IncludeSubdomains),
					SubpathMatching:     optBoolToString(item.Redirect.SubpathMatching),
					StatusCode:          flatteners.Int64(int64(cloudflare.Int(item.Redirect.StatusCode))),
					PreservePathSuffix:  optBoolToString(item.Redirect.PreservePathSuffix),
					PreserveQueryString: optBoolToString(item.Redirect.PreserveQueryString),
				},
			}
		}

		models = append(models, model)
	}

	return models
}

func buildListItemsCreateRequest(items []ListItemModel) []cloudflare.ListItemCreateRequest {
	var listItems []cloudflare.ListItemCreateRequest

	for _, item := range items {
		payload := cloudflare.ListItemCreateRequest{
			Comment: item.Comment.ValueString(),
		}
		value := item.Value[0]

		if !value.IP.IsNull() {
			payload.IP = cloudflare.StringPtr(value.IP.ValueString())
		}

		if !value.ASN.IsNull() {
			asn := uint32(value.ASN.ValueInt64())
			payload.ASN = &asn
		}

		if len(value.Hostname) > 0 && !value.Hostname[0].URLHostname.IsNull() {
			payload.Hostname = &cloudflare.Hostname{
				UrlHostname: value.Hostname[0].URLHostname.ValueString(),
			}
		}

		if len(value.Redirect) > 0 {
			redirect := value.Redirect[0]
			var statusCode *int = nil
			if !redirect.StatusCode.IsNull() {
				statusCode = cloudflare.IntPtr(int(redirect.StatusCode.ValueInt64()))
			}
			payload.Redirect = &cloudflare.Redirect{
				SourceUrl:           redirect.SourceURL.ValueString(),
				TargetUrl:           redirect.TargetURL.ValueString(),
				StatusCode:          statusCode,
				IncludeSubdomains:   stringToOptBool(redirect.IncludeSubdomains.ValueString()),
				SubpathMatching:     stringToOptBool(redirect.SubpathMatching.ValueString()),
				PreserveQueryString: stringToOptBool(redirect.PreserveQueryString.ValueString()),
				PreservePathSuffix:  stringToOptBool(redirect.PreservePathSuffix.ValueString()),
			}
		}

		listItems = append(listItems, payload)
	}

	return listItems
}

func stringToOptBool(s string) *bool {
	switch s {
	case "enabled":
		return cloudflare.BoolPtr(true)
	case "disabled":
		return cloudflare.BoolPtr(false)
	default:
		return nil
	}
}

func optBoolToString(b *bool) types.String {
	if b != nil {
		if *b {
			return types.StringValue("enabled")
		}
		return types.StringValue("disabled")
	}
	return types.StringNull()
}
