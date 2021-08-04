package cloudflare

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareNotificationPolicyWebhooks() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareNotificationPolicyWebhooksCreate,
		Read:   resourceCloudflareNotificationPolicyWebhooksRead,
		Update: resourceCloudflareNotificationPolicyWebhooksUpdate,
		Delete: resourceCloudflareNotificationPolicyWebhooksDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareNotificationPolicyWebhooksImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_success": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_failure": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudflareNotificationPolicyWebhooksCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	notificationWebhooks := buildNotificationPolicyWebhooks(d)

	webhooksDestination, err := client.CreateNotificationWebhooks(context.Background(), accountID, &notificationWebhooks)

	if err != nil {
		return fmt.Errorf("error connecting webhooks destination %s: %s", notificationWebhooks.Name, err)
	}
	formattedWebhookID, err := uuid.Parse(webhooksDestination.Result.ID)
	if err != nil {
		return fmt.Errorf("error setting notification webhooks: %s", err)
	}

	d.SetId(formattedWebhookID.String())

	return resourceCloudflareNotificationPolicyWebhooksRead(d, meta)
}

func resourceCloudflareNotificationPolicyWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	webhooksDestinationID := d.Id()
	accountID := d.Get("account_id").(string)

	notificationWebhooks, err := client.GetNotificationWebhooks(context.Background(), accountID, webhooksDestinationID)

	name := d.Get("name").(string)
	if err != nil {
		return fmt.Errorf("error retrieving notification webhooks %s: %s", name, err)
	}
	id, err := uuid.Parse(notificationWebhooks.Result.ID)
	if err != nil {
		return fmt.Errorf("error setting notification webhooks %s: %s", name, err)
	}
	d.Set("id", id.String())
	d.Set("name", notificationWebhooks.Result.Name)
	d.Set("url", notificationWebhooks.Result.URL)
	d.Set("created", notificationWebhooks.Result.CreatedAt.Format(time.RFC3339))
	d.Set("type", notificationWebhooks.Result.Type)

	if notificationWebhooks.Result.LastSuccess != nil {
		d.Set("last_success", notificationWebhooks.Result.LastSuccess.Format(time.RFC3339))
	}
	if notificationWebhooks.Result.LastFailure != nil {
		d.Set("last_failure", notificationWebhooks.Result.LastFailure.Format(time.RFC3339))
	}

	return nil
}

func resourceCloudflareNotificationPolicyWebhooksUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	webhooksID := d.Id()
	accountID := d.Get("account_id").(string)

	notificationWebhooks := buildNotificationPolicyWebhooks(d)

	_, err := client.UpdateNotificationWebhooks(context.Background(), accountID, webhooksID, &notificationWebhooks)

	if err != nil {
		return fmt.Errorf("error updating notification webhooks destination %s: %s", webhooksID, err)
	}

	return resourceCloudflareNotificationPolicyWebhooksRead(d, meta)
}

func resourceCloudflareNotificationPolicyWebhooksDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	webhooksID := d.Id()
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteNotificationWebhooks(context.Background(), accountID, webhooksID)

	if err != nil {
		return fmt.Errorf("error deleting notification webhooks destination %s: %s", webhooksID, err)
	}
	return nil
}

func resourceCloudflareNotificationPolicyWebhooksImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/webhooksID\"", d.Id())
	}

	accountID, webhooksID := attributes[0], attributes[1]
	d.SetId(webhooksID)
	d.Set("account_id", accountID)

	resourceCloudflareNotificationPolicyWebhooksRead(d, meta)

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
