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

		Schema: resourceCloudflareWaitingRoomEventSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
	}
}

func expandWaitingRoomEvent(d *schema.ResourceData) (cloudflare.WaitingRoomEvent, error) {
	var disableSessionRenewal *bool
	switch b := d.Get("disable_session_renewal").(type) {
	case bool:
		disableSessionRenewal = &b
	case nil:
		disableSessionRenewal = nil
	}
	eventStartTime, err := time.Parse(time.RFC3339, d.Get("event_start_time").(string))
	if err != nil {
		return cloudflare.WaitingRoomEvent{}, err
	}

	eventEndTime, err := time.Parse(time.RFC3339, d.Get("event_end_time").(string))
	if err != nil {
		return cloudflare.WaitingRoomEvent{}, err
	}

	prequeueStartTime := time.Time{}
	if t, ok := d.GetOk("prequeue_start_time"); ok {
		prequeueStartTime, err = time.Parse(time.RFC3339, t.(string))
		if err != nil {
			return cloudflare.WaitingRoomEvent{}, err
		}
	}

	return cloudflare.WaitingRoomEvent{
		ID:                    d.Id(),
		Name:                  d.Get("name").(string),
		EventStartTime:        eventStartTime,
		EventEndTime:          eventEndTime,
		PrequeueStartTime:     prequeueStartTime,
		Description:           d.Get("description").(string),
		QueueingMethod:        d.Get("queueing_method").(string),
		ShuffleAtEventStart:   d.Get("shuffle_at_event_start").(bool),
		Suspended:             d.Get("suspended").(bool),
		TotalActiveUsers:      d.Get("total_active_users").(int),
		NewUsersPerMinute:     d.Get("new_users_per_minute").(int),
		CustomPageHTML:        d.Get("custom_page_html").(string),
		SessionDuration:       d.Get("session_duration").(int),
		DisableSessionRenewal: disableSessionRenewal,
	}, nil
}

func resourceCloudflareWaitingRoomEventCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get("zone_id").(string)

	newWaitingRoomEvent, err := expandWaitingRoomEvent(d)
	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error building waiting room event %q: %w", name, err)
	}

	waitingRoomEvent, err := client.CreateWaitingRoomEvent(context.Background(), zoneID, waitingRoomID, newWaitingRoomEvent)

	if err != nil {
		return fmt.Errorf("error creating waiting room event %q: %w", d.Get("name").(string), err)
	}

	d.SetId(waitingRoomEvent.ID)

	return resourceCloudflareWaitingRoomEventRead(d, meta)
}

func resourceCloudflareWaitingRoomEventRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
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
		return fmt.Errorf("error getting waiting room event %q: %w", name, err)
	}
	d.Set("name", waitingRoomEvent.Name)
	d.Set("event_start_time", waitingRoomEvent.EventStartTime.Format(time.RFC3339))
	d.Set("event_end_time", waitingRoomEvent.EventEndTime.Format(time.RFC3339))
	if !waitingRoomEvent.PrequeueStartTime.IsZero() {
		d.Set("prequeue_start_time", waitingRoomEvent.PrequeueStartTime.Format(time.RFC3339))
	}
	
	if waitingRoomEvent.Description != "" {
		d.Set("description", waitingRoomEvent.Description)
	}
	
	if waitingRoomEvent.QueueingMethod != "" {
		d.Set("queueing_method", waitingRoomEvent.QueueingMethod)
	}
	
	d.Set("shuffle_at_event_start", waitingRoomEvent.ShuffleAtEventStart)
	d.Set("suspended", waitingRoomEvent.Suspended)
	
	if waitingRoomEvent.TotalActiveUsers != 0 {
		d.Set("total_active_users", waitingRoomEvent.TotalActiveUsers)
	}
	
	if waitingRoomEvent.NewUsersPerMinute != 0 {
		d.Set("new_users_per_minute", waitingRoomEvent.NewUsersPerMinute)
	}
	
	if waitingRoomEvent.CustomPageHTML != "" {
		d.Set("custom_page_html", waitingRoomEvent.CustomPageHTML)
	}
	d.Set("session_duration", waitingRoomEvent.SessionDuration)
	if waitingRoomEvent.DisableSessionRenewal != nil {
		d.Set("disable_session_renewal", waitingRoomEvent.DisableSessionRenewal)
	}
	return nil
}

func resourceCloudflareWaitingRoomEventUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get("zone_id").(string)

	waitingRoomEvent, err := expandWaitingRoomEvent(d)
	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error building waiting room event %q: %w", name, err)
	}

	_, err = client.ChangeWaitingRoomEvent(context.Background(), zoneID, waitingRoomID, waitingRoomEvent)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("error updating waiting room event %q: %w", name, err)
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
		return fmt.Errorf("error deleting waiting room event %q: %w", name, err)
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

	err = resourceCloudflareWaitingRoomEventRead(d, meta)
	return []*schema.ResourceData{d}, err
}
