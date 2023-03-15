package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWorkerCronTrigger() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerCronTriggerSchema(),
		CreateContext: resourceCloudflareWorkerCronTriggerUpdate,
		ReadContext:   resourceCloudflareWorkerCronTriggerRead,
		UpdateContext: resourceCloudflareWorkerCronTriggerUpdate,
		DeleteContext: resourceCloudflareWorkerCronTriggerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkerCronTriggerImport,
		},
		Description: heredoc.Doc(fmt.Sprintf(`
			Worker Cron Triggers allow users to map a cron expression to a Worker script
			using a %s listener that enables Workers to be executed on a
			schedule. Worker Cron Triggers are ideal for running periodic jobs for
			maintenance or calling third-party APIs to collect up-to-date data.
		`, "`ScheduledEvent`")),
	}
}

// resourceCloudflareWorkerCronTriggerUpdate is used for creation and updates of
// Worker Cron Triggers as the remote API endpoint is shared uses HTTP PUT.
func resourceCloudflareWorkerCronTriggerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	scriptName := d.Get("script_name").(string)

	crons := transformSchemaToWorkerCronTriggerStruct(d)
	_, err := client.UpdateWorkerCronTriggers(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateWorkerCronTriggersParams{
		ScriptName: scriptName,
		Crons:      crons,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update Worker Cron Trigger: %w", err))
	}

	d.SetId(stringChecksum(scriptName))

	return nil
}

func resourceCloudflareWorkerCronTriggerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	scriptName := d.Get("script_name").(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := cloudflare.ListWorkerCronTriggersParams{
		ScriptName: scriptName,
	}

	s, err := client.ListWorkerCronTriggers(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			d.SetId("")
			return nil
		}

		return diag.FromErr(fmt.Errorf("failed to read Worker Cron Trigger: %w", err))
	}

	if err := d.Set("schedules", transformWorkerCronTriggerStructToSet(s)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set schedules attribute: %w", err))
	}

	return nil
}

func resourceCloudflareWorkerCronTriggerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	scriptName := d.Get("script_name").(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.UpdateWorkerCronTriggers(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateWorkerCronTriggersParams{
		ScriptName: scriptName,
		Crons:      []cloudflare.WorkerCronTrigger{},
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func resourceCloudflareWorkerCronTriggerImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/scriptName"`, d.Id())
	}

	accountID, scriptName := attributes[0], attributes[1]

	d.Set("script_name", scriptName)
	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(stringChecksum(scriptName))

	resourceCloudflareWorkerCronTriggerRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func transformWorkerCronTriggerStructToSet(triggers []cloudflare.WorkerCronTrigger) *schema.Set {
	returnSet := schema.NewSet(schema.HashString, []interface{}{})

	for _, trigger := range triggers {
		returnSet.Add(trigger.Cron)
	}

	return returnSet
}

func transformSchemaToWorkerCronTriggerStruct(d *schema.ResourceData) []cloudflare.WorkerCronTrigger {
	triggers := []cloudflare.WorkerCronTrigger{}
	schedules := d.Get("schedules").(*schema.Set).List()

	for _, schedule := range schedules {
		triggers = append(triggers, cloudflare.WorkerCronTrigger{Cron: schedule.(string)})
	}

	return triggers
}
