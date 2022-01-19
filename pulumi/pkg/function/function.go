package function

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudfunctions"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "cloud-testing/pkg/bucket"

)


func NewFunction(ctx *pulumi.Context, storageLocation string, functionLocation string, name string, args cloudfunctions.FunctionArgs) (*cloudfunctions.Function,error){
    bkt, err := bucket.NewBucket(ctx, storageLocation, name)
    if err != nil {
        return nil, err
    }

    archive, err := storage.NewBucketObject(ctx, name + "_archive", &storage.BucketObjectArgs{
		Bucket: bkt.Name,
        Source: pulumi.NewFileArchive(functionLocation),
	})
	if err != nil {
		return nil, err
	}
	function, err := cloudfunctions.NewFunction(ctx, name, &cloudfunctions.FunctionArgs {
        Description: args.Description,
		Runtime: args.Runtime,
		SourceArchiveBucket: bkt.Name,
		SourceArchiveObject: archive.Name,
        TriggerHttp: args.TriggerHttp,
		EntryPoint: args.EntryPoint,
        Region: args.Region,
    })
	if err != nil {
		return nil, err
	}
    _, err = cloudfunctions.NewFunctionIamMember(ctx, name + "_invoker", &cloudfunctions.FunctionIamMemberArgs{
		Project:       function.Project,
		Region:        function.Region,
		CloudFunction: function.Name,
		Role:          pulumi.String("roles/cloudfunctions.invoker"),
		Member:        pulumi.String("allUsers"),
	})
	if err != nil {
		return nil, err
	}
    return function, nil
}

