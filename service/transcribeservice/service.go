// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribeservice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// TranscribeService provides the API operation methods for making requests to
// Amazon Transcribe Service. See this package's package overview docs
// for details on the service.
//
// TranscribeService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type TranscribeService struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*TranscribeService)

// Used for custom request initialization logic
var initRequest func(*TranscribeService, *aws.Request)

// Service information constants
const (
	ServiceName = "transcribe" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName  // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the TranscribeService client with a config.
//
// Example:
//     // Create a TranscribeService client from just a config.
//     svc := transcribeservice.New(myConfig)
func New(config aws.Config) *TranscribeService {
	var signingName string
	signingName = "transcribe"
	signingRegion := config.Region

	svc := &TranscribeService{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-10-26",
				JSONVersion:   "1.1",
				TargetPrefix:  "Transcribe",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a TranscribeService operation and runs any
// custom request initialization.
func (c *TranscribeService) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
