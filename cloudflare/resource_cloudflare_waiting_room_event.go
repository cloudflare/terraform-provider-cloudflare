package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomEvent() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWaitingRoomEventCreate,
		Read:   resourceCloudflareWaitingRoomEventRead,
		Update: resourceCloudflareWaitingRoomEventUpdate,
		Delete: resourceCloudflareWaitingRoomEventDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWaitingRoomEventImport,
		},

		Schema: resourceCloudflareWaitingRoomSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
	}
}

func buildWaitingRoomEvent(d *schema.ResourceData) cloudflare.WaitingRoomEvent {
	disableSessionRenewal := d.Get("disable_session_renewal").(bool)
	return cloudflare.WaitingRoomEvent{
		ID:                    d.Id(),
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		QueueingMethod:        d.Get("queueing_method").(string),
		ShuffleAtEventStart:   d.Get("shuffle_at_evend_start").(bool),
		Suspended:             d.Get("suspended").(bool),
		TotalActiveUsers:      d.Get("total_active_users").(int),
		NewUsersPerMinute:     d.Get("new_users_per_minute").(int),
		CustomPageHTML:        d.Get("custom_page_html").(string),
		SessionDuration:       d.Get("session_duration").(int),
		DisableSessionRenewal: &disableSessionRenewal,
	}
}

func resourceCloudflareWaitingRoomEventCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get("zone_id").(string)

	newWaitingRoomEvent := buildWaitingRoomEvent(d)

	waitingRoomEvent, err := client.CreateWaitingRoomEvent(context.Background(), zoneID, waitingRoomID, newWaitingRoomEvent)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error creating waiting room event %q: %s", name, err)
	}

	d.SetId(waitingRoomEvent.ID)

	return resourceCloudflareWaitingRoomEventRead(d, meta)
}

func resourceCloudflareWaitingRoomEventRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("zone_id").(string)
	zoneID := d.Get("zone_id").(string)
	waitingRoomEventID := d.Id()

	waitingRoomEvent, err := client.WaitingRoomEvent(context.Background(), zoneID, waitingRoomID, waitingRoomEventID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing waiting room event from state because it's not found in API")
			d.SetId("")
			return nil
		}
		name := d.Get("name").(string)
		return fmt.Errorf("error getting waiting room event %q: %s", name, err)
	}
	d.Set("name", waitingRoomEvent.Name)
	d.Set("description", waitingRoomEvent.Description)
	d.Set("suspended", waitingRoomEvent.Suspended)
	d.Set("new_users_per_minute", waitingRoomEvent.NewUsersPerMinute)
	d.Set("total_active_users", waitingRoomEvent.TotalActiveUsers)
	d.Set("session_duration", waitingRoomEvent.SessionDuration)
	d.Set("disable_session_renewal", waitingRoomEvent.DisableSessionRenewal)
	d.Set("custom_page_html", waitingRoomEvent.CustomPageHTML)
	return nil
}

func resourceCloudflareWaitingRoomEventUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get("zone_id").(string)

	waitingRoomEvent := buildWaitingRoomEvent(d)

	_, err := client.ChangeWaitingRoomEvent(context.Background(), zoneID, waitingRoomID, waitingRoomEvent)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error updating waiting room event %q: %s", name, err)
	}

	return resourceCloudflareWaitingRoomEventRead(d, meta)
}

func resourceCloudflareWaitingRoomEventDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomEventID := d.Id()
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteWaitingRoomEvent(context.Background(), zoneID, waitingRoomID, waitingRoomEventID)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error deleting waiting room event %q: %s", name, err)
	}

	return nil
}

func resourceCloudflareWaitingRoomEventImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 3)
	var zoneID string
	var waitingRoomID string
	var waitingRoomEventID string
	if len(idAttr) == 3 {
		zoneID = idAttr[0]
		waitingRoomID = idAttr[1]
		waitingRoomEventID = idAttr[3]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/waitingRoomID/eventID\" for import", d.Id())
	}

	waitingRoomEvent, err := client.WaitingRoomEvent(context.Background(), zoneID, waitingRoomID, waitingRoomEventID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Waiting room event %s", waitingRoomID)
	}

	d.SetId(waitingRoomEvent.ID)
	d.Set("waiting_room_id", waitingRoomID)
	d.Set("zone_id", zoneID)

	resourceCloudflareWaitingRoomEventRead(d, meta)
	return []*schema.ResourceData{d}, nil
}
