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

func resourceCloudflareQueue() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareQueueSchema(),
		CreateContext: resourceCloudflareQueueCreate,
		ReadContext:   resourceCloudflareQueueRead,
		UpdateContext: resourceCloudflareQueueUpdate,
		DeleteContext: resourceCloudflareQueueDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareQueueImport,
		},
		Description: "Provides the ability to manage Cloudflare Workers Queue features.",
	}
}

func resourceCloudflareQueueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	queueName := d.Get("name").(string)

	req := cloudflare.CreateQueueParams{
		Name: queueName,
	}
	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Workers Queue from struct: %+v", req))

	r, err := client.CreateQueue(ctx, cloudflare.AccountIdentifier(accountID), req)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating workers queue"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	if d.Get("consumers") != nil {
		consumers := d.Get("consumers").([]interface{})
		for _, consumer := range consumers {
			consumerMap := consumer.(map[string]interface{})
			settings := consumerMap["settings"].([]interface{})[0].(map[string]interface{})
			req := cloudflare.CreateQueueConsumerParams{
				QueueName: queueName,
				Consumer: cloudflare.QueueConsumer{
					Service: "workers",
					Name: consumerMap["script_name"].(string),
					QueueName: queueName,
					Environment: consumerMap["environment"].(string),
					ScriptName: consumerMap["script_name"].(string),
					Settings: cloudflare.QueueConsumerSettings{
						BatchSize: settings["batch_size"].(int),
						MaxRetires: settings["max_retries"].(int),
						MaxWaitTime: settings["max_wait_time_ms"].(int),
					},
				},
			}
			
			_, err := client.CreateQueueConsumer(ctx, cloudflare.AccountIdentifier(accountID), req)
			// println(fmt.Sprintf("result: %+v", result))
			if err != nil {
				return diag.FromErr(errors.Wrap(err, "error creating workers queue consumer"))
			}
		}
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Workers Queue ID: %s. Name: %s", d.Id(), queueName))

	return resourceCloudflareQueueRead(ctx, d, meta)
}

func resourceCloudflareQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	queueID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	resp, _, err := client.ListQueues(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.ListQueuesParams{})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error reading queues"))
	}

	var queue cloudflare.Queue
	for _, r := range resp {
		if r.ID == queueID {
			queue = r
			d.Set("name", r.Name)
			consumers := make([]map[string]interface{}, 0)
			if len(r.Consumers) > 0 {
				for _, consumer := range r.Consumers {
					consumerMap := map[string]interface{}{
						"environment": consumer.Environment,
						"queue_name": consumer.QueueName,
						"script_name": consumer.ScriptName,
						"settings": []map[string]interface{}{
							{
								"batch_size": consumer.Settings.BatchSize,
								"max_retries": consumer.Settings.MaxRetires,
								"max_wait_time_ms": consumer.Settings.MaxWaitTime,
							},
						},
					}
					consumers = append(consumers, consumerMap)
				}
				d.Set("consumers", consumers)
			}
			break
		}
	}

	if queue.ID == "" {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceCloudflareQueueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	existingName, updatedName := d.GetChange("name")

	_, err := client.UpdateQueue(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateQueueParams{
		Name:        existingName.(string),
		UpdatedName: updatedName.(string),
	})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating workers queue"))
	}

	return resourceCloudflareQueueRead(ctx, d, meta)
}

func resourceCloudflareQueueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Workers Queue with id: %+v", d.Id()))

	err := client.DeleteQueue(ctx, cloudflare.AccountIdentifier(accountID), d.Get("name").(string))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting workers queue"))
	}

	d.SetId("")
	return nil
}

func resourceCloudflareQueueImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)
	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/queueID\"", d.Id())
	}

	accountID, queueID := attributes[0], attributes[1]
	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Queue id %s for account %s", queueID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(queueID)

	resourceCloudflareQueueRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
