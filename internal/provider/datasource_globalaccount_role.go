package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/SAP/terraform-provider-btp/internal/btpcli"
)

func newGlobalaccountRoleDataSource() datasource.DataSource {
	return &globalaccountRoleDataSource{}
}

type globalaccountRoleScope struct {
	Name                         types.String `tfsdk:"name"`
	Description                  types.String `tfsdk:"description"`
	CustomGrantAsAuthorityToApps types.Set    `tfsdk:"custom_grant_as_authority_to_apps"`
	CustomGrantedApps            types.Set    `tfsdk:"custom_granted_apps"`
	GrantAsAuthorityToApps       types.Set    `tfsdk:"grant_as_authority_to_apps"`
	GrantedApps                  types.Set    `tfsdk:"granted_apps"`
}

type globalaccountRoleDataSourceConfig struct {
	/* INPUT */
	Id                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	RoleTemplateAppId types.String `tfsdk:"app_id"`
	RoleTemplateName  types.String `tfsdk:"role_template_name"`
	/* OUTPUT */
	Description types.String          `tfsdk:"description"`
	IsReadOnly  types.Bool            `tfsdk:"read_only"`
	Scopes      []subaccountRoleScope `tfsdk:"scopes"`
}

type globalaccountRoleDataSource struct {
	cli *btpcli.ClientFacade
}

func (ds *globalaccountRoleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_globalaccount_role", req.ProviderTypeName)
}

func (ds *globalaccountRoleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	ds.cli = req.ProviderData.(*btpcli.ClientFacade)
}

func (ds *globalaccountRoleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Gets details about a specific global account role.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{ // required hashicorps terraform plugin testing framework
				DeprecationMessage:  "Use the `btp_globalaccount` datasource instead",
				MarkdownDescription: "The ID of the global account",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the role.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"app_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the xsuaa application.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"role_template_name": schema.StringAttribute{
				MarkdownDescription: "The name of the role template.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the role.",
				Computed:            true,
			},
			"read_only": schema.BoolAttribute{
				MarkdownDescription: "Shows whether the role can be modified or not.",
				Computed:            true,
			},
			"scopes": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the scope.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "The description of the scope.",
							Computed:            true,
						},
						"custom_grant_as_authority_to_apps": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"custom_granted_apps": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"grant_as_authority_to_apps": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"granted_apps": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
				},
				MarkdownDescription: "The scopes available with this role.",
				Computed:            true,
			},
		},
	}
}

func (ds *globalaccountRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data globalaccountRoleDataSourceConfig

	diags := req.Config.Get(ctx, &data)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cliRes, _, err := ds.cli.Security.Role.GetByGlobalAccount(ctx, data.Name.ValueString(), data.RoleTemplateAppId.ValueString(), data.RoleTemplateName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("API Error Reading Resource Role (Global Account)", fmt.Sprintf("%s", err))
		return
	}

	data.Id = types.StringValue(ds.cli.GetGlobalAccountSubdomain())
	data.Description = types.StringValue(cliRes.Description)
	data.IsReadOnly = types.BoolValue(cliRes.IsReadOnly)
	data.Scopes = []subaccountRoleScope{}

	for _, scope := range cliRes.Scopes {
		scopeVal := subaccountRoleScope{
			Name:        types.StringValue(scope.Name),
			Description: types.StringValue(scope.Description),
		}

		scopeVal.CustomGrantAsAuthorityToApps, diags = types.SetValueFrom(ctx, types.StringType, scope.CustomGrantAsAuthorityToApps)
		resp.Diagnostics.Append(diags...)

		scopeVal.CustomGrantedApps, diags = types.SetValueFrom(ctx, types.StringType, scope.CustomGrantedApps)
		resp.Diagnostics.Append(diags...)

		scopeVal.GrantAsAuthorityToApps, diags = types.SetValueFrom(ctx, types.StringType, scope.GrantAsAuthorityToApps)
		resp.Diagnostics.Append(diags...)

		scopeVal.GrantedApps, diags = types.SetValueFrom(ctx, types.StringType, scope.GrantedApps)
		resp.Diagnostics.Append(diags...)

		data.Scopes = append(data.Scopes, scopeVal)
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
