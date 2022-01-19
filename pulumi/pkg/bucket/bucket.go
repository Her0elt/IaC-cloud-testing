package bucket


import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewBucket(ctx *pulumi.Context, location string, functionName string) (*storage.Bucket, error) {
	bucket, err := storage.NewBucket(ctx, functionName +"_bucket", &storage.BucketArgs{
		Location: pulumi.String(location),
	})
	if err != nil {
	    return nil, err
    }
    return bucket, nil
}
