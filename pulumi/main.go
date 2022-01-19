package main

import (
	local_cloudrun "cloud-testing/pkg/cloudrun"
	"cloud-testing/pkg/function"
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
func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		err := CreateCloudRun(ctx, "album-app", "europe-west1", "europe-west1-docker.pkg.dev/hello-cloud-338214/cloud-testing/albums_app", true, 100)
		//err := CreateFunction(ctx, "EU", "../functions", "albums_function", "Albums function", "go116", 128, true, "GetAlbums", "europe-west1")
		if err != nil {
			return err
		}
		return nil
	})
}
