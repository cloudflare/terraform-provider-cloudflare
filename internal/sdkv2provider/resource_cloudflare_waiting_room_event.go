package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoomEvent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareWaitingRoomEventCreate,
		ReadContext:   resourceCloudflareWaitingRoomEventRead,
		UpdateContext: resourceCloudflareWaitingRoomEventUpdate,
		DeleteContext: resourceCloudflareWaitingRoomEventDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWaitingRoomEventImport,
		},

		Schema:      resourceCloudflareWaitingRoomEventSchema(),
		Description: "Provides a Cloudflare Waiting Room Event resource.",
	}
}

func expandWaitingRoomEvent(d *schema.ResourceData) (cloudflare.WaitingRoomEvent, error) {
	var disableSessionRenewal *bool
	if b, ok := d.GetOk("disable_session_renewal"); ok {
		disableSessionRenewal = cloudflare.BoolPtr(b.(bool))
	}

	eventStartTime, err := time.Parse(time.RFC3339, d.Get("event_start_time").(string))
	if err != nil {
		return cloudflare.WaitingRoomEvent{}, err
	}

	eventEndTime, err := time.Parse(time.RFC3339, d.Get("event_end_time").(string))
	if err != nil {
		return cloudflare.WaitingRoomEvent{}, err
	}

	var prequeueStartTime *time.Time
	if t, ok := d.GetOk("prequeue_start_time"); ok {
		prequeueStartTimeValue, err := time.Parse(time.RFC3339, t.(string))
		prequeueStartTime = cloudflare.TimePtr(prequeueStartTimeValue)
		if err != nil {
			return cloudflare.WaitingRoomEvent{}, err
		}
	}

	return cloudflare.WaitingRoomEvent{
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

func resourceCloudflareWaitingRoomEventCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	waitingRoomEventName := d.Get("name").(string)
	newWaitingRoomEvent, err := expandWaitingRoomEvent(d)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error building waiting room event %q: %w", waitingRoomEventName, err))
	}

	waitingRoomEvent, err := client.CreateWaitingRoomEvent(ctx, zoneID, waitingRoomID, newWaitingRoomEvent)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating waiting room event %q: %w", waitingRoomEventName, err))
	}

	d.SetId(waitingRoomEvent.ID)

	return resourceCloudflareWaitingRoomEventRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomEventRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoomEvent, err := client.WaitingRoomEvent(ctx, zoneID, waitingRoomID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing waiting room event from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error getting waiting room event %q: %w", d.Get("name").(string), err))
	}

	d.Set("name", waitingRoomEvent.Name)
	d.Set("event_start_time", waitingRoomEvent.EventStartTime.Format(time.RFC3339))
	d.Set("event_end_time", waitingRoomEvent.EventEndTime.Format(time.RFC3339))
	d.Set("session_duration", waitingRoomEvent.SessionDuration)

	if waitingRoomEvent.PrequeueStartTime != nil {
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

	if waitingRoomEvent.DisableSessionRenewal != nil {
		d.Set("disable_session_renewal", waitingRoomEvent.DisableSessionRenewal)
	}

	return nil
}

func resourceCloudflareWaitingRoomEventUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomEventID := d.Id()
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	waitingRoomEventName := d.Get("name").(string)
	waitingRoomEvent, err := expandWaitingRoomEvent(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error building waiting room event %q: %w", waitingRoomEventName, err))
	}
	waitingRoomEvent.ID = waitingRoomEventID

	_, err = client.ChangeWaitingRoomEvent(ctx, zoneID, waitingRoomID, waitingRoomEvent)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating waiting room event %q: %w", waitingRoomEventName, err))
	}

	return resourceCloudflareWaitingRoomEventRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomEventDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomEventID := d.Id()
	waitingRoomID := d.Get("waiting_room_id").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteWaitingRoomEvent(ctx, zoneID, waitingRoomID, waitingRoomEventID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting waiting room event %q: %w", d.Get("name").(string), err))
	}

	return nil
}

func resourceCloudflareWaitingRoomEventImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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

	waitingRoomEvent, err := client.WaitingRoomEvent(ctx, zoneID, waitingRoomID, waitingRoomEventID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Waiting room event %s", waitingRoomID)
	}

	d.SetId(waitingRoomEvent.ID)
	d.Set("waiting_room_id", waitingRoomID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)

	resourceCloudflareWaitingRoomEventRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
