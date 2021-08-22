package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareStaticRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareStaticRouteCreate,
		Read:   resourceCloudflareStaticRouteRead,
		Update: resourceCloudflareStaticRouteUpdate,
		Delete: resourceCloudflareStaticRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareStaticRouteImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nexthop": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Required: false,
				// API does not allow to reset weights when attribute isn't send. To avoid generating unnecessary changes
				// we will trigger a re-create when weights change
				ForceNew: true,
			},
			"colo_regions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"colo_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceCloudflareStaticRouteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	newStaticRoute, err := client.CreateMagicTransitStaticRoute(context.Background(), staticRouteFromResource(d))

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating static route for prefix %s", d.Get("prefix").(string)))
	}

	d.SetId(newStaticRoute[0].ID)

	return resourceCloudflareStaticRouteRead(d, meta)
}

func resourceCloudflareStaticRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/routeID\"", d.Id())
	}

	accountID, routeID := attributes[0], attributes[1]
	d.SetId(routeID)
	d.Set("account_id", accountID)
	client.AccountID = accountID

	resourceCloudflareStaticRouteRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareStaticRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	staticRoute, err := client.GetMagicTransitStaticRoute(context.Background(), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Route not found") {
			log.Printf("[INFO] Static Route %s not found", d.Id())
			d.SetId("")
			return nil
		}
		return errors.Wrap(err, fmt.Sprintf("error reading Static Route ID %q", d.Id()))
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

func resourceCloudflareStaticRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	_, err := client.UpdateMagicTransitStaticRoute(context.Background(), d.Id(), staticRouteFromResource(d))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating static route with ID %q", d.Id()))
	}

	return resourceCloudflareStaticRouteRead(d, meta)
}

func resourceCloudflareStaticRouteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	log.Printf("[INFO] Deleting Static Route:  %s", d.Id())

	_, err := client.DeleteMagicTransitStaticRoute(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Static Route: %s", err)
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
