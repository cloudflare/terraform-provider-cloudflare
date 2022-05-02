package cloudflare

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareNotificationPolicyWebhooks() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareNotificationPolicyWebhooksSchema(),
		CreateContext: resourceCloudflareNotificationPolicyWebhooksCreate,
		ReadContext:   resourceCloudflareNotificationPolicyWebhooksRead,
		UpdateContext: resourceCloudflareNotificationPolicyWebhooksUpdate,
		DeleteContext: resourceCloudflareNotificationPolicyWebhooksDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareNotificationPolicyWebhooksImport,
		},
	}
}

func resourceCloudflareNotificationPolicyWebhooksCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	notificationWebhooks := buildNotificationPolicyWebhooks(d)

	webhooksDestination, err := client.CreateNotificationWebhooks(ctx, accountID, &notificationWebhooks)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error connecting webhooks destination %s: %s", notificationWebhooks.Name, err))
	}
	formattedWebhookID, err := uuid.Parse(webhooksDestination.Result.ID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting notification webhooks: %s", err))
	}

	d.SetId(formattedWebhookID.String())

	return resourceCloudflareNotificationPolicyWebhooksRead(ctx, d, meta)
}

func resourceCloudflareNotificationPolicyWebhooksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	webhooksDestinationID := d.Id()
	accountID := d.Get("account_id").(string)

	notificationWebhooks, err := client.GetNotificationWebhooks(ctx, accountID, webhooksDestinationID)

	name := d.Get("name").(string)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error retrieving notification webhooks %s: %s", name, err))
	}

	d.Set("name", notificationWebhooks.Result.Name)
	d.Set("url", notificationWebhooks.Result.URL)
	d.Set("created_at", notificationWebhooks.Result.CreatedAt.Format(time.RFC3339))
	d.Set("type", notificationWebhooks.Result.Type)

	if notificationWebhooks.Result.LastSuccess != nil {
		d.Set("last_success", notificationWebhooks.Result.LastSuccess.Format(time.RFC3339))
	}
	if notificationWebhooks.Result.LastFailure != nil {
		d.Set("last_failure", notificationWebhooks.Result.LastFailure.Format(time.RFC3339))
	}

	return nil
}

func resourceCloudflareNotificationPolicyWebhooksUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	webhooksID := d.Id()
	accountID := d.Get("account_id").(string)

	notificationWebhooks := buildNotificationPolicyWebhooks(d)

	_, err := client.UpdateNotificationWebhooks(ctx, accountID, webhooksID, &notificationWebhooks)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating notification webhooks destination %s: %s", webhooksID, err))
	}

	return resourceCloudflareNotificationPolicyWebhooksRead(ctx, d, meta)
}

func resourceCloudflareNotificationPolicyWebhooksDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	webhooksID := d.Id()
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteNotificationWebhooks(ctx, accountID, webhooksID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting notification webhooks destination %s: %s", webhooksID, err))
	}
	return nil
}

func resourceCloudflareNotificationPolicyWebhooksImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/webhooksID\"", d.Id())
	}

	accountID, webhooksID := attributes[0], attributes[1]
	d.SetId(webhooksID)
	d.Set("account_id", accountID)

	resourceCloudflareNotificationPolicyWebhooksRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil

}

func buildNotificationPolicyWebhooks(d *schema.ResourceData) cloudflare.NotificationUpsertWebhooks {
	webhooks := cloudflare.NotificationUpsertWebhooks{}

	if name, ok := d.GetOk("name"); ok {
		webhooks.Name = name.(string)
	}

	if url, ok := d.GetOk("url"); ok {
		webhooks.URL = url.(string)
	}

	if secret, ok := d.GetOk("secret"); ok {
		webhooks.Secret = secret.(string)
	}

	return webhooks
}
