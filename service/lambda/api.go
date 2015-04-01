// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambda provides a client for AWS Lambda.
package lambda

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AddEventSourceRequest generates a request for the AddEventSource operation.
func (c *Lambda) AddEventSourceRequest(input *AddEventSourceInput) (req *aws.Request, output *EventSourceConfiguration) {
	if opAddEventSource == nil {
		opAddEventSource = &aws.Operation{
			Name:       "AddEventSource",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-11-13/event-source-mappings/",
		}
	}

	req = c.newRequest(opAddEventSource, input, output)
	output = &EventSourceConfiguration{}
	req.Data = output
	return
}

// Identifies a stream as an event source for an AWS Lambda function. It can
// be either an Amazon Kinesis stream or a Amazon DynamoDB stream. AWS Lambda
// invokes the specified function when records are posted to the stream.
//
// This is the pull model, where AWS Lambda invokes the function. For more
// information, go to AWS Lambda: How it Works (http://docs.aws.amazon.com/lambda/latest/dg/lambda-introduction.html)
// in the AWS Lambda Developer Guide.
//
// This association between an Amazon Kinesis stream and an AWS Lambda function
// is called the event source mapping. You provide the configuration information
// (for example, which stream to read from and which AWS Lambda function to
// invoke) for the event source mapping in the request body.
//
//  Each event source, such as a Kinesis stream, can only be associated with
// one AWS Lambda function. If you call AddEventSource for an event source that
// is already mapped to another AWS Lambda function, the existing mapping is
// updated to call the new function instead of the old one.
//
// This operation requires permission for the iam:PassRole action for the IAM
// role. It also requires permission for the lambda:AddEventSource action.
func (c *Lambda) AddEventSource(input *AddEventSourceInput) (output *EventSourceConfiguration, err error) {
	req, out := c.AddEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddEventSource *aws.Operation

// DeleteFunctionRequest generates a request for the DeleteFunction operation.
func (c *Lambda) DeleteFunctionRequest(input *DeleteFunctionInput) (req *aws.Request, output *DeleteFunctionOutput) {
	if opDeleteFunction == nil {
		opDeleteFunction = &aws.Operation{
			Name:       "DeleteFunction",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = c.newRequest(opDeleteFunction, input, output)
	output = &DeleteFunctionOutput{}
	req.Data = output
	return
}

// Deletes the specified Lambda function code and configuration.
//
// This operation requires permission for the lambda:DeleteFunction action.
func (c *Lambda) DeleteFunction(input *DeleteFunctionInput) (output *DeleteFunctionOutput, err error) {
	req, out := c.DeleteFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteFunction *aws.Operation

// GetEventSourceRequest generates a request for the GetEventSource operation.
func (c *Lambda) GetEventSourceRequest(input *GetEventSourceInput) (req *aws.Request, output *EventSourceConfiguration) {
	if opGetEventSource == nil {
		opGetEventSource = &aws.Operation{
			Name:       "GetEventSource",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/event-source-mappings/{UUID}",
		}
	}

	req = c.newRequest(opGetEventSource, input, output)
	output = &EventSourceConfiguration{}
	req.Data = output
	return
}

// Returns configuration information for the specified event source mapping
// (see AddEventSource).
//
// This operation requires permission for the lambda:GetEventSource action.
func (c *Lambda) GetEventSource(input *GetEventSourceInput) (output *EventSourceConfiguration, err error) {
	req, out := c.GetEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetEventSource *aws.Operation

// GetFunctionRequest generates a request for the GetFunction operation.
func (c *Lambda) GetFunctionRequest(input *GetFunctionInput) (req *aws.Request, output *GetFunctionOutput) {
	if opGetFunction == nil {
		opGetFunction = &aws.Operation{
			Name:       "GetFunction",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = c.newRequest(opGetFunction, input, output)
	output = &GetFunctionOutput{}
	req.Data = output
	return
}

// Returns the configuration information of the Lambda function and a presigned
// URL link to the .zip file you uploaded with UploadFunction so you can download
// the .zip file. Note that the URL is valid for up to 10 minutes. The configuration
// information is the same information you provided as parameters when uploading
// the function.
//
// This operation requires permission for the lambda:GetFunction action.
func (c *Lambda) GetFunction(input *GetFunctionInput) (output *GetFunctionOutput, err error) {
	req, out := c.GetFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetFunction *aws.Operation

// GetFunctionConfigurationRequest generates a request for the GetFunctionConfiguration operation.
func (c *Lambda) GetFunctionConfigurationRequest(input *GetFunctionConfigurationInput) (req *aws.Request, output *FunctionConfiguration) {
	if opGetFunctionConfiguration == nil {
		opGetFunctionConfiguration = &aws.Operation{
			Name:       "GetFunctionConfiguration",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/configuration",
		}
	}

	req = c.newRequest(opGetFunctionConfiguration, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

// Returns the configuration information of the Lambda function. This the same
// information you provided as parameters when uploading the function by using
// UploadFunction.
//
// This operation requires permission for the lambda:GetFunctionConfiguration
// operation.
func (c *Lambda) GetFunctionConfiguration(input *GetFunctionConfigurationInput) (output *FunctionConfiguration, err error) {
	req, out := c.GetFunctionConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetFunctionConfiguration *aws.Operation

// InvokeAsyncRequest generates a request for the InvokeAsync operation.
func (c *Lambda) InvokeAsyncRequest(input *InvokeAsyncInput) (req *aws.Request, output *InvokeAsyncOutput) {
	if opInvokeAsync == nil {
		opInvokeAsync = &aws.Operation{
			Name:       "InvokeAsync",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/invoke-async/",
		}
	}

	req = c.newRequest(opInvokeAsync, input, output)
	output = &InvokeAsyncOutput{}
	req.Data = output
	return
}

// Submits an invocation request to AWS Lambda. Upon receiving the request,
// Lambda executes the specified function asynchronously. To see the logs generated
// by the Lambda function execution, see the CloudWatch logs console.
//
// This operation requires permission for the lambda:InvokeAsync action.
func (c *Lambda) InvokeAsync(input *InvokeAsyncInput) (output *InvokeAsyncOutput, err error) {
	req, out := c.InvokeAsyncRequest(input)
	output = out
	err = req.Send()
	return
}

var opInvokeAsync *aws.Operation

// ListEventSourcesRequest generates a request for the ListEventSources operation.
func (c *Lambda) ListEventSourcesRequest(input *ListEventSourcesInput) (req *aws.Request, output *ListEventSourcesOutput) {
	if opListEventSources == nil {
		opListEventSources = &aws.Operation{
			Name:       "ListEventSources",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/event-source-mappings/",
		}
	}

	req = c.newRequest(opListEventSources, input, output)
	output = &ListEventSourcesOutput{}
	req.Data = output
	return
}

// Returns a list of event source mappings you created using the AddEventSource
// (see AddEventSource), where you identify a stream as event source. This list
// does not include Amazon S3 event sources.
//
// For each mapping, the API returns configuration information. You can optionally
// specify filters to retrieve specific event source mappings.
//
// This operation requires permission for the lambda:ListEventSources action.
func (c *Lambda) ListEventSources(input *ListEventSourcesInput) (output *ListEventSourcesOutput, err error) {
	req, out := c.ListEventSourcesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListEventSources *aws.Operation

// ListFunctionsRequest generates a request for the ListFunctions operation.
func (c *Lambda) ListFunctionsRequest(input *ListFunctionsInput) (req *aws.Request, output *ListFunctionsOutput) {
	if opListFunctions == nil {
		opListFunctions = &aws.Operation{
			Name:       "ListFunctions",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/",
		}
	}

	req = c.newRequest(opListFunctions, input, output)
	output = &ListFunctionsOutput{}
	req.Data = output
	return
}

// Returns a list of your Lambda functions. For each function, the response
// includes the function configuration information. You must use GetFunction
// to retrieve the code for your function.
//
// This operation requires permission for the lambda:ListFunctions action.
func (c *Lambda) ListFunctions(input *ListFunctionsInput) (output *ListFunctionsOutput, err error) {
	req, out := c.ListFunctionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListFunctions *aws.Operation

// RemoveEventSourceRequest generates a request for the RemoveEventSource operation.
func (c *Lambda) RemoveEventSourceRequest(input *RemoveEventSourceInput) (req *aws.Request, output *RemoveEventSourceOutput) {
	if opRemoveEventSource == nil {
		opRemoveEventSource = &aws.Operation{
			Name:       "RemoveEventSource",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2014-11-13/event-source-mappings/{UUID}",
		}
	}

	req = c.newRequest(opRemoveEventSource, input, output)
	output = &RemoveEventSourceOutput{}
	req.Data = output
	return
}

// Removes an event source mapping. This means AWS Lambda will no longer invoke
// the function for events in the associated source.
//
// This operation requires permission for the lambda:RemoveEventSource action.
func (c *Lambda) RemoveEventSource(input *RemoveEventSourceInput) (output *RemoveEventSourceOutput, err error) {
	req, out := c.RemoveEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRemoveEventSource *aws.Operation

// UpdateFunctionConfigurationRequest generates a request for the UpdateFunctionConfiguration operation.
func (c *Lambda) UpdateFunctionConfigurationRequest(input *UpdateFunctionConfigurationInput) (req *aws.Request, output *FunctionConfiguration) {
	if opUpdateFunctionConfiguration == nil {
		opUpdateFunctionConfiguration = &aws.Operation{
			Name:       "UpdateFunctionConfiguration",
			HTTPMethod: "PUT",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/configuration",
		}
	}

	req = c.newRequest(opUpdateFunctionConfiguration, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

// Updates the configuration parameters for the specified Lambda function by
// using the values provided in the request. You provide only the parameters
// you want to change. This operation must only be used on an existing Lambda
// function and cannot be used to update the function's code.
//
// This operation requires permission for the lambda:UpdateFunctionConfiguration
// action.
func (c *Lambda) UpdateFunctionConfiguration(input *UpdateFunctionConfigurationInput) (output *FunctionConfiguration, err error) {
	req, out := c.UpdateFunctionConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateFunctionConfiguration *aws.Operation

// UploadFunctionRequest generates a request for the UploadFunction operation.
func (c *Lambda) UploadFunctionRequest(input *UploadFunctionInput) (req *aws.Request, output *FunctionConfiguration) {
	if opUploadFunction == nil {
		opUploadFunction = &aws.Operation{
			Name:       "UploadFunction",
			HTTPMethod: "PUT",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = c.newRequest(opUploadFunction, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

// Creates a new Lambda function or updates an existing function. The function
// metadata is created from the request parameters, and the code for the function
// is provided by a .zip file in the request body. If the function name already
// exists, the existing Lambda function is updated with the new code and metadata.
//
// This operation requires permission for the lambda:UploadFunction action.
func (c *Lambda) UploadFunction(input *UploadFunctionInput) (output *FunctionConfiguration, err error) {
	req, out := c.UploadFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opUploadFunction *aws.Operation

type AddEventSourceInput struct {
	// The largest number of records that AWS Lambda will give to your function
	// in a single event. The default is 100 records.
	BatchSize *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream that is the event
	// source. Any record added to this stream causes AWS Lambda to invoke your
	// Lambda function. AWS Lambda POSTs the Amazon Kinesis event, containing records,
	// to your Lambda function as JSON.
	EventSource *string `type:"string" required:"true"`

	// The Lambda function to invoke when AWS Lambda detects an event on the stream.
	FunctionName *string `type:"string" required:"true"`

	// A map (key-value pairs) defining the configuration for AWS Lambda to use
	// when reading the event source. Currently, AWS Lambda supports only the InitialPositionInStream
	// key. The valid values are: "TRIM_HORIZON" and "LATEST". The default value
	// is "TRIM_HORIZON". For more information, go to ShardIteratorType (http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType)
	// in the Amazon Kinesis Service API Reference.
	Parameters *map[string]*string `type:"map"`

	// The ARN of the IAM role (invocation role) that AWS Lambda can assume to read
	// from the stream and invoke the function.
	Role *string `type:"string" required:"true"`

	metadataAddEventSourceInput `json:"-", xml:"-"`
}

type metadataAddEventSourceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteFunctionInput struct {
	// The Lambda function to delete.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	metadataDeleteFunctionInput `json:"-", xml:"-"`
}

type metadataDeleteFunctionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteFunctionOutput struct {
	metadataDeleteFunctionOutput `json:"-", xml:"-"`
}

type metadataDeleteFunctionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes mapping between an Amazon Kinesis stream and a Lambda function.
type EventSourceConfiguration struct {
	// The largest number of records that AWS Lambda will POST in the invocation
	// request to your function.
	BatchSize *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream that is the source
	// of events.
	EventSource *string `type:"string"`

	// The Lambda function to invoke when AWS Lambda detects an event on the stream.
	FunctionName *string `type:"string"`

	// Indicates whether the event source mapping is currently honored. Events are
	// only processes if IsActive is true.
	IsActive *bool `type:"boolean"`

	// The UTC time string indicating the last time the event mapping was updated.
	LastModified *string `type:"string"`

	// The map (key-value pairs) defining the configuration for AWS Lambda to use
	// when reading the event source.
	Parameters *map[string]*string `type:"map"`

	// The ARN of the IAM role (invocation role) that AWS Lambda can assume to read
	// from the stream and invoke the function.
	Role *string `type:"string"`

	// The description of the health of the event source mapping. Valid values are:
	// "PENDING", "OK", and "PROBLEM:message". Initially this staus is "PENDING".
	// When AWS Lambda begins processing events, it changes the status to "OK".
	Status *string `type:"string"`

	// The AWS Lambda assigned opaque identifier for the mapping.
	UUID *string `type:"string"`

	metadataEventSourceConfiguration `json:"-", xml:"-"`
}

type metadataEventSourceConfiguration struct {
	SDKShapeTraits bool `type:"structure"`
}

// The object for the Lambda function location.
type FunctionCodeLocation struct {
	// The presigned URL you can use to download the function's .zip file that you
	// previously uploaded. The URL is valid for up to 10 minutes.
	Location *string `type:"string"`

	// The repository from which you can download the function.
	RepositoryType *string `type:"string"`

	metadataFunctionCodeLocation `json:"-", xml:"-"`
}

type metadataFunctionCodeLocation struct {
	SDKShapeTraits bool `type:"structure"`
}

// A complex type that describes function metadata.
type FunctionConfiguration struct {
	// The size, in bytes, of the function .zip file you uploaded.
	CodeSize *int64 `type:"long"`

	// A Lambda-assigned unique identifier for the current function code and related
	// configuration.
	ConfigurationID *string `locationName:"ConfigurationId" type:"string"`

	// The user-provided description.
	Description *string `type:"string"`

	// The Amazon Resource Name (ARN) assigned to the function.
	FunctionARN *string `type:"string"`

	// The name of the function.
	FunctionName *string `type:"string"`

	// The function Lambda calls to begin executing your function.
	Handler *string `type:"string"`

	// The timestamp of the last time you updated the function.
	LastModified *string `type:"string"`

	// The memory size, in MB, you configured for the function. Must be a multiple
	// of 64 MB.
	MemorySize *int64 `type:"integer"`

	// The type of the Lambda function you uploaded.
	Mode *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role that Lambda assumes when it
	// executes your function to access any other Amazon Web Services (AWS) resources.
	Role *string `type:"string"`

	// The runtime environment for the Lambda function.
	Runtime *string `type:"string"`

	// The function execution time at which Lambda should terminate the function.
	// Because the execution time has cost implications, we recommend you set this
	// value based on your expected execution time. The default is 3 seconds.
	Timeout *int64 `type:"integer"`

	metadataFunctionConfiguration `json:"-", xml:"-"`
}

type metadataFunctionConfiguration struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetEventSourceInput struct {
	// The AWS Lambda assigned ID of the event source mapping.
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"`

	metadataGetEventSourceInput `json:"-", xml:"-"`
}

type metadataGetEventSourceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetFunctionConfigurationInput struct {
	// The name of the Lambda function for which you want to retrieve the configuration
	// information.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	metadataGetFunctionConfigurationInput `json:"-", xml:"-"`
}

type metadataGetFunctionConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetFunctionInput struct {
	// The Lambda function name.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	metadataGetFunctionInput `json:"-", xml:"-"`
}

type metadataGetFunctionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// This response contains the object for AWS Lambda function location (see API_FunctionCodeLocation
type GetFunctionOutput struct {
	// The object for the Lambda function location.
	Code *FunctionCodeLocation `type:"structure"`

	// A complex type that describes function metadata.
	Configuration *FunctionConfiguration `type:"structure"`

	metadataGetFunctionOutput `json:"-", xml:"-"`
}

type metadataGetFunctionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type InvokeAsyncInput struct {
	// The Lambda function name.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	// JSON that you want to provide to your Lambda function as input.
	InvokeArgs []byte `type:"blob" required:"true"`

	metadataInvokeAsyncInput `json:"-", xml:"-"`
}

type metadataInvokeAsyncInput struct {
	SDKShapeTraits bool `type:"structure" payload:"InvokeArgs"`
}

// Upon success, it returns empty response. Otherwise, throws an exception.
type InvokeAsyncOutput struct {
	// It will be 202 upon success.
	Status *int64 `location:"statusCode" type:"integer"`

	metadataInvokeAsyncOutput `json:"-", xml:"-"`
}

type metadataInvokeAsyncOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListEventSourcesInput struct {
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream.
	EventSourceARN *string `location:"querystring" locationName:"EventSource" type:"string"`

	// The name of the AWS Lambda function.
	FunctionName *string `location:"querystring" locationName:"FunctionName" type:"string"`

	// Optional string. An opaque pagination token returned from a previous ListEventSources
	// operation. If present, specifies to continue the list from where the returning
	// call left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Optional integer. Specifies the maximum number of event sources to return
	// in response. This value must be greater than 0.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	metadataListEventSourcesInput `json:"-", xml:"-"`
}

type metadataListEventSourcesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains a list of event sources (see API_EventSourceConfiguration)
type ListEventSourcesOutput struct {
	// An arrary of EventSourceConfiguration objects.
	EventSources []*EventSourceConfiguration `type:"list"`

	// A string, present if there are more event source mappings.
	NextMarker *string `type:"string"`

	metadataListEventSourcesOutput `json:"-", xml:"-"`
}

type metadataListEventSourcesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListFunctionsInput struct {
	// Optional string. An opaque pagination token returned from a previous ListFunctions
	// operation. If present, indicates where to continue the listing.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Optional integer. Specifies the maximum number of AWS Lambda functions to
	// return in response. This parameter value must be greater than 0.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	metadataListFunctionsInput `json:"-", xml:"-"`
}

type metadataListFunctionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Contains a list of AWS Lambda function configurations (see API_FunctionConfiguration.
type ListFunctionsOutput struct {
	// A list of Lambda functions.
	Functions []*FunctionConfiguration `type:"list"`

	// A string, present if there are more functions.
	NextMarker *string `type:"string"`

	metadataListFunctionsOutput `json:"-", xml:"-"`
}

type metadataListFunctionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RemoveEventSourceInput struct {
	// The event source mapping ID.
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"`

	metadataRemoveEventSourceInput `json:"-", xml:"-"`
}

type metadataRemoveEventSourceInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RemoveEventSourceOutput struct {
	metadataRemoveEventSourceOutput `json:"-", xml:"-"`
}

type metadataRemoveEventSourceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateFunctionConfigurationInput struct {
	// A short user-defined function description. Lambda does not use this value.
	// Assign a meaningful description as you see fit.
	Description *string `location:"querystring" locationName:"Description" type:"string"`

	// The name of the Lambda function.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	// The function that Lambda calls to begin executing your function. For Node.js,
	// it is the module-name.export value in your function.
	Handler *string `location:"querystring" locationName:"Handler" type:"string"`

	// The amount of memory, in MB, your Lambda function is given. Lambda uses this
	// memory size to infer the amount of CPU allocated to your function. Your function
	// use-case determines your CPU and memory requirements. For example, a database
	// operation might need less memory compared to an image processing function.
	// The default value is 128 MB. The value must be a multiple of 64 MB.
	MemorySize *int64 `location:"querystring" locationName:"MemorySize" type:"integer"`

	// The Amazon Resource Name (ARN) of the IAM role that Lambda will assume when
	// it executes your function.
	Role *string `location:"querystring" locationName:"Role" type:"string"`

	// The function execution time at which Lambda should terminate the function.
	// Because the execution time has cost implications, we recommend you set this
	// value based on your expected execution time. The default is 3 seconds.
	Timeout *int64 `location:"querystring" locationName:"Timeout" type:"integer"`

	metadataUpdateFunctionConfigurationInput `json:"-", xml:"-"`
}

type metadataUpdateFunctionConfigurationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UploadFunctionInput struct {
	// A short, user-defined function description. Lambda does not use this value.
	// Assign a meaningful description as you see fit.
	Description *string `location:"querystring" locationName:"Description" type:"string"`

	// The name you want to assign to the function you are uploading. The function
	// names appear in the console and are returned in the ListFunctions API. Function
	// names are used to specify functions to other AWS Lambda APIs, such as InvokeAsync.
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"`

	// A .zip file containing your packaged source code. For more information about
	// creating a .zip file, go to AWS LambdaL How it Works (http://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events.html)
	// in the AWS Lambda Developer Guide.
	FunctionZip []byte `type:"blob" required:"true"`

	// The function that Lambda calls to begin execution. For Node.js, it is the
	// module-name.export value in your function.
	Handler *string `location:"querystring" locationName:"Handler" type:"string" required:"true"`

	// The amount of memory, in MB, your Lambda function is given. Lambda uses this
	// memory size to infer the amount of CPU allocated to your function. Your function
	// use-case determines your CPU and memory requirements. For example, database
	// operation might need less memory compared to image processing function. The
	// default value is 128 MB. The value must be a multiple of 64 MB.
	MemorySize *int64 `location:"querystring" locationName:"MemorySize" type:"integer"`

	// How the Lambda function will be invoked. Lambda supports only the "event"
	// mode.
	Mode *string `location:"querystring" locationName:"Mode" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role that Lambda assumes when it
	// executes your function to access any other Amazon Web Services (AWS) resources.
	Role *string `location:"querystring" locationName:"Role" type:"string" required:"true"`

	// The runtime environment for the Lambda function you are uploading. Currently,
	// Lambda supports only "nodejs" as the runtime.
	Runtime *string `location:"querystring" locationName:"Runtime" type:"string" required:"true"`

	// The function execution time at which Lambda should terminate the function.
	// Because the execution time has cost implications, we recommend you set this
	// value based on your expected execution time. The default is 3 seconds.
	Timeout *int64 `location:"querystring" locationName:"Timeout" type:"integer"`

	metadataUploadFunctionInput `json:"-", xml:"-"`
}

type metadataUploadFunctionInput struct {
	SDKShapeTraits bool `type:"structure" payload:"FunctionZip"`
}