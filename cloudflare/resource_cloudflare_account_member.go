package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCloudflareAccountMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccountMemberCreate,
		Read:   resourceCloudflareAccountMemberRead,
		Update: resourceCloudflareAccountMemberUpdate,
		Delete: resourceCloudflareAccountMemberDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccountMemberImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
			},

			"role_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					roleIDs := d.Get("role_ids").([]interface{})
					if arrayContains(old, roleIDs) && arrayContains(new, roleIDs) {
						return true
					}

					return false
				},
			},
		},
	}
}

func resourceCloudflareAccountMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	_, err := client.AccountMember(client.OrganizationID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Member not found") ||
			strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing account member from state because it's not present in API")
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(d.Id())

	return nil
}

func resourceCloudflareAccountMemberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting Cloudflare account member ID: %s", d.Id())

	err := client.DeleteAccountMember(client.OrganizationID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting Cloudflare account member: %s", err)
	}

	return nil
}

func resourceCloudflareAccountMemberCreate(d *schema.ResourceData, meta interface{}) error {
	memberEmailAddress := d.Get("email_address").(string)
	requestedMemberRoles := d.Get("role_ids").([]interface{})

	client := meta.(*cloudflare.API)

	var accountMemberRoleIDs []string
	for _, roleID := range requestedMemberRoles {
		accountMemberRoleIDs = append(accountMemberRoleIDs, roleID.(string))
	}

	r, err := client.CreateAccountMember(client.OrganizationID, memberEmailAddress, accountMemberRoleIDs)

	if err != nil {
		return fmt.Errorf("error creating Cloudflare account member: %s", err)
	}

	if r.ID == "" {
		return fmt.Errorf("failed to find ID in create response; resource was empty")
	}

	d.SetId(r.ID)

	return resourceCloudflareAccountMemberRead(d, meta)
}

func resourceCloudflareAccountMemberUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountRoles := []cloudflare.AccountRole{}
	memberRoles := d.Get("role_ids").([]interface{})

	for _, r := range memberRoles {
		accountRole, _ := client.AccountRole(client.OrganizationID, r.(string))
		accountRoles = append(accountRoles, accountRole)
	}

	updatedAccountMember := cloudflare.AccountMember{Roles: accountRoles}
	_, err := client.UpdateAccountMember(client.OrganizationID, d.Id(), updatedAccountMember)
	if err != nil {
		return fmt.Errorf("failed to update Cloudflare account member: %s", err)
	}

	return resourceCloudflareAccountMemberRead(d, meta)
}

func resourceCloudflareAccountMemberImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup the account member
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var accountMemberID string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		accountMemberID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"accountID/accountMemberID\" for import", d.Id())
	}

	member, err := client.AccountMember(accountID, accountMemberID)
	if err != nil {
		return nil, fmt.Errorf("unable to find account member with ID %q: %q", accountMemberID, err)
	}

	log.Printf("[INFO] Found account member: %s", member.User.Email)

	var memberIDs []string
	for _, role := range member.Roles {
		memberIDs = append(memberIDs, role.ID)
	}

	d.Set("email_address", member.User.Email)
	d.Set("role_ids", memberIDs)
	d.SetId(accountMemberID)

	return []*schema.ResourceData{d}, nil
}

func arrayContains(a string, list []interface{}) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
