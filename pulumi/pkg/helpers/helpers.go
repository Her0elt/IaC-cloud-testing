package helpers


import (
	local_cloudrun "cloud-testing/pkg/cloudrun"
	"cloud-testing/pkg/function"
	"cloud-testing/pkg/registry"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateFunction(ctx *pulumi.Context, bucket_location string, function_location string, function_name, description string, runtime string, mem int, http bool, entryPoint string, region string) error {
	_, err := function.NewFunction(ctx, bucket_location, function_location, function_name,
		cloudfunctions.FunctionArgs{
			Description:       pulumi.String(description),
			Runtime:           pulumi.String(runtime),
			AvailableMemoryMb: pulumi.Int(mem),
			TriggerHttp:       pulumi.Bool(http),
			EntryPoint:        pulumi.String(entryPoint),
			Region:            pulumi.String(region),
		})
	if err != nil {
		return err
	}
	return nil
}

func CreateCloudRun(ctx *pulumi.Context, app_name string, app_location string, image_location string, latestRevision bool, percent int) error {
	_, err := local_cloudrun.NewService(ctx, app_name, cloudrun.ServiceArgs{
		Location: pulumi.String(app_location),
		Template: &cloudrun.ServiceTemplateArgs{
			Spec: &cloudrun.ServiceTemplateSpecArgs{
				Containers: cloudrun.ServiceTemplateSpecContainerArray{
					&cloudrun.ServiceTemplateSpecContainerArgs{
						Image: pulumi.String(image_location),
					},
				},
			},
		},
		Traffics: cloudrun.ServiceTrafficArray{
			&cloudrun.ServiceTrafficArgs{
				LatestRevision: pulumi.Bool(latestRevision),
				Percent:        pulumi.Int(percent),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateRegistry(ctx *pulumi.Context, repoName string, repoId string, repoLocation string, repoDisc string, repoFormat string) (*artifactregistry.Repository, error) {
	reg, err := registry.NewRepository(ctx, repoName, artifactregistry.RepositoryArgs{
		Location:     pulumi.String(repoLocation),
		RepositoryId: pulumi.String(repoId),
		Description:  pulumi.String(repoDisc),
		Format:       pulumi.String(repoFormat),
	})
	if err != nil {
		return nil, err
	}
	return reg, nil
}

func UploadDockerImage(ctx *pulumi.Context, imageName string, tagName string, appLocation string, reg *artifactregistry.Repository) (pulumi.StringOutput, error) {
	image, err := registry.UploadDockerImage(ctx, imageName, tagName, appLocation, reg)
	if err != nil {
		return pulumi.String("").ToStringOutput(), err
	}
	return image.BaseImageName, nil
}

