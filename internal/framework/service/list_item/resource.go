package list_item

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ListItemResource{}
var _ resource.ResourceWithImportState = &ListItemResource{}

func NewResource() resource.Resource {
	return &ListItemResource{}
}

// ListItemResource defines the resource implementation.
type ListItemResource struct {
	client *muxclient.Client
}

func (r *ListItemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list_item"
}

func (r *ListItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ListItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ListItemModelV1

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	item, err := createListItem(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to create list item: %s", data.ID), err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, toListItemModel(data.AccountID.ValueString(), data.ListID.ValueString(), item))...)
}

func (r *ListItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ListItemModelV1

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	listItem, err := getListItemModel(ctx, r.client, data)
	if listItem.ID.ValueString() == "" {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed reading List Item with ID: %q", data.ID), err.Error())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, listItem)...)
}

func (r *ListItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ListItemModelV1

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	itemID, err := createListItem(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to create list item: %s", data.ID), err.Error())
		return
	}
	data.ID = types.StringValue(itemID.ID)

	listItem, err := getListItemModel(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed reading List Item with ID: %q", data.ID), err.Error())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, listItem)...)
}

func (r *ListItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ListItemModelV1

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.DeleteListItems(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), cfv1.ListDeleteItemsParams{
		ID: data.ListID.ValueString(),
		Items: cfv1.ListItemDeleteRequest{
			Items: []cfv1.ListItemDeleteItemRequest{{ID: data.ID.ValueString()}},
		},
	})

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to delete list item: %s", data.ID), err.Error())
		return
	}
}

func (r *ListItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 3 {
		resp.Diagnostics.AddError("error importing list item ", "invalid ID specified. Please specify the ID as \"accountID/listID/itemID\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("list_id"), idparts[1],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[2],
	)...)
}

func toListItemModel(accountID string, listID string, item cfv1.ListItem) *ListItemModelV1 {
	model := &ListItemModelV1{
		AccountID: types.StringValue(accountID),
		ListID:    types.StringValue(listID),
		ID:        types.StringValue(item.ID),
		Comment:   flatteners.String(item.Comment),
	}
	if item.ASN != nil {
		model.ASN = types.Int64Value(int64(*item.ASN))
	}
	if item.IP != nil {
		model.IP = types.StringValue(cfv1.String(item.IP))
	}
	if item.Hostname != nil {
		model.Hostname = []*ListItemHostnameModel{
			{
				URLHostname: types.StringValue(item.Hostname.UrlHostname),
			},
		}
	}
	if item.Redirect != nil {
		model.Redirect = []*ListItemRedirectModelV1{
			{
				SourceURL:           types.StringValue(item.Redirect.SourceUrl),
				TargetURL:           types.StringValue(item.Redirect.TargetUrl),
				IncludeSubdomains:   flatteners.Bool(item.Redirect.IncludeSubdomains),
				SubpathMatching:     flatteners.Bool(item.Redirect.SubpathMatching),
				StatusCode:          types.Int64Value(int64(cfv1.Int(item.Redirect.StatusCode))),
				PreservePathSuffix:  flatteners.Bool(item.Redirect.PreservePathSuffix),
				PreserveQueryString: flatteners.Bool(item.Redirect.PreserveQueryString),
			},
		}
	}
	return model
}

func buildListItemCreateRequest(d *ListItemModelV1) cfv1.ListItemCreateRequest {
	itemType := listItemType(d)

	request := cfv1.ListItemCreateRequest{
		Comment: d.Comment.ValueString(),
	}

	switch itemType {
	case "ip":
		request.IP = cfv1.StringPtr(d.IP.ValueString())
	case "asn":
		request.ASN = cfv1.Uint32Ptr(uint32(d.ASN.ValueInt64()))
	case "hostname":
		request.Hostname = &cfv1.Hostname{
			UrlHostname: *cfv1.StringPtr(d.Hostname[0].URLHostname.ValueString()),
		}
	case "redirect":
		redirect := d.Redirect[0]
		request.Redirect = &cfv1.Redirect{
			SourceUrl:           redirect.SourceURL.ValueString(),
			TargetUrl:           redirect.TargetURL.ValueString(),
			StatusCode:          cfv1.IntPtr(int(redirect.StatusCode.ValueInt64())),
			IncludeSubdomains:   redirect.IncludeSubdomains.ValueBoolPointer(),
			SubpathMatching:     redirect.SubpathMatching.ValueBoolPointer(),
			PreservePathSuffix:  redirect.PreservePathSuffix.ValueBoolPointer(),
			PreserveQueryString: redirect.PreserveQueryString.ValueBoolPointer(),
		}
	}
	return request
}

func listItemType(d *ListItemModelV1) string {
	if d.IP.ValueString() != "" {
		return "ip"
	}
	if d.ASN.ValueInt64() > 0 {
		return "asn"
	}
	if d.Hostname != nil {
		return "hostname"
	}
	if d.Redirect != nil {
		return "redirect"
	}
	return ""
}

// getSearchTerm takes the schema and works out which "type" we are looking for
// and returns it.
func getSearchTerm(d *ListItemModelV1) string {
	if d.IP.ValueString() != "" {
		return d.IP.ValueString()
	}

	if d.ASN.ValueInt64() > 0 {
		return strconv.Itoa(int(d.ASN.ValueInt64()))
	}
	if d.Hostname != nil {
		return d.Hostname[0].URLHostname.ValueString()
	}
	if d.Redirect != nil {
		return d.Redirect[0].SourceURL.ValueString()
	}

	return ""
}

func getListItemModel(ctx context.Context, client *muxclient.Client, data *ListItemModelV1) (*ListItemModelV1, error) {
	listItem, err := client.V1.GetListItem(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), data.ListID.ValueString(), data.ID.ValueString())
	if err != nil {
		var notFoundError *cfv1.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("List item %s no longer exists", data.ID.ValueString()))
			return &ListItemModelV1{}, nil
		}
		return &ListItemModelV1{}, errors.Wrap(err, fmt.Sprintf("error reading List Item with ID %q", data.ID.ValueString()))
	}

	return toListItemModel(data.AccountID.ValueString(), data.ListID.ValueString(), listItem), nil
}

func createListItem(ctx context.Context, client *muxclient.Client, data *ListItemModelV1) (cfv1.ListItem, error) {
	listItemType := listItemType(data)
	listID := data.ListID.ValueString()
	accountIdentifier := cfv1.AccountIdentifier(data.AccountID.ValueString())
	list, err := client.V1.GetList(ctx, accountIdentifier, data.ListID.ValueString())
	if err != nil {
		return cfv1.ListItem{}, fmt.Errorf("unable to find list with id %s", listID)
	}

	if list.Kind != listItemType {
		return cfv1.ListItem{}, fmt.Errorf("items of type %s can not be added to lists of type %s", listItemType, list.Kind)
	}

	_, err = client.V1.CreateListItem(ctx, accountIdentifier, cfv1.ListCreateItemParams{
		ID:   data.ListID.ValueString(),
		Item: buildListItemCreateRequest(data),
	},
	)
	if err != nil {
		return cfv1.ListItem{}, fmt.Errorf("failed to create list item: %w", err)
	}

	// terraform-plugin-framework doesn't have a built in retryable HTTP client (yet)
	// so we use a simple loop with a break when we get the data we expect here.
	var items []cfv1.ListItem
	for attempts := 1; attempts < 5 && len(items) != 1; attempts++ {
		// this is extremely inefficient however, it's the only option as the list
		// service uses a polling model and does not expose the ID.
		searchTerm := getSearchTerm(data)
		items, err = client.V1.ListListItems(ctx, accountIdentifier, cfv1.ListListItemsParams{
			ID:     listID,
			Search: searchTerm,
		})
		if len(items) == 1 {
			break
		}

		for _, item := range items {
			if item.Redirect != nil && item.Redirect.SourceUrl == searchTerm {
				items = []cfv1.ListItem{item}
				break
			}

			if item.IP != nil && cfv1.String(item.IP) == searchTerm {
				items = []cfv1.ListItem{item}
				break
			}
		}
		if len(items) == 1 {
			break
		}

		//lintignore:R018
		time.Sleep(time.Duration(attempts) * time.Second)
	}

	if len(items) != 1 {
		return cfv1.ListItem{}, errors.New("failed to match exactly one list item")
	}

	return items[0], nil
}
