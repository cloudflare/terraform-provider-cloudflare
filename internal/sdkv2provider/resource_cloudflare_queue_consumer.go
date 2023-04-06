package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareQueueConsumer() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareQueueConsumerSchema(),
		CreateContext: resourceCloudflareQueueConsumerCreate,
		ReadContext:   resourceCloudflareQueueConsumerRead,
		UpdateContext: resourceCloudflareQueueConsumerUpdate,
		DeleteContext: resourceCloudflareQueueConsumerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareQueueConsumerImport,
		},
		Description: "Provides the ability to manage Cloudflare Queue Consumers.",
	}
}

func resourceCloudflareQueueConsumerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	queueName := d.Get("queue_name").(string)
	scriptName := d.Get("script_name").(string)
	deadLetterQueue := d.Get("dead_letter_queue").(string)
	environment := d.Get("environment").(string)

	settings := cloudflare.QueueConsumerSettings{}
	if v, ok := d.GetOk("settings"); ok {
		settingsSet := v.(*schema.Set)
		for _, s := range settingsSet.List() {
			settingsMap := s.(map[string]interface{})
			settings.BatchSize = settingsMap["max_batch_size"].(int)
			settings.MaxRetires = settingsMap["max_retries"].(int)
			settings.MaxWaitTime = settingsMap["max_wait_time"].(int)
		}
	}

	params := cloudflare.CreateQueueConsumerParams{
		QueueName: queueName,
		Consumer: cloudflare.QueueConsumer{
			DeadLetterQueue: deadLetterQueue,
			Environment:     environment,
			ScriptName:      scriptName,
			Settings:        settings,
		},
	}

	_, err := api.CreateQueueConsumer(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare Queue Consumer: %s", err))
	}

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Queue Consumer; Queue Name: %s. Script: %s", queueName, scriptName))

	return resourceCloudflareQueueConsumerRead(ctx, d, meta)
}

func resourceCloudflareQueueConsumerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	queueName := d.Get("queue_name").(string)

	consumers, _, err := api.ListQueueConsumers(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListQueueConsumersParams{QueueName: queueName})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing Cloudflare Queue Consumers: %s", err))
	}

	for _, consumer := range consumers {
		if consumer.QueueName == queueName { // assumes Queues can only have one consumer which is true as of writing this code
			d.Set("dead_letter_queue", consumer.DeadLetterQueue)
			d.Set("environment", consumer.Environment)
			d.Set("script_name", consumer.ScriptName)
			d.Set("settings", consumer.Settings)
		}
	}

	return nil
}

func resourceCloudflareQueueConsumerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	queueName := d.Get("queue_name").(string)
	scriptName := d.Get("script_name").(string)
	deadLetterQueue := d.Get("dead_letter_queue").(string)
	environment := d.Get("environment").(string)

	settings := cloudflare.QueueConsumerSettings{}
	if v, ok := d.GetOk("settings"); ok {
		settingsSet := v.(*schema.Set)
		for _, s := range settingsSet.List() {
			settingsMap := s.(map[string]interface{})
			settings.BatchSize = settingsMap["max_batch_size"].(int)
			settings.MaxRetires = settingsMap["max_retries"].(int)
			settings.MaxWaitTime = settingsMap["max_wait_time"].(int)
		}
	}

	params := cloudflare.UpdateQueueConsumerParams{
		QueueName: queueName,
		Consumer: cloudflare.QueueConsumer{
			DeadLetterQueue: deadLetterQueue,
			Environment:     environment,
			ScriptName:      scriptName,
			Settings:        settings,
		},
	}

	_, err := api.UpdateQueueConsumer(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating Cloudflare Queue Consumer"))
	}

	return resourceCloudflareQueueConsumerRead(ctx, d, meta)
}

func resourceCloudflareQueueConsumerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	queueName := d.Get("queue_name").(string)
	scriptName := d.Get("script_name").(string)

	params := cloudflare.DeleteQueueConsumerParams{
		QueueName:    queueName,
		ConsumerName: scriptName,
	}

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Queue Consumer for queue: %+v", queueName))

	err := api.DeleteQueueConsumer(ctx, cloudflare.AccountIdentifier(accountID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting Cloudflare Queue Consumer"))
	}

	return nil
}

func resourceCloudflareQueueConsumerImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)
	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/queueID/scriptName\"", d.Id())
	}

	accountID, queueID, scriptName := attributes[0], attributes[1], attributes[2]
	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Queue Consumer id %s for account %s and script %s", queueID, accountID, scriptName))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(queueID)
	d.Set("script_name", scriptName)

	resourceCloudflareQueueConsumerRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
