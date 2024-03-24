// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	appstream_sdkv2 "github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDirectoryConfig,
			TypeName: "aws_appstream_directory_config",
		},
		{
			Factory:  ResourceFleet,
			TypeName: "aws_appstream_fleet",
			Name:     "Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFleetStackAssociation,
			TypeName: "aws_appstream_fleet_stack_association",
		},
		{
			Factory:  ResourceImageBuilder,
			TypeName: "aws_appstream_image_builder",
			Name:     "Image Builder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceStack,
			TypeName: "aws_appstream_stack",
			Name:     "Stack",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_appstream_user",
		},
		{
			Factory:  ResourceUserStackAssociation,
			TypeName: "aws_appstream_user_stack_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppStream
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appstream_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return appstream_sdkv2.NewFromConfig(cfg, func(o *appstream_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
