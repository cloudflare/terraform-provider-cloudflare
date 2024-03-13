package access_mutual_tls_hostname_settings

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AccessMutualTLSHostnameSettingsResource{}
var _ resource.ResourceWithImportState = &AccessMutualTLSHostnameSettingsResource{}

func NewResource() resource.Resource {
	return &AccessMutualTLSHostnameSettingsResource{}
}

type AccessMutualTLSHostnameSettingsResource struct {
	client *cloudflare.API
}

func (r *AccessMutualTLSHostnameSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_mutual_tls_hostname_settings"
}

func (r *AccessMutualTLSHostnameSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AccessMutualTLSHostnameSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AccessMutualTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	identifier := getIdentifier(data)

	updatedSettings, err := r.update(ctx, data, identifier)
	if err != nil {
		resp.Diagnostics.AddError("error updating Access Mutual TLS Hostname Settings", err.Error())
		return
	}

	data = buildAccessMutualTLSHostnameSettingsModel(updatedSettings, data.ZoneID, data.AccountID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessMutualTLSHostnameSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessMutualTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	identifier := getIdentifier(data)

	accessMutualTLSHostnameSettings, err := r.client.GetAccessMutualTLSHostnameSettings(ctx, identifier)
	if err != nil {
		resp.Diagnostics.AddError("error reading Access Mutual TLS Hostname Settings", err.Error())
		return
	}
	data = buildAccessMutualTLSHostnameSettingsModel(accessMutualTLSHostnameSettings, data.ZoneID, data.AccountID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Helper function used by both Update and Create
// We take the difference in hostnames between state and plan to determine which hostnames to delete.
// Those hostnames are updated with empty settings to delete them.
func (r *AccessMutualTLSHostnameSettingsResource) update(ctx context.Context, data *AccessMutualTLSHostnameSettingsModel, identifier *cloudflare.ResourceContainer) ([]cloudflare.AccessMutualTLSHostnameSettings, error) {
	currentSettings, err := r.client.GetAccessMutualTLSHostnameSettings(ctx, identifier)
	if err != nil {
		return nil, fmt.Errorf("error finding Access Mutual TLS Hostname Settings: %w", err)
	}
	currentHostnames := make(map[string]struct{})
	for _, setting := range currentSettings {
		currentHostnames[setting.Hostname] = struct{}{}
	}

	updatedSettings := make([]cloudflare.AccessMutualTLSHostnameSettings, 0)
	updatedHostnames := make(map[string]struct{})

	settings := data.Settings
	for _, setting := range settings {
		updatedSetting := cloudflare.AccessMutualTLSHostnameSettings{
			Hostname:                    setting.Hostname.ValueString(),
			ChinaNetwork:                setting.ChinaNetwork.ValueBoolPointer(),
			ClientCertificateForwarding: setting.ClientCertificateForwarding.ValueBoolPointer(),
		}
		updatedSettings = append(updatedSettings, updatedSetting)
		updatedHostnames[updatedSetting.Hostname] = struct{}{}
	}

	for hostname := range currentHostnames {
		if _, ok := updatedHostnames[hostname]; !ok {
			// Hostname has been removed
			updatedSettings = append(updatedSettings, cloudflare.AccessMutualTLSHostnameSettings{
				Hostname: hostname,
			})
		}
	}

	updatedAccessMutualTLSHostnameSettings := cloudflare.UpdateAccessMutualTLSHostnameSettingsParams{
		Settings: updatedSettings,
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Mutal TLS Hostname Settings from struct: %+v", updatedAccessMutualTLSHostnameSettings))

	resultUpdatedSettings, err := r.client.UpdateAccessMutualTLSHostnameSettings(ctx, identifier, updatedAccessMutualTLSHostnameSettings)
	if err != nil {
		return nil, fmt.Errorf("error updating Access Mutual TLS Hostname Settings for %s %q: %w", identifier.Level, identifier.Identifier, err)
	}
	return resultUpdatedSettings, nil
}

func (r *AccessMutualTLSHostnameSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AccessMutualTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	identifier := getIdentifier(data)

	updatedSettings, err := r.update(ctx, data, identifier)
	if err != nil {
		resp.Diagnostics.AddError("error updating Access Mutual TLS Hostname Settings", err.Error())
		return
	}

	data = buildAccessMutualTLSHostnameSettingsModel(updatedSettings, data.ZoneID, data.AccountID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessMutualTLSHostnameSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AccessMutualTLSHostnameSettingsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	identifier := getIdentifier(data)

	currentSettings, err := r.client.GetAccessMutualTLSHostnameSettings(ctx, identifier)
	if err != nil {
		resp.Diagnostics.AddError("error finding Access Mutual TLS Hostname Settings", err.Error())
		return
	}
	updatedSettings := make([]cloudflare.AccessMutualTLSHostnameSettings, 0)
	for _, setting := range currentSettings {
		updatedSetting := cloudflare.AccessMutualTLSHostnameSettings{
			Hostname: setting.Hostname,
		}
		updatedSettings = append(updatedSettings, updatedSetting)
	}

	// To actually delete the settings we issue an update for the changed hostnames with all fields set to false
	deletedSettings := cloudflare.UpdateAccessMutualTLSHostnameSettingsParams{
		Settings: updatedSettings,
	}

	_, err = r.client.UpdateAccessMutualTLSHostnameSettings(ctx, identifier, deletedSettings)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error removing Access Mutual TLS Hostname Settings for %s %q", identifier.Level, identifier.Identifier), err.Error())
		return
	}
}

func (r *AccessMutualTLSHostnameSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	attributes := strings.Split(req.ID, "/")

	invalidIDMessage := "invalid ID (\"%s\") specified, should be in format \"account/<account_id>\" or \"zone/<zone_id>\""
	if len(attributes) != 2 {
		resp.Diagnostics.AddError("error importing Access Mutual TLS Hostname Settings", fmt.Sprintf(invalidIDMessage, req.ID))
		return
	}

	identifierType, identifierID := attributes[0], attributes[1]

	if !(identifierType == "zone" || identifierType == "account") {
		resp.Diagnostics.AddError("invalid id specified", fmt.Sprintf(invalidIDMessage, req.ID))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Mutual TLS Hostname Settings: for %s %s", identifierType, identifierID))

	schemaIdentifierName := "account_id"
	if identifierType == "zone" {
		schemaIdentifierName = "zone_id"
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(schemaIdentifierName), identifierID)...)
}

func buildAccessMutualTLSHostnameSettingsModel(settings []cloudflare.AccessMutualTLSHostnameSettings, zoneID, accountID basetypes.StringValue) *AccessMutualTLSHostnameSettingsModel {
	model := &AccessMutualTLSHostnameSettingsModel{
		ZoneID:    zoneID,
		AccountID: accountID,
	}
	for _, setting := range settings {
		model.Settings = append(model.Settings, Settings{
			Hostname:                    types.StringValue(setting.Hostname),
			ChinaNetwork:                types.BoolValue(cloudflare.Bool(setting.ChinaNetwork)),
			ClientCertificateForwarding: types.BoolValue(cloudflare.Bool(setting.ClientCertificateForwarding)),
		})
	}
	return model
}

func getIdentifier(data *AccessMutualTLSHostnameSettingsModel) *cloudflare.ResourceContainer {
	accountID := data.AccountID
	zoneID := data.ZoneID

	var identifier *cloudflare.ResourceContainer
	if accountID.ValueString() != "" {
		identifier = cloudflare.AccountIdentifier(accountID.ValueString())
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID.ValueString())
	}

	return identifier
}
