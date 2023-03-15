package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAddressMap() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAddressMapSchema(),
		CreateContext: resourceCloudflareAddressMapCreate,
		ReadContext:   resourceCloudflareAddressMapRead,
		UpdateContext: resourceCloudflareAddressMapUpdate,
		DeleteContext: resourceCloudflareAddressMapDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAddressMapImport,
		},
		Description: heredoc.Doc(`
			Provides the ability to manage IP addresses that can be used by DNS records when
			they are proxied through Cloudflare.
		`),
	}
}

func resourceCloudflareAddressMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	payload := cloudflare.CreateAddressMapParams{}

	if v, ok := d.GetOk("description"); ok {
		desc := v.(string)
		payload.Description = &desc
	}

	if v, ok := d.GetOk("enabled"); ok {
		enabled := v.(bool)
		payload.Enabled = &enabled
	}

	if v, ok := d.GetOk("ips"); ok {
		payload.IPs = getIPsFromSchema(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("memberships"); ok {
		payload.Memberships = getMembershipContainersFromSchema(v.(*schema.Set).List())
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating AddressMap from struct: %+v", payload))

	addressMap, err := client.CreateAddressMap(ctx, cloudflare.AccountIdentifier(accountID), payload)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating AddressMap: %w", err))
	}

	d.SetId(addressMap.ID)
	return resourceCloudflareAddressMapRead(ctx, d, meta)
}

func resourceCloudflareAddressMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	addressMap, err := client.GetAddressMap(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading address map %q: %w", d.Id(), err))
	}

	d.Set("description", cloudflare.String(addressMap.Description))
	d.Set("default_sni", cloudflare.String(addressMap.DefaultSNI))
	d.Set("enabled", cloudflare.Bool(addressMap.Enabled))
	d.Set("can_delete", cloudflare.Bool(addressMap.Deletable))
	d.Set("can_modify_ips", cloudflare.Bool(addressMap.CanModifyIPs))
	d.Set("ips", convertIPsToSchema(addressMap.IPs))
	d.Set("memberships", convertMembershipsToSchema(addressMap.Memberships))

	return nil
}

func resourceCloudflareAddressMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	hasChanges := false
	payload := cloudflare.UpdateAddressMapParams{ID: d.Id()}

	if d.HasChange("enabled") {
		hasChanges = true
		payload.Enabled = cloudflare.BoolPtr(d.Get("enabled").(bool))
	}

	if d.HasChange("description") {
		hasChanges = true
		payload.Description = cloudflare.StringPtr(d.Get("description").(string))
	}

	if d.HasChange("default_sni") {
		hasChanges = true
		payload.DefaultSNI = cloudflare.StringPtr(d.Get("default_sni").(string))
	}

	if hasChanges {
		tflog.Debug(ctx, fmt.Sprintf("Updating AddressMap from struct: %+v", payload))

		if _, err := client.UpdateAddressMap(ctx, cloudflare.AccountIdentifier(accountID), payload); err != nil {
			return diag.FromErr(fmt.Errorf("error updating address map %q: %w", d.Id(), err))
		}
	}

	membershipDiff := make(map[cloudflare.AddressMapMembershipContainer]int)
	if d.HasChange("memberships") {
		old, new := d.GetChange("memberships")
		oldMembers, newMembers := getMembershipsFromSchema(old.(*schema.Set).List()), getMembershipsFromSchema(new.(*schema.Set).List())

		for _, member := range newMembers {
			membershipDiff[cloudflare.AddressMapMembershipContainer{Identifier: member.Identifier, Kind: member.Kind}] += 1
		}
		for _, member := range oldMembers {
			membershipDiff[cloudflare.AddressMapMembershipContainer{Identifier: member.Identifier, Kind: member.Kind}] -= 1
		}
	}

	ipsDiff := make(map[string]int)
	if d.HasChange("ips") {
		old, new := d.GetChange("ips")
		oldIPs, newIPs := getIPsFromSchema(old.(*schema.Set).List()), getIPsFromSchema(new.(*schema.Set).List())
		for _, ip := range newIPs {
			ipsDiff[ip] += 1
		}
		for _, ip := range oldIPs {
			ipsDiff[ip] -= 1
		}
	}

	// Add ip addresses before adding any memberships
	for ip, flag := range ipsDiff {
		if flag > 0 {
			tflog.Debug(ctx, fmt.Sprintf("Adding ip %q to AddressMap %q", ip, d.Id()))

			if err := client.CreateIPAddressToAddressMap(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.CreateIPAddressToAddressMapParams{ID: d.Id(), IP: ip}); err != nil {
				return diag.FromErr(fmt.Errorf("error adding ip %q from address map %q: %w", ip, d.Id(), err))
			}
		}
	}

	// Add memberships
	for member, flag := range membershipDiff {
		if flag > 0 {
			tflog.Debug(ctx, fmt.Sprintf("Adding membership %v to AddressMap %q", member, d.Id()))

			params := cloudflare.CreateMembershipToAddressMapParams{ID: d.Id(), Membership: member}
			if err := client.CreateMembershipToAddressMap(ctx, cloudflare.AccountIdentifier(accountID), params); err != nil {
				return diag.FromErr(fmt.Errorf("error adding %v from address map %q: %w", member, d.Id(), err))
			}
		}
	}

	// Remove memberships before removing any ip address
	for member, flag := range membershipDiff {
		if flag < 0 {
			tflog.Debug(ctx, fmt.Sprintf("Removing membership %v from AddressMap %q", member, d.Id()))

			params := cloudflare.DeleteMembershipFromAddressMapParams{ID: d.Id(), Membership: member}
			if err := client.DeleteMembershipFromAddressMap(ctx, cloudflare.AccountIdentifier(accountID), params); err != nil {
				return diag.FromErr(fmt.Errorf("error removing %v from address map %q: %w", member, d.Id(), err))
			}
		}
	}

	// Remove ip addresses
	for ip, flag := range ipsDiff {
		if flag < 0 {
			tflog.Debug(ctx, fmt.Sprintf("Removing ip %q from AddressMap %q", ip, d.Id()))

			if err := client.DeleteIPAddressFromAddressMap(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.DeleteIPAddressFromAddressMapParams{ID: d.Id(), IP: ip}); err != nil {
				return diag.FromErr(fmt.Errorf("error removing ip %q from address map %q: %w", ip, d.Id(), err))
			}
		}
	}

	return resourceCloudflareAddressMapRead(ctx, d, meta)
}

func resourceCloudflareAddressMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	if err := client.DeleteAddressMap(ctx, cloudflare.AccountIdentifier(accountID), d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting address map %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareAddressMapImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)
	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "<account_id>/<address_map_id>"`, d.Id())
	}

	accountID, addressMapID := attributes[0], attributes[1]
	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Address Map: id %s for account %s", addressMapID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(addressMapID)

	if readErr := resourceCloudflareAddressMapRead(ctx, d, meta); readErr != nil {
		return nil, fmt.Errorf("failed to read Address Map state")
	}

	return []*schema.ResourceData{d}, nil
}

func convertIPsToSchema(ips []cloudflare.AddressMapIP) []interface{} {
	data := []interface{}{}
	for _, ip := range ips {
		data = append(data, map[string]string{
			"ip": ip.IP,
		})
	}
	return data
}

func getIPsFromSchema(values []interface{}) []string {
	ips := []string{}

	for _, value := range values {
		m := value.(map[string]interface{})
		ips = append(ips, m["ip"].(string))
	}

	return ips
}

func convertMembershipsToSchema(members []cloudflare.AddressMapMembership) []interface{} {
	data := []interface{}{}
	for _, member := range members {
		data = append(data, map[string]interface{}{
			"identifier": member.Identifier,
			"kind":       string(member.Kind),
			"can_delete": member.Deletable,
		})
	}
	return data
}

func getMembershipsFromSchema(values []interface{}) []cloudflare.AddressMapMembership {
	memberships := []cloudflare.AddressMapMembership{}

	for _, value := range values {
		m := value.(map[string]interface{})
		memberships = append(memberships, cloudflare.AddressMapMembership{
			Identifier: m["identifier"].(string),
			Kind:       cloudflare.AddressMapMembershipKind(m["kind"].(string)),
			Deletable:  cloudflare.BoolPtr(m["can_delete"].(bool)),
		})
	}

	return memberships
}

func getMembershipContainersFromSchema(values []interface{}) []cloudflare.AddressMapMembershipContainer {
	memberships := []cloudflare.AddressMapMembershipContainer{}

	for _, value := range values {
		m := value.(map[string]interface{})
		memberships = append(memberships, cloudflare.AddressMapMembershipContainer{
			Identifier: m["identifier"].(string),
			Kind:       cloudflare.AddressMapMembershipKind(m["kind"].(string)),
		})
	}

	return memberships
}
