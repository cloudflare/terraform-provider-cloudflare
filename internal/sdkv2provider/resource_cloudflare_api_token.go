package sdkv2provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareApiToken() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareApiTokenSchema(),
		CreateContext: resourceCloudflareApiTokenCreate,
		ReadContext:   resourceCloudflareApiTokenRead,
		UpdateContext: resourceCloudflareApiTokenUpdate,
		DeleteContext: resourceCloudflareApiTokenDelete,
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare API tokens.

			Read more about permission groups and their applicable scopes in the
			[developer documentation](https://developers.cloudflare.com/api/tokens/create/permissions).
		`),
	}
}

func buildAPIToken(d *schema.ResourceData) cloudflare.APIToken {
	token := cloudflare.APIToken{}

	token.Name = d.Get("name").(string)
	token.Policies = resourceDataToApiTokenPolices(d)

	ipsIn := []string{}
	ipsNotIn := []string{}
	if ips, ok := d.GetOk("condition.0.request_ip.0.in"); ok {
		ipsIn = expandInterfaceToStringList(ips.(*schema.Set).List())
	}

	if ips, ok := d.GetOk("condition.0.request_ip.0.not_in"); ok {
		ipsNotIn = expandInterfaceToStringList(ips.(*schema.Set).List())
	}

	if len(ipsIn) > 0 || len(ipsNotIn) > 0 {
		token.Condition = &cloudflare.APITokenCondition{
			RequestIP: &cloudflare.APITokenRequestIPCondition{},
		}

		if len(ipsIn) > 0 {
			token.Condition.RequestIP.In = ipsIn
		}

		if len(ipsNotIn) > 0 {
			token.Condition.RequestIP.NotIn = ipsNotIn
		}
	}
	if before, ok := d.GetOk("not_before"); ok {
		notBefore, _ := time.Parse(time.RFC3339Nano, before.(string))
		token.NotBefore = &notBefore
	}
	if expires, ok := d.GetOk("expires_on"); ok {
		expiresOn, _ := time.Parse(time.RFC3339Nano, expires.(string))
		token.ExpiresOn = &expiresOn
	}

	return token
}

func resourceCloudflareApiTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	name := d.Get("name").(string)

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare API Token: name %s", name))

	t := buildAPIToken(d)
	t, err := client.CreateAPIToken(ctx, t)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare API Token %q: %w", name, err))
	}

	d.SetId(t.ID)
	d.Set("status", t.Status)
	d.Set("value", t.Value)

	return resourceCloudflareApiTokenRead(ctx, d, meta)
}

func resourceDataToApiTokenPolices(d *schema.ResourceData) []cloudflare.APITokenPolicies {
	policies := d.Get("policy").(*schema.Set).List()
	var cfPolicies []cloudflare.APITokenPolicies

	for _, p := range policies {
		policy := p.(map[string]interface{})

		permissionGroups := expandInterfaceToStringList(policy["permission_groups"].(*schema.Set).List())
		if len(permissionGroups) == 0 {
			continue
		}
		var cfPermissionGroups []cloudflare.APITokenPermissionGroups
		for _, pg := range permissionGroups {
			cfPermissionGroups = append(cfPermissionGroups, cloudflare.APITokenPermissionGroups{
				ID: pg,
			})
		}

		cfResources := map[string]interface{}{}
		for k, v := range policy["resources"].(map[string]interface{}) {
			// value can be object or just a string ("*"), try to convert it to map
			obj := map[string]string{}
			if err := json.Unmarshal([]byte(v.(string)), &obj); err == nil {
				cfResources[k] = obj
			} else {
				cfResources[k] = v
			}
		}

		cfPolicies = append(cfPolicies, cloudflare.APITokenPolicies{
			Effect:           policy["effect"].(string),
			Resources:        cfResources,
			PermissionGroups: cfPermissionGroups,
		})
	}

	return cfPolicies
}

func resourceCloudflareApiTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	tokenID := d.Id()

	t, err := client.GetAPIToken(ctx, tokenID)

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare API Token: %+v", t))

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Cloudflare API Token %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Cloudflare API Token %q: %w", d.Id(), err))
	}

	policies := []map[string]interface{}{}

	for _, p := range t.Policies {
		permissionGroups := []string{}
		for _, v := range p.PermissionGroups {
			permissionGroups = append(permissionGroups, v.ID)
		}

		policies = append(policies, map[string]interface{}{
			"resources":         p.Resources,
			"permission_groups": permissionGroups,
			"effect":            p.Effect,
		})
	}

	d.Set("name", t.Name)
	d.Set("policy", policies)
	d.Set("status", t.Status)
	d.Set("issued_on", t.IssuedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", t.ModifiedOn.Format(time.RFC3339Nano))

	if t.ExpiresOn != nil {
		d.Set("expires_on", t.ExpiresOn.Format(time.RFC3339Nano))
	}

	if t.NotBefore != nil {
		d.Set("not_before", t.NotBefore.Format(time.RFC3339Nano))
	}

	var ipIn []string
	var ipNotIn []string
	if t.Condition != nil && t.Condition.RequestIP != nil && t.Condition.RequestIP.In != nil {
		ipIn = t.Condition.RequestIP.In
	}

	if t.Condition != nil && t.Condition.RequestIP != nil && t.Condition.RequestIP.NotIn != nil {
		ipNotIn = t.Condition.RequestIP.NotIn
	}

	if len(ipIn) > 0 || len(ipNotIn) > 0 {
		d.Set("condition", []map[string]interface{}{{
			"request_ip": []map[string]interface{}{{
				"not_in": ipNotIn,
				"in":     ipIn,
			}},
		}})
	}

	return nil
}

func resourceCloudflareApiTokenUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	name := d.Get("name").(string)
	tokenID := d.Id()

	t := buildAPIToken(d)

	tflog.Info(ctx, fmt.Sprintf("Updating Cloudflare API Token: name %s", name))

	_, err := client.UpdateAPIToken(ctx, tokenID, t)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Cloudflare API Token %q: %w", name, err))
	}

	return resourceCloudflareApiTokenRead(ctx, d, meta)
}

func resourceCloudflareApiTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	tokenID := d.Id()

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare API Token: id %s", tokenID))

	err := client.DeleteAPIToken(ctx, tokenID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare API Token: %w", err))
	}

	return nil
}
