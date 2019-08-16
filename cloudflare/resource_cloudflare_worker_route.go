package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareWorkerRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWorkerRouteCreate,
		Read:   resourceCloudflareWorkerRouteRead,
		Update: resourceCloudflareWorkerRouteUpdate,
		Delete: resourceCloudflareWorkerRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWorkerRouteImport,
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// Deprecated: "`zone` is deprecated in favour of explicit `zone_id` and will be removed in the next major release",
			},

			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"multi_script": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"pattern": {
				Type:     schema.TypeString,
				Required: true,
			},

			"script_name": {
				Type:     schema.TypeString,
				Optional: true,
				// enabled is used for single-script, script_name is used for multi-script
				ConflictsWith: []string{"enabled"},
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				// enabled is used for single-script, script_name is used for multi-script
				ConflictsWith: []string{"script_name"},
			},
		},
	}
}

func getRouteFromResource(d *schema.ResourceData) cloudflare.WorkerRoute {
	route := cloudflare.WorkerRoute{
		ID:      d.Id(),
		Pattern: d.Get("pattern").(string),
	}
	scriptName := d.Get("script_name").(string)
	if scriptName != "" {
		route.Script = scriptName
	} else {
		route.Enabled = d.Get("enabled").(bool)
	}
	return route
}

func resourceCloudflareWorkerRouteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	route := getRouteFromResource(d)

	zoneName := d.Get("zone").(string)
	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return fmt.Errorf("error finding zone %q: %s", zoneName, err)
	}
	d.Set("zone_id", zoneId)
	d.Set("multi_script", route.Script != "")

	log.Printf("[INFO] Creating Cloudflare Worker Route from struct: %+v", route)

	r, err := client.CreateWorkerRoute(zoneId, route)
	if err != nil {
		return errors.Wrap(err, "error creating worker route")
	}

	if r.ID == "" {
		return fmt.Errorf("failed to find id in Create response; resource was empty")
	}

	d.SetId(r.ID)

	log.Printf("[INFO] Cloudflare Worker Route ID: %s", d.Id())

	return nil
}

func resourceCloudflareWorkerRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	routeId := d.Id()

	// There isn't a dedicated endpoint for retrieving a specific route, so we
	// list all routes and find the target route by comparing IDs
	resp, err := client.ListWorkerRoutes(zoneId)

	if err != nil {
		return errors.Wrap(err, "error reading worker routes")
	}

	var route cloudflare.WorkerRoute
	for _, r := range resp.Routes {
		if r.ID == routeId {
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

	if d.Get("multi_script").(bool) {
		d.Set("script_name", route.Script)
	} else {
		d.Set("enabled", route.Enabled)
	}

	return nil
}

func resourceCloudflareWorkerRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	route := getRouteFromResource(d)

	log.Printf("[INFO] Updating Cloudflare Worker Route from struct: %+v", route)

	_, err := client.UpdateWorkerRoute(zoneId, route.ID, route)
	if err != nil {
		return errors.Wrap(err, "error updating worker route")
	}
	d.Set("multi_script", route.Script != "")

	return nil
}

func resourceCloudflareWorkerRouteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	route := getRouteFromResource(d)

	log.Printf("[INFO] Deleting Cloudflare Worker Route from zone %+v with id: %+v", zoneId, route.ID)

	_, err := client.DeleteWorkerRoute(zoneId, route.ID)
	if err != nil {
		return errors.Wrap(err, "error deleting worker route")
	}

	return nil
}

func resourceCloudflareWorkerRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	isEnterpriseWorker := false

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneName string
	var routeId string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		routeId = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneName/routeId\"", d.Id())
	}

	zoneID, err := client.ZoneIDByName(zoneName)
	routes, err := client.ListWorkerRoutes(zoneID)

	for _, r := range routes.Routes {
		if r.ID == routeId && client.OrganizationID != "" {
			isEnterpriseWorker = true
		}
	}

	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}

	d.Set("zone", zoneName)
	d.Set("zone_id", zoneID)
	d.Set("multi_script", isEnterpriseWorker)
	d.SetId(routeId)

	return []*schema.ResourceData{d}, nil
}
