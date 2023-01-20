package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareAccounts() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc("Data source for looking up Cloudflare Accounts."),
		ReadContext: dataSourceCloudflareAccountsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The account name to target for the resource",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Account ID.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "Account name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"type": {
							Description: "Account subscription type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"enforce_twofactor": {
							Description: "Whether 2FA is enforced on the account.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareAccountsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountName := d.Get("name").(string)

	tflog.Debug(ctx, "reading accounts")

	accounts, _, err := client.Accounts(ctx, cloudflare.AccountsListParams{Name: accountName})
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare accounts: %w", err))
	}

	accountIds := make([]string, 0)
	accountDetails := make([]interface{}, 0)

	for _, a := range accounts {
		accountDetails = append(accountDetails, map[string]interface{}{
			"id":                a.ID,
			"type":              a.Type,
			"name":              a.Name,
			"enforce_twofactor": a.Settings.EnforceTwoFactor,
		})
		accountIds = append(accountIds, a.ID)
	}

	err = d.Set("accounts", accountDetails)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error setting accounts: %w", err))
	}

	d.SetId(stringListChecksum(accountIds))
	return nil
}
