package cloudflare

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareNotificationPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareNotificationPolicySchema(),
		CreateContext: resourceCloudflareNotificationPolicyCreate,
		ReadContext: resourceCloudflareNotificationPolicyRead,
		UpdateContext: resourceCloudflareNotificationPolicyUpdate,
		DeleteContext: resourceCloudflareNotificationPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceNotificationPolicyImport,
		},
	}
}

func resourceCloudflareNotificationPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	notificationPolicy := buildNotificationPolicy(d)

	policy, err := client.CreateNotificationPolicy(context.Background(), accountID, notificationPolicy)

	if err != nil {
		return fmt.Errorf("error creating policy %s: %s", notificationPolicy.Name, err)
	}
	d.SetId(policy.Result.ID)

	return resourceCloudflareNotificationPolicyRead(d, meta)
}

func resourceCloudflareNotificationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	policyID := d.Id()
	accountID := d.Get("account_id").(string)

	policy, err := client.GetNotificationPolicy(context.Background(), accountID, policyID)

	name := d.Get("name").(string)
	if err != nil {
		return fmt.Errorf("error retrieving notification policy %s: %s", name, err)
	}

	d.Set("name", policy.Result.Name)
	d.Set("enabled", policy.Result.Enabled)
	d.Set("alert_type", policy.Result.AlertType)
	d.Set("description", policy.Result.Description)
	d.Set("created", policy.Result.Created.Format(time.RFC3339))
	d.Set("modified", policy.Result.Modified.Format(time.RFC3339))

	if policy.Result.Filters != nil && len(policy.Result.Filters) > 0 {
		if err := d.Set("filters", flattenNotificationPolicyFilter(policy.Result.Filters)); err != nil {
			return fmt.Errorf("failed to set filters: %s", err)
		}
	}

	if err := d.Set("email_integration", setNotificationMechanisms(policy.Result.Mechanisms["email"])); err != nil {
		return fmt.Errorf("failed to set email integration: %s", err)
	}

	if err := d.Set("pagerduty_integration", setNotificationMechanisms(policy.Result.Mechanisms["pagerduty"])); err != nil {
		return fmt.Errorf("failed to set pagerduty integration: %s", err)
	}

	if err := d.Set("webhooks_integration", setNotificationMechanisms(policy.Result.Mechanisms["webhooks"])); err != nil {
		return fmt.Errorf("failed to set webhooks integration: %s", err)
	}

	return nil
}

func resourceCloudflareNotificationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	policyID := d.Id()
	accountID := d.Get("account_id").(string)

	notificationPolicy := buildNotificationPolicy(d)
	notificationPolicy.ID = policyID

	_, err := client.UpdateNotificationPolicy(context.Background(), accountID, &notificationPolicy)

	if err != nil {
		return fmt.Errorf("error updating notification policy %s: %s", policyID, err)
	}

	return resourceCloudflareNotificationPolicyRead(d, meta)
}

func resourceCloudflareNotificationPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	policyID := d.Id()
	accountID := d.Get("account_id").(string)

	_, err := client.DeleteNotificationPolicy(context.Background(), accountID, policyID)

	if err != nil {
		return fmt.Errorf("error deleting notification policy %s: %s", policyID, err)
	}
	return nil
}

func resourceNotificationPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/policyID\"", d.Id())
	}

	accountID, policyID := attributes[0], attributes[1]
	d.SetId(policyID)
	d.Set("account_id", accountID)

	resourceCloudflareNotificationPolicyRead(d, meta)

	return []*schema.ResourceData{d}, nil

}

func buildNotificationPolicy(d *schema.ResourceData) cloudflare.NotificationPolicy {
	notificationPolicy := cloudflare.NotificationPolicy{}
	notificationPolicy.Mechanisms = make(map[string]cloudflare.NotificationMechanismIntegrations)
	notificationPolicy.Conditions = make(map[string]interface{})
	notificationPolicy.Filters = make(map[string][]string)

	if name, ok := d.GetOk("name"); ok {
		notificationPolicy.Name = name.(string)
	}

	if description, ok := d.GetOk("description"); ok {
		notificationPolicy.Description = description.(string)
	}

	if enabled, ok := d.GetOk("enabled"); ok {
		notificationPolicy.Enabled = enabled.(bool)
	}

	if alertType, ok := d.GetOk("alert_type"); ok {
		notificationPolicy.AlertType = alertType.(string)
	}

	if emails, ok := d.GetOk("email_integration"); ok {
		notificationPolicy.Mechanisms["email"] = getNotificationMechanisms(emails.(*schema.Set))
	}

	if webhooks, ok := d.GetOk("webhooks_integration"); ok {
		notificationPolicy.Mechanisms["webhooks"] = getNotificationMechanisms(webhooks.(*schema.Set))
	}

	if pagerduty, ok := d.GetOk("pagerduty_integration"); ok {
		notificationPolicy.Mechanisms["pagerduty"] = getNotificationMechanisms(pagerduty.(*schema.Set))
	}

	if filters, ok := d.GetOk("filters"); ok {
		notificationPolicy.Filters = expandNotificationPolicyFilter(filters.([]interface{}))
	}

	return notificationPolicy
}

func expandNotificationPolicyFilter(list []interface{}) map[string][]string {
	filters := make(map[string][]string)
	for _, listItem := range list {
		for k, mapItem := range listItem.(map[string]interface{}) {
			for _, v := range mapItem.(*schema.Set).List() {
				filters[k] = append(filters[k], v.(string))
			}
		}
	}
	return filters
}

func flattenNotificationPolicyFilter(filters map[string][]string) []interface{} {
	filtersMap := make(map[string]interface{})
	for k, v := range filters {
		set := schema.NewSet(schema.HashString, []interface{}{})
		for _, value := range v {
			set.Add(value)
		}
		filtersMap[k] = set
	}
	return []interface{}{filtersMap}
}

func getNotificationMechanisms(s *schema.Set) []cloudflare.NotificationMechanismData {
	var notificationMechanisms []cloudflare.NotificationMechanismData

	for _, m := range s.List() {
		mechanism := m.(map[string]interface{})
		data := cloudflare.NotificationMechanismData{
			ID:   mechanism["id"].(string),
			Name: mechanism["name"].(string),
		}
		notificationMechanisms = append(notificationMechanisms, data)
	}

	return notificationMechanisms
}

func setNotificationMechanisms(md []cloudflare.NotificationMechanismData) *schema.Set {
	mechanisms := make([]interface{}, 0)
	data := make(map[string]interface{})

	for _, m := range md {
		data["name"] = m.Name
		data["id"] = m.ID
		mechanisms = append(mechanisms, data)
	}

	return schema.NewSet(schema.HashResource(mechanismData), mechanisms)
}
