package cloudrun

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewService(ctx *pulumi.Context, name string, args cloudrun.ServiceArgs) (*cloudrun.Service, error){
    svc, err := cloudrun.NewService(ctx, name, &cloudrun.ServiceArgs{
			Location: args.Location,
			Template: args.Template,
            Traffics: args.Traffics,
        })
		if err != nil {
			return nil, err
		}

        _, err = cloudrun.NewIamMember(ctx, name + "_invoker", &cloudrun.IamMemberArgs{
			Location: svc.Location,
			Service: svc.Name,
    		Role:     pulumi.String("roles/run.invoker"),
			Member:   pulumi.String("allUsers"),
		})
		if err != nil {
			return nil, err
		}
		return svc, nil

}



