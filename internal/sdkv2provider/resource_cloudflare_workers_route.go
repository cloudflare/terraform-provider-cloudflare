package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerRoute() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWorkerRouteSchema(),
		CreateContext: resourceCloudflareWorkerRouteCreate,
		ReadContext:   resourceCloudflareWorkerRouteRead,
		UpdateContext: resourceCloudflareWorkerRouteUpdate,
		DeleteContext: resourceCloudflareWorkerRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWorkerRouteImport,
		},
		Description: heredoc.Doc("Provides a Cloudflare worker route resource. A route will also require a `cloudflare_worker_script`."),
	}
}

func resourceCloudflareWorkerRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	params := cloudflare.CreateWorkerRouteParams{
		Pattern: d.Get("pattern").(string),
		Script:  d.Get("script_name").(string),
	}

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Worker Route from struct: %+v", params))

	r, err := client.CreateWorkerRoute(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating worker route"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in Create response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Worker Route ID: %s", d.Id()))

	return nil
}

func resourceCloudflareWorkerRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	routeID := d.Id()

	// There isn't a dedicated endpoint for retrieving a specific route, so we
	// list all routes and find the target route by comparing IDs
	resp, err := client.ListWorkerRoutes(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListWorkerRoutesParams{})

	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error reading worker routes"))
	}

	var route cloudflare.WorkerRoute
	for _, r := range resp.Routes {
		if r.ID == routeID {
			route = r
			break
		}
	}

	// If the resource is deleted, we should set the ID to "" and not
	// return an error according to the terraform spec
	if route.ID == "" {
		d.SetId("")
		return nil
	}

	d.Set("pattern", route.Pattern)
	d.Set("script_name", route.ScriptName)

	return nil
}

func resourceCloudflareWorkerRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	params := cloudflare.UpdateWorkerRouteParams{
		ID:      d.Id(),
		Pattern: d.Get("pattern").(string),
		Script:  d.Get("script_name").(string),
	}

	log.Printf("[INFO] Updating Cloudflare Worker Route from struct: %+v", params)

	_, err := client.UpdateWorkerRoute(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error updating worker route"))
	}

	return nil
}

func resourceCloudflareWorkerRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	log.Printf("[INFO] Deleting Cloudflare Worker Route from zone %+v with id: %+v", zoneID, d.Id())

	_, err := client.DeleteWorkerRoute(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error deleting worker route"))
	}

	return nil
}

func resourceCloudflareWorkerRouteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var routeID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		routeID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/routeID\"", d.Id())
	}

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(routeID)

	resourceCloudflareWorkerRouteRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
