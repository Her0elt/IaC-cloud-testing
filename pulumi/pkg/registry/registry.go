package registry

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi-docker/sdk/v3/go/docker"

)

func NewRepository(ctx *pulumi.Context, name string, args artifactregistry.RepositoryArgs) (*artifactregistry.Repository, error){
    reg, err := artifactregistry.NewRepository(ctx, name, &artifactregistry.RepositoryArgs{
			Location:     args.Location,
			RepositoryId: args.RepositoryId,
			Description:  args.Description,
			Format:      args.Format,
		})
		if err != nil {
			return nil, err
		}
		return reg, nil
    }

func UploadDockerImage(ctx *pulumi.Context,  name string, tagName string, app_location string, reg *artifactregistry.Repository) (*docker.Image ,error){
    imageName := pulumi.Sprintf("%s-docker.pkg.dev/%s/%s/"+tagName, reg.Location, reg.Project, reg.RepositoryId)

    registryInfo := docker.ImageRegistryArgs{}
    image, err := docker.NewImage(ctx, name, &docker.ImageArgs{
            Build:     &docker.DockerBuildArgs{Context: pulumi.String(app_location)},
            ImageName: imageName,
            Registry:  registryInfo,
        })

        // Export the base and specific version image name.
        ctx.Export("baseImageName", image.BaseImageName)
        ctx.Export("fullImageName", image.ImageName)
    if err !=nil{
        return nil, err
    }
    return image, nil
}

