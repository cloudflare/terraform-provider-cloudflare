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
	tflog.Info(ctx, "In create")
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		tflog.Info(ctx, "Has an error")
		return
	}

	tflog.Info(ctx, data.ZoneID.ValueString())
	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	tflog.Info(ctx, fmt.Sprintf("Before update config params %s", rc))

	x := data.toZarazConfigParams(ctx)
	str := spew.Sdump(x)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ %s", str))

	response, err := r.client.UpdateZarazConfig(ctx, rc, data.toZarazConfigParams(ctx))
	if err != nil {
		resp.Diagnostics.AddError("failed to create Zaraz config", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())

	response, err := r.client.GetZarazConfig(ctx, rc)
	if err != nil {
		resp.Diagnostics.AddError("failed reading D1 database", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("failed to update Zaraz Config", "Not implemented")
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

	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	response, err := r.client.UpdateZarazConfig(ctx, rc, cloudflare.UpdateZarazConfigParams{
		DebugKey: "123",
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to delete Zaraz config", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)
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
		DebugKey:     z.Config.DebugKey.String(),
		ZarazVersion: z.Config.ZarazVersion.ValueInt64(),
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

	return zz
}

func (zt *ZarazTool) toZarazToolParams(ctx context.Context) cloudflare.ZarazTool {
	tflog.Info(ctx, "Converting tool")
	zarazTool := cloudflare.ZarazTool{
		Name:             zt.Name.ValueString(),
		Type:             cloudflare.ZarazToolType(zt.Type.ValueString()),
		Enabled:          zt.Enabled.ValueBoolPointer(),
		DefaultFields:    zt.DefaultFields,
		BlockingTriggers: zt.BlockingTriggers,
		Settings:         zt.Settings,
		DefaultPurpose:   zt.DefaultPurpose.ValueString(),
	}
	tflog.Info(ctx, "after Converting tool")

	zarazTool.Component = zt.Component.ValueString()
	zarazTool.Library = zt.Library.ValueString()
	zarazTool.Permissions = zt.Permissions
	if !reflect.ValueOf(zt.Worker).IsNil() {
		zarazTool.Worker = &cloudflare.ZarazWorker{
			WorkerTag:         zt.Worker.WorkerTag,
			EscapedWorkerName: zt.Worker.EscapedWorkerName,
		}
	}

	str := spew.Sdump(zarazTool)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ TOOOOL %s", str))
	return zarazTool

}

func (ztr *ZarazTrigger) toZarazTriggerParams(ctx context.Context) cloudflare.ZarazTrigger {
	tflog.Info(ctx, "Converting trigger")
	zarazTrigger := cloudflare.ZarazTrigger{
		Name: ztr.Name.ValueString(),
	}
	tflog.Info(ctx, "after Converting trigger")

	zarazTrigger.Description = ztr.Name.ValueString()

	if !reflect.ValueOf(ztr.LoadRules).IsNil() {
		var len = len(ztr.LoadRules)
		zarazTrigger.LoadRules = make([]cloudflare.ZarazTriggerRule, len)
		for _, rule := range ztr.LoadRules {
			zarazTrigger.LoadRules = append(zarazTrigger.LoadRules, cloudflare.ZarazTriggerRule{
				Id: rule.Id.ValueString(),
			})
		}
	}

	if !reflect.ValueOf(ztr.ExcludeRules).IsNil() {
		var len = len(ztr.LoadRules)
		zarazTrigger.ExcludeRules = make([]cloudflare.ZarazTriggerRule, len)
		for _, rule := range ztr.LoadRules {
			zarazTrigger.ExcludeRules = append(zarazTrigger.ExcludeRules, cloudflare.ZarazTriggerRule{
				Id: rule.Id.ValueString(),
			})
		}
	}

	str := spew.Sdump(zarazTrigger)
	tflog.Info(ctx, fmt.Sprintf("ZARAAAAZ TRIGGGEEEER %s", str))
	return zarazTrigger

}
