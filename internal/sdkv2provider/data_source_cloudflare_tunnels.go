package sdkv2provider

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareTunnels() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc(`
			Use this data source to lookup a single [Cloudflare Tunnel](https://developers.cloudflare.com/api/operations/cloudflare-tunnel-get-a-cloudflare-tunnel).
		`),
		ReadContext: dataSourceCloudflareTunnelsRead,
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"include_token": {
				Description: "Whether to include the tunnel token in the response.",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel_id": {
							Description: "UUID of the tunnel",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": {
							Description: "User-friendly name of the tunnel.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"status": {
							Description: "Current status of the tunnel. One of: inactive, degraded, healthy, down",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"remote_config": {
							Description: "Whether the tunnel can be configured remotely from the Zero Trust dashboard.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},

			"tunnels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Description: "Current status of the tunnel. One of: inactive, degraded, healthy, down",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "User-friendly name of the tunnel.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"remote_config": {
							Description: "Whether the tunnel can be configured remotely from the Zero Trust dashboard.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"cname": {
							Description: "Usable CNAME for accessing the tunnel.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tunnel_token": {
							Description: "Token used by connector to authenticate and run the tunnel.",
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Optional:    true,
						},
						"created_at": {
							Description: "Timestamp of when the tunnel was created.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"deleted_at": {
							Description: "Timestamp of when the tunnel was deleted.",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareTunnelsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	filter, err := buildTunnelFilter(d.Get("filter"))
	if err != nil {
		return diag.FromErr(err)
	}

	listParams := cloudflare.TunnelListParams{}
	listParams.PerPage = 1000

	tunnels, _, err := client.ListTunnels(ctx, accountID, listParams)
	tflog.Info(ctx, fmt.Sprintf("found tunnels %s", "err"))

	if err != nil {
		return diag.FromErr(err)
	}

	filteredTunnels := make([]interface{}, 0)
	filteredTunnelIds := make([]string, 0)

	for _, tunnel := range tunnels {
		if matchTunnelFilter(filter, tunnel) {
			result := map[string]interface{}{
				"name":          tunnel.Name,
				"status":        tunnel.Status,
				"remote_config": tunnel.RemoteConfig,
				"cname":         fmt.Sprintf("%s.%s", tunnel.ID, argoTunnelCNAME),
				"created_at":    tunnel.CreatedAt.Format(time.RFC3339),
			}

			if tunnel.DeletedAt != nil {
				result["deleted_at"] = tunnel.DeletedAt.Format(time.RFC3339)
			}

			if d.Get("include_token").(bool) {
				token, err := client.GetTunnelToken(ctx, accountID, tunnel.ID)
				if err != nil {
					return diag.FromErr(fmt.Errorf("error fetching tunnel token %s: %w", tunnel.ID, err))
				}

				result["tunnel_token"] = token
			}

			filteredTunnelIds = append(filteredTunnelIds, tunnel.ID)
			filteredTunnels = append(filteredTunnels, result)

		} else {
			tflog.Debug(ctx, fmt.Sprintf("tunnel %s did not match filter", tunnel.ID))
		}
	}

	err = d.Set("tunnels", filteredTunnels)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(stringListChecksum(filteredTunnelIds))

	return nil
}

func matchTunnelFilter(filter *searchFilterTunnels, tunnel cloudflare.Tunnel) bool {
	if filter.Name != nil && !filter.Name.MatchString(tunnel.Name) {
		return false
	}

	if filter.TunnelID != "" && filter.TunnelID != tunnel.ID {
		return false
	}

	if filter.Status != "" && filter.Status != tunnel.Status {
		return false
	}

	if filter.RemoteConfig && !tunnel.RemoteConfig {
		return false
	}

	return true
}

func buildTunnelFilter(d interface{}) (*searchFilterTunnels, error) {
	config := d.([]interface{})
	filter := &searchFilterTunnels{}

	if len(config) == 0 || config[0] == nil {
		return filter, nil
	}

	cfg := config[0].(map[string]interface{})

	for k, v := range cfg {
		switch k {

		case "name":
			pattern, err := regexp.Compile(v.(string))
			if err != nil {
				return nil, err
			}

			filter.Name = pattern

		case "tunnel_id":
			filter.TunnelID = v.(string)

		case "status":
			filter.Status = v.(string)

		case "remote_config":
			filter.RemoteConfig = v.(bool)

		}
	}

	return filter, nil
}

type searchFilterTunnels struct {
	Name         *regexp.Regexp
	TunnelID     string
	Status       string
	RemoteConfig bool
}
