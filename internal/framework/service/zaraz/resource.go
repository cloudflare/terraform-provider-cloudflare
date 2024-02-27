package zaraz

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZarazConfigResource{}
var _ resource.ResourceWithImportState = &ZarazConfigResource{}

func NewResource() resource.Resource {
	return &ZarazConfigResource{}
}

// ZarazConfigResource defines the resource implementation.
type ZarazConfigResource struct {
	client *cloudflare.API
}

func (r *ZarazConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zaraz_config"
}

func (r *ZarazConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZarazConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, data.ZoneID.ValueString())
	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())

	x := data.toZarazConfigParams(ctx)
	str := spew.Sdump(x)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ %s", str))

	_, err := r.client.UpdateZarazConfig(ctx, rc, data.toZarazConfigParams(ctx))
	if err != nil {
		resp.Diagnostics.AddError("failed to create Zaraz config", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	zoneId := data.ZoneID.ValueString()

	rc := cloudflare.ZoneIdentifier(zoneId)

	_, err := r.client.GetZarazConfig(ctx, rc)
	if err != nil {
		resp.Diagnostics.AddError("failed reading Zaraz config", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, data.ZoneID.ValueString())
	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())

	x := data.toZarazConfigParams(ctx)
	str := spew.Sdump(x)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ %s", str))

	_, err := r.client.UpdateZarazConfig(ctx, rc, data.toZarazConfigParams(ctx))
	if err != nil {
		resp.Diagnostics.AddError("failed to update Zaraz config", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Todo: Can we implment Zaraz TF without delete?
// If no, then what does delete for zaraz config mean in the terraform context?
// Does deleting a config mean, resetting the config?
func (r *ZarazConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	zoneId := data.ZoneID.ValueString()
	zoneId = strings.ReplaceAll(zoneId, "\\\"", "")
	tflog.Info(ctx, fmt.Sprintf("Zone ID %s", zoneId))
	zoneId = strings.ReplaceAll(zoneId, "\"", "")
	tflog.Info(ctx, fmt.Sprintf("Zone ID %s", zoneId))

	// TODO call the reset ZarazConfig API here to reset the config
	// rc := cloudflare.ZoneIdentifier(zoneId)
	// response, err := r.client.UpdateZarazConfig(ctx, rc, cloudflare.UpdateZarazConfigParams{
	// 	DebugKey: "123",
	// })
	// if err != nil {
	// 	resp.Diagnostics.AddError("failed to delete Zaraz config", err.Error())
	// 	return
	// }

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"invalid import identifier",
			fmt.Sprintf("expected import identifier to be resourceLevel/resourceIdentifier/rulesetID. got: %q", req.ID),
		)
		return
	}
	resourceLevel, resourceIdentifier, rulesetID := idParts[0], idParts[1], idParts[2]

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), rulesetID)...)
	if resourceLevel == "zone" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), resourceIdentifier)...)
	}
}

func toZarazResourceModel(ctx context.Context, zoneID, accountID basetypes.StringValue, in cloudflare.UpdateZarazConfigParams) *ZarazConfigModel {
	return nil
}

func (z *ZarazConfigModel) toZarazConfigParams(ctx context.Context) cloudflare.UpdateZarazConfigParams {
	zz := cloudflare.UpdateZarazConfigParams{
		DebugKey:      z.Config.DebugKey.ValueString(),
		ZarazVersion:  z.Config.ZarazVersion.ValueInt64(),
		HistoryChange: z.Config.HistoryChange.ValueBoolPointer(),
	}

	if z.Config.Tools != nil {
		zz.Tools = make(map[string]cloudflare.ZarazTool)
		for toolId, tool := range z.Config.Tools {
			zz.Tools[toolId] = tool.toZarazToolParams(ctx)

		}
	}

	if z.Config.Triggers != nil {
		zz.Triggers = make(map[string]cloudflare.ZarazTrigger)
		for triggerId, trigger := range z.Config.Triggers {
			zz.Triggers[triggerId] = trigger.toZarazTriggerParams(ctx)
		}
	}

	if z.Config.Settings != nil {
		zz.Settings = cloudflare.ZarazConfigSettings{
			AutoInjectScript:    z.Config.Settings.AutoInjectScript.ValueBoolPointer(),
			InjectIframes:       z.Config.Settings.InjectIframes.ValueBoolPointer(),
			Ecommerce:           z.Config.Settings.Ecommerce.ValueBoolPointer(),
			HideQueryParams:     z.Config.Settings.HideQueryParams.ValueBoolPointer(),
			HideIpAddress:       z.Config.Settings.HideIpAddress.ValueBoolPointer(),
			HideUserAgent:       z.Config.Settings.HideUserAgent.ValueBoolPointer(),
			HideExternalReferer: z.Config.Settings.HideExternalReferer.ValueBoolPointer(),
			CookieDomain:        z.Config.Settings.CookieDomain.ValueString(),
			InitPath:            z.Config.Settings.InitPath.ValueString(),
			ScriptPath:          z.Config.Settings.ScriptPath.ValueString(),
			TrackPath:           z.Config.Settings.TrackPath.ValueString(),
			EventsApiPath:       z.Config.Settings.EventsApiPath.ValueString(),
			McRootPath:          z.Config.Settings.McRootPath.ValueString(),
		}

		if !reflect.ValueOf(z.Config.Settings.ContextEnricher).IsNil() {
			zz.Settings.ContextEnricher = &cloudflare.ZarazWorker{
				WorkerTag:         z.Config.Settings.ContextEnricher.WorkerTag.ValueString(),
				EscapedWorkerName: z.Config.Settings.ContextEnricher.EscapedWorkerName.ValueString(),
				MutableId:         z.Config.Settings.ContextEnricher.MutableId.ValueString(),
			}
		}
	}

	return zz
}

func (zt *ZarazTool) toZarazToolParams(ctx context.Context) cloudflare.ZarazTool {
	zarazTool := cloudflare.ZarazTool{
		Name:             zt.Name.ValueString(),
		Type:             cloudflare.ZarazToolType(zt.Type.ValueString()),
		Enabled:          zt.Enabled.ValueBoolPointer(),
		DefaultFields:    zt.DefaultFields,
		BlockingTriggers: zt.BlockingTriggers,
		Settings:         zt.Settings,
		DefaultPurpose:   zt.DefaultPurpose.ValueString(),
	}

	zarazTool.Component = zt.Component.ValueString()
	zarazTool.Library = zt.Library.ValueString()
	zarazTool.Permissions = zt.Permissions
	if !reflect.ValueOf(zt.Worker).IsNil() {
		zarazTool.Worker = &cloudflare.ZarazWorker{
			WorkerTag:         zt.Worker.WorkerTag.ValueString(),
			EscapedWorkerName: zt.Worker.EscapedWorkerName.ValueString(),
			MutableId:         zt.Worker.MutableId.ValueString(),
		}
	}

	if !reflect.ValueOf(zt.Actions).IsNil() {
		zarazTool.Actions = make(map[string]cloudflare.ZarazAction)
		for actionID, action := range zt.Actions {
			zarazTool.Actions[actionID] = cloudflare.ZarazAction{
				BlockingTriggers: toStringArray(ctx, action.BlockingTriggers),
				FiringTriggers:   toStringArray(ctx, action.FiringTriggers),
				Data:             action.Data,
				ActionType:       action.ActionType.ValueString(),
			}
		}
	}

	str := spew.Sdump(zarazTool)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ TOOOOL %s", str))
	return zarazTool

}

func toStringArray(ctx context.Context, arrayValues []types.String) []string {
	result := make([]string, len(arrayValues))
	for index, value := range arrayValues {
		result[index] = value.ValueString()
	}

	return result
}

func (ztr *ZarazTrigger) toZarazTriggerParams(ctx context.Context) cloudflare.ZarazTrigger {
	zarazTrigger := cloudflare.ZarazTrigger{
		Name: ztr.Name.ValueString(),
	}
	zarazTrigger.Description = ztr.Name.ValueString()
	if !reflect.ValueOf(ztr.LoadRules).IsNil() {
		zarazTrigger.LoadRules = make([]cloudflare.ZarazTriggerRule, 0)
		for _, rule := range ztr.LoadRules {
			zarazTrigger.LoadRules = append(zarazTrigger.LoadRules, cloudflare.ZarazTriggerRule{
				Match: rule.Match.ValueString(),
				Op:    rule.Op.ValueString(),
				Value: rule.Value.ValueString(),
			})
		}
	}

	if !reflect.ValueOf(ztr.ExcludeRules).IsNil() {
		zarazTrigger.ExcludeRules = make([]cloudflare.ZarazTriggerRule, 0)
		for _, rule := range ztr.ExcludeRules {
			zarazTrigger.ExcludeRules = append(zarazTrigger.ExcludeRules, cloudflare.ZarazTriggerRule{
				Id: rule.Id.ValueString(),
			})
		}
	}

	str := spew.Sdump(zarazTrigger)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ TRIGGGEEEER %s", str))
	return zarazTrigger

}
