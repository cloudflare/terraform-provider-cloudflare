package sdkv2provider

import (
	"context"
	"fmt"
	"log"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareLogpushOwnershipChallenge() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareLogpushOwnershipChallengeSchema(),
		CreateContext: resourceCloudflareLogpushOwnershipChallengeCreate,
		UpdateContext: resourceCloudflareLogpushOwnershipChallengeCreate,
		ReadContext:   resourceCloudflareLogpushOwnershipChallengeNoop,
		DeleteContext: resourceCloudflareLogpushOwnershipChallengeNoop,
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare Logpush ownership
			challenges to use in a Logpush Job. On it's own, doesn't do much
			however this resource should be used in conjunction to create
			Logpush jobs.
		`),
	}
}

func resourceCloudflareLogpushOwnershipChallengeNoop(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceCloudflareLogpushOwnershipChallengeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	destinationConf := d.Get("destination_conf").(string)
	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	challenge, err := client.GetLogpushOwnershipChallenge(ctx, identifier, cloudflare.GetLogpushOwnershipChallengeParams{DestinationConf: destinationConf})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error requesting ownership challenge for %s: %w", identifier, err))
	}

	// The ownership challenge doesn't have a unique identifier so we generate it
	// here from the filename which will be unique.
	d.SetId(stringChecksum(challenge.Filename))
	d.Set("ownership_challenge_filename", challenge.Filename)

	log.Printf("[INFO] Created Cloudflare Logpush Ownership Challenge for %s: %s", identifier, d.Id())

	return nil
}
