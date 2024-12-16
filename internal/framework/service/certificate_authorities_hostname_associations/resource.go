package certificate_authorities_hostname_associations

import (
	"context"
	"fmt"
	"strings"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CertificateAuthoritiesHostnameAssociationsResource{}
var _ resource.ResourceWithImportState = &CertificateAuthoritiesHostnameAssociationsResource{}

func NewResource() resource.Resource {
	return &CertificateAuthoritiesHostnameAssociationsResource{}
}

type CertificateAuthoritiesHostnameAssociationsResource struct {
	client *muxclient.Client
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_authorities_hostname_associations"
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	updatedHostnames, err := r.update(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("error updating Certificate Authorities Hostname Associations", err.Error())
		return
	}

	data = buildCertificateAuthoritiesHostnameAssociationsModel(updatedHostnames, data.MTLSCertificateID, data.ZoneID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	identifier := cfv1.ZoneIdentifier(data.ZoneID.ValueString())
	params := cfv1.ListCertificateAuthoritiesHostnameAssociationsParams{
		MTLSCertificateID: data.MTLSCertificateID.ValueString(),
	}

	hostnames, err := r.client.V1.ListCertificateAuthoritiesHostnameAssociations(ctx, identifier, params)
	if err != nil {
		resp.Diagnostics.AddError("error reading Access Mutual TLS Hostname Settings", err.Error())
		return
	}
	data = buildCertificateAuthoritiesHostnameAssociationsModel(hostnames, data.MTLSCertificateID, data.ZoneID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Helper function used by both Update, Create and Delete
func (r *CertificateAuthoritiesHostnameAssociationsResource) update(ctx context.Context, data *CertificateAuthoritiesHostnameAssociationsModel) ([]cfv1.HostnameAssociation, error) {
	updatedHostnames := []cfv1.HostnameAssociation{}
	identifier := cfv1.ZoneIdentifier(data.ZoneID.ValueString())

	hostnames := data.Hostnames
	for _, hostname := range hostnames {
		updatedHostnames = append(updatedHostnames, hostname.ValueString())
	}

	updatedCertificateAuthoritiesHostnameAssociations := cfv1.UpdateCertificateAuthoritiesHostnameAssociationsParams{
		MTLSCertificateID: data.MTLSCertificateID.ValueString(),
		Hostnames:         updatedHostnames,
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Certificate Authorities Hostname Associations from struct: %+v", updatedCertificateAuthoritiesHostnameAssociations))

	resultUpdatedHostnames, err := r.client.V1.UpdateCertificateAuthoritiesHostnameAssociations(ctx, identifier, updatedCertificateAuthoritiesHostnameAssociations)
	if err != nil {
		return nil, fmt.Errorf("error updating Certificate Authorities Hostname Associations for %s %q %s: %w", identifier.Level, identifier.Identifier, data.MTLSCertificateID, err)
	}
	return resultUpdatedHostnames, nil
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	updatedHostnames, err := r.update(ctx, data)
	if err != nil {
		resp.Diagnostics.AddError("error updating Certificate Authorities Hostname Associations", err.Error())
		return
	}

	data = buildCertificateAuthoritiesHostnameAssociationsModel(updatedHostnames, data.MTLSCertificateID, data.ZoneID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CertificateAuthoritiesHostnameAssociationsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	identifier := cfv1.ZoneIdentifier(data.ZoneID.ValueString())

	// To delete the associations we issue and update an empty array of hostnames
	deletedCertificateAuthoritiesHostnameAssociations := cfv1.UpdateCertificateAuthoritiesHostnameAssociationsParams{
		MTLSCertificateID: data.MTLSCertificateID.ValueString(),
		Hostnames:         []cfv1.HostnameAssociation{},
	}

	_, err := r.client.V1.UpdateCertificateAuthoritiesHostnameAssociations(ctx, identifier, deletedCertificateAuthoritiesHostnameAssociations)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error removing Certificate Authorities Hostname Associations for %s %q", identifier.Level, identifier.Identifier), err.Error())
		return
	}
}

func (r *CertificateAuthoritiesHostnameAssociationsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	attributes := strings.Split(req.ID, "/")

	invalidIDMessage := "invalid ID (\"%s\") specified, should be in format \"<zone_id>\" or \"<zone_id>/<mtls_certificate_id>\""
	if len(attributes) != 1 || len(attributes) != 2 {
		resp.Diagnostics.AddError("error importing Certificate Authorities Hostname Associations", fmt.Sprintf(invalidIDMessage, req.ID))
		return
	}

	zoneID := attributes[0]
	mtlsCertificateID := ""
	if len(attributes) == 2 {
		mtlsCertificateID = attributes[1]
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Certificate Authorities Hostname Associations: for %s %s", zoneID, mtlsCertificateID))

	resp.Diagnostics.Append(
		resp.State.SetAttribute(ctx, path.Root("zone_id"), zoneID)...,
	)

	resp.Diagnostics.Append(
		resp.State.SetAttribute(ctx, path.Root("mtls_certificate_id"), mtlsCertificateID)...,
	)
}

func buildCertificateAuthoritiesHostnameAssociationsModel(hostnames []cfv1.HostnameAssociation, mtlsCertificateID basetypes.StringValue, zoneID basetypes.StringValue) *CertificateAuthoritiesHostnameAssociationsModel {
	model := &CertificateAuthoritiesHostnameAssociationsModel{
		ZoneID:            zoneID,
		MTLSCertificateID: mtlsCertificateID,
	}
	for _, hostname := range hostnames {
		model.Hostnames = append(model.Hostnames, types.StringValue(hostname))
	}
	return model
}
