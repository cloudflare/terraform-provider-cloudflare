package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoom() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareWaitingRoomCreate,
		ReadContext:   resourceCloudflareWaitingRoomRead,
		UpdateContext: resourceCloudflareWaitingRoomUpdate,
		DeleteContext: resourceCloudflareWaitingRoomDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWaitingRoomImport,
		},

		Schema: resourceCloudflareWaitingRoomSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
	}
}

func buildWaitingRoom(d *schema.ResourceData) cloudflare.WaitingRoom {
	return cloudflare.WaitingRoom{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		Suspended:             d.Get("suspended").(bool),
		Host:                  d.Get("host").(string),
		Path:                  d.Get("path").(string),
		TotalActiveUsers:      d.Get("total_active_users").(int),
		NewUsersPerMinute:     d.Get("new_users_per_minute").(int),
		CustomPageHTML:        d.Get("custom_page_html").(string),
		SessionDuration:       d.Get("session_duration").(int),
		JsonResponseEnabled:   d.Get("json_response_enabled").(bool),
		QueueAll:              d.Get("queue_all").(bool),
		DisableSessionRenewal: d.Get("disable_session_renewal").(bool),
	}
}

func resourceCloudflareWaitingRoomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	newWaitingRoom := buildWaitingRoom(d)

	waitingRoom, err := client.CreateWaitingRoom(context.Background(), zoneID, newWaitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error creating waiting room %q: %s", name, err))
	}

	d.SetId(waitingRoom.ID)

	return resourceCloudflareWaitingRoomRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	waitingRoom, err := client.WaitingRoom(context.Background(), zoneID, waitingRoomID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing waiting room from state because it's not found in API")
			d.SetId("")
			return nil
		}
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error getting waiting room %q: %s", name, err))
	}
	d.SetId(waitingRoom.ID)
	d.Set("name", waitingRoom.Name)
	d.Set("description", waitingRoom.Description)
	d.Set("suspended", waitingRoom.Suspended)
	d.Set("host", waitingRoom.Host)
	d.Set("path", waitingRoom.Path)
	d.Set("queue_all", waitingRoom.QueueAll)
	d.Set("new_users_per_minute", waitingRoom.NewUsersPerMinute)
	d.Set("total_active_users", waitingRoom.TotalActiveUsers)
	d.Set("session_duration", waitingRoom.SessionDuration)
	d.Set("disable_session_renewal", waitingRoom.DisableSessionRenewal)
	d.Set("custom_page_html", waitingRoom.CustomPageHTML)
	d.Set("json_response_enabled", waitingRoom.JsonResponseEnabled)
	return nil
}

func resourceCloudflareWaitingRoomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	waitingRoom := buildWaitingRoom(d)

	_, err := client.ChangeWaitingRoom(context.Background(), zoneID, waitingRoomID, waitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error updating waiting room %q: %s", name, err))
	}

	return resourceCloudflareWaitingRoomRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteWaitingRoom(context.Background(), zoneID, waitingRoomID)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error deleting waiting room %q: %s", name, err))
	}

	return nil
}

func resourceCloudflareWaitingRoomImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var waitingRoomID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		waitingRoomID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/waitingRoomID\" for import", d.Id())
	}

	waitingRoom, err := client.WaitingRoom(context.Background(), zoneID, waitingRoomID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Waiting room %s", waitingRoomID)
	}

	d.SetId(waitingRoom.ID)
	d.Set("zone_id", zoneID)

	resourceCloudflareWaitingRoomRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
