// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified destination of the specified delivery stream. Use this
// operation to change the destination type (for example, to replace the Amazon S3
// destination with Amazon Redshift) or change the parameters associated with a
// destination (for example, to change the bucket name of the Amazon S3
// destination). The update might not occur immediately. The target delivery stream
// remains active while the configurations are updated, so data writes to the
// delivery stream can continue during this process. The updated configurations are
// usually effective within a few minutes. Switching between Amazon OpenSearch
// Service and other services is not supported. For an Amazon OpenSearch Service
// destination, you can only update to another Amazon OpenSearch Service
// destination. If the destination type is the same, Kinesis Data Firehose merges
// the configuration parameters specified with the destination configuration that
// already exists on the delivery stream. If any of the parameters are not
// specified in the call, the existing values are retained. For example, in the
// Amazon S3 destination, if EncryptionConfiguration is not specified, then the
// existing EncryptionConfiguration is maintained on the destination. If the
// destination type is not the same, for example, changing the destination from
// Amazon S3 to Amazon Redshift, Kinesis Data Firehose does not merge any
// parameters. In this case, all parameters must be specified. Kinesis Data
// Firehose uses CurrentDeliveryStreamVersionId to avoid race conditions and
// conflicting merges. This is a required field, and the service updates the
// configuration only if the existing configuration has a version ID that matches.
// After the update is applied successfully, the version ID is updated, and can be
// retrieved using DescribeDeliveryStream . Use the new version ID to set
// CurrentDeliveryStreamVersionId in the next call.
func (c *Client) UpdateDestination(ctx context.Context, params *UpdateDestinationInput, optFns ...func(*Options)) (*UpdateDestinationOutput, error) {
	if params == nil {
		params = &UpdateDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDestination", params, optFns, c.addOperationUpdateDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDestinationInput struct {

	// Obtain this value from the VersionId result of DeliveryStreamDescription . This
	// value is required, and helps the service perform conditional operations. For
	// example, if there is an interleaving update and this value is null, then the
	// update destination fails. After the update is successful, the VersionId value
	// is updated. The service then performs a merge of the old configuration with the
	// new configuration.
	//
	// This member is required.
	CurrentDeliveryStreamVersionId *string

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// The ID of the destination.
	//
	// This member is required.
	DestinationId *string

	// Describes an update for a destination in the Serverless offering for Amazon
	// OpenSearch Service.
	AmazonOpenSearchServerlessDestinationUpdate *types.AmazonOpenSearchServerlessDestinationUpdate

	// Describes an update for a destination in Amazon OpenSearch Service.
	AmazonopensearchserviceDestinationUpdate *types.AmazonopensearchserviceDestinationUpdate

	// Describes an update for a destination in Amazon ES.
	ElasticsearchDestinationUpdate *types.ElasticsearchDestinationUpdate

	// Describes an update for a destination in Amazon S3.
	ExtendedS3DestinationUpdate *types.ExtendedS3DestinationUpdate

	// Describes an update to the specified HTTP endpoint destination.
	HttpEndpointDestinationUpdate *types.HttpEndpointDestinationUpdate

	// Describes an update for a destination in Amazon Redshift.
	RedshiftDestinationUpdate *types.RedshiftDestinationUpdate

	// [Deprecated] Describes an update for a destination in Amazon S3.
	//
	// Deprecated: This member has been deprecated.
	S3DestinationUpdate *types.S3DestinationUpdate

	// Describes an update for a destination in Splunk.
	SplunkDestinationUpdate *types.SplunkDestinationUpdate

	noSmithyDocumentSerde
}

type UpdateDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDestination{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateDestinationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDestination(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "UpdateDestination",
	}
}

type opUpdateDestinationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateDestinationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateDestinationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "firehose"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "firehose"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("firehose")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateDestinationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateDestinationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
