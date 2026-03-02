package container_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a Cloudflare Workers Container Application.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The container application ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The account identifier to target for the resource.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"application_id": schema.StringAttribute{
				Description:   "The container application ID assigned by the API.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"script_name": schema.StringAttribute{
				Description:   "The name of the Workers script that this container is associated with.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"class_name": schema.StringAttribute{
				Description:   "The Durable Object class name that this container backs.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The container application name. Defaults to {script_name}-{class_name}.",
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"image": schema.StringAttribute{
				Description: "The container image URI (e.g. registry.example.com/image:tag).",
				Required:    true,
			},
			"max_instances": schema.Int64Attribute{
				Description: "Maximum number of container instances. Defaults to 20.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(20),
			},
			"instance_type": schema.StringAttribute{
				Description: "The instance type for the container (lite, dev, basic, standard, standard-1, standard-2, standard-3, standard-4). Mutually exclusive with vcpu/memory_mib.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("lite"),
			},
			"vcpu": schema.Float64Attribute{
				Description: "Custom vCPU allocation. Mutually exclusive with instance_type.",
				Optional:    true,
			},
			"memory_mib": schema.Int64Attribute{
				Description: "Custom memory allocation in MiB. Mutually exclusive with instance_type.",
				Optional:    true,
			},
			"disk_size_mb": schema.Int64Attribute{
				Description: "Disk size in MB.",
				Optional:    true,
			},
			"scheduling_policy": schema.StringAttribute{
				Description: "The scheduling policy (default, regional, moon, gpu, fill_metals).",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("default"),
			},
			"rollout_step_percentage": schema.Int64Attribute{
				Description: "Rollout step percentage for progressive deployments.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
			},
			"rollout_kind": schema.StringAttribute{
				Description: "Rollout kind: full_auto, full_manual, or none.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("full_auto"),
			},
			"rollout_active_grace_period": schema.Int64Attribute{
				Description: "Rollout active grace period in seconds.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"constraints": schema.SingleNestedAttribute{
				Description: "Placement constraints for container instances.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"tiers": schema.ListAttribute{
						Description: "Tier constraints (e.g. [1, 2]).",
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"regions": schema.ListAttribute{
						Description: "Region constraints.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"cities": schema.ListAttribute{
						Description: "City constraints.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"affinities": schema.SingleNestedAttribute{
				Description: "Affinity preferences for container placement.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"colocation": schema.StringAttribute{
						Description: "Colocation affinity (datacenter).",
						Optional:    true,
					},
					"hardware_generation": schema.StringAttribute{
						Description: "Hardware generation affinity (highest-overall-performance).",
						Optional:    true,
					},
				},
			},
			"observability": schema.SingleNestedAttribute{
				Description: "Observability configuration. Defaults to logs enabled.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"logs_enabled": schema.BoolAttribute{
						Description: "Whether container logs are enabled.",
						Optional:    true,
					},
				},
			},
			"wrangler_ssh": schema.SingleNestedAttribute{
				Description: "SSH configuration for the container.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether SSH is enabled.",
						Required:    true,
					},
					"port": schema.Int64Attribute{
						Description: "SSH port number.",
						Optional:    true,
					},
				},
			},
			"authorized_keys": schema.ListNestedAttribute{
				Description: "SSH authorized keys for container access.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name for the authorized key.",
							Required:    true,
						},
						"public_key": schema.StringAttribute{
							Description: "SSH public key.",
							Required:    true,
						},
					},
				},
			},
			"trusted_user_ca_keys": schema.ListNestedAttribute{
				Description: "Trusted user CA keys for SSH certificate authentication.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name for the CA key.",
							Optional:    true,
						},
						"public_key": schema.StringAttribute{
							Description: "CA public key.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (r *ContainerApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}
