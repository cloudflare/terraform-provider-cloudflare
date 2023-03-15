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
	"github.com/pkg/errors"
)

func resourceCloudflareStaticRoute() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareStaticRouteSchema(),
		CreateContext: resourceCloudflareStaticRouteCreate,
		ReadContext:   resourceCloudflareStaticRouteRead,
		UpdateContext: resourceCloudflareStaticRouteUpdate,
		DeleteContext: resourceCloudflareStaticRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareStaticRouteImport,
		},
		Description: heredoc.Doc(`
			Provides a resource, that manages Cloudflare static routes for Magic
			Transit or Magic WAN. Static routes are used to route traffic
			through GRE tunnels.
		`),
	}
}

func resourceCloudflareStaticRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	newStaticRoute, err := client.CreateMagicTransitStaticRoute(ctx, accountID, staticRouteFromResource(d))

	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating static route for prefix %s", d.Get("prefix").(string))))
	}

	d.SetId(newStaticRoute[0].ID)

	return resourceCloudflareStaticRouteRead(ctx, d, meta)
}

func resourceCloudflareStaticRouteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/routeID\"", d.Id())
	}

	accountID, routeID := attributes[0], attributes[1]
	d.SetId(routeID)
	d.Set(consts.AccountIDSchemaKey, accountID)

	resourceCloudflareStaticRouteRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareStaticRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	staticRoute, err := client.GetMagicTransitStaticRoute(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Route not found") {
			tflog.Info(ctx, fmt.Sprintf("Static Route %s not found", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading Static Route ID %q", d.Id())))
	}

	d.Set("prefix", staticRoute.Prefix)
	d.Set("nexthop", staticRoute.Nexthop)
	d.Set("priority", staticRoute.Priority)
	d.Set("weight", staticRoute.Weight)

	if len(staticRoute.Description) > 0 {
		d.Set("description", staticRoute.Description)
	}

	if len(staticRoute.Scope.ColoRegions) > 0 {
		d.Set("colo_regions", staticRoute.Scope.ColoRegions)
	}

	if len(staticRoute.Scope.ColoNames) > 0 {
		d.Set("colo_names", staticRoute.Scope.ColoNames)
	}

	return nil
}

func resourceCloudflareStaticRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	_, err := client.UpdateMagicTransitStaticRoute(ctx, accountID, d.Id(), staticRouteFromResource(d))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating static route with ID %q", d.Id())))
	}

	return resourceCloudflareStaticRouteRead(ctx, d, meta)
}

func resourceCloudflareStaticRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Static Route:  %s", d.Id()))

	_, err := client.DeleteMagicTransitStaticRoute(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Static Route: %w", err))
	}

	return nil
}

func staticRouteFromResource(d *schema.ResourceData) cloudflare.MagicTransitStaticRoute {
	staticRoute := cloudflare.MagicTransitStaticRoute{
		Prefix:   d.Get("prefix").(string),
		Nexthop:  d.Get("nexthop").(string),
		Priority: d.Get("priority").(int),
	}

	description, descriptionOk := d.GetOk("description")
	if descriptionOk {
		staticRoute.Description = description.(string)
	}

	weight, weightOk := d.GetOk("weight")
	if weightOk {
		staticRoute.Weight = weight.(int)
	}

	coloRegions, coloRegionsOk := d.GetOk("colo_regions")
	if coloRegionsOk {
		staticRoute.Scope.ColoRegions = expandInterfaceToStringList(coloRegions)
	}

	coloNames, coloNamesOk := d.GetOk("colo_names")
	if coloNamesOk {
		staticRoute.Scope.ColoNames = expandInterfaceToStringList(coloNames)
	}

	return staticRoute
}
