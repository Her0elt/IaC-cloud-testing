package main

import (
	"cloud-testing/pkg/helpers"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		reg, err := helpers.CreateRegistry(ctx, "cloud-testing", "cloud-testing", "europe-west1", "test", "DOCKER")
		if err != nil {
			return err
		}
		url, err := helpers.UploadDockerImage(ctx, "albums_app", "albums_app", "../app", reg)
		if err != nil {
			return err
		}
		url.ApplyT(func(url string) string {
			helpers.CreateCloudRun(ctx, "album-app", "europe-west1", url, true, 100)
			return url
		})
		//err := CreateFunction(ctx, "EU", "../functions", "albums_function", "Albums function", "go116", 128, true, "GetAlbums", "europe-west1")

		return nil
	})
}
