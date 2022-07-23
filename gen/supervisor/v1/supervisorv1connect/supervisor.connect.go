// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: supervisor/v1/supervisor.proto

package supervisorv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/submaline/operation-service/gen/supervisor/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SupervisorServiceName is the fully-qualified name of the SupervisorService service.
	SupervisorServiceName = "supervisor.v1.SupervisorService"
)

// SupervisorServiceClient is a client for the supervisor.v1.SupervisorService service.
type SupervisorServiceClient interface {
	CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error)
	CreateProfile(context.Context, *connect_go.Request[v1.CreateProfileRequest]) (*connect_go.Response[v1.CreateProfileResponse], error)
	RecordOperation(context.Context, *connect_go.Request[v1.RecordOperationRequest]) (*connect_go.Response[v1.RecordOperationResponse], error)
}

// NewSupervisorServiceClient constructs a client for the supervisor.v1.SupervisorService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSupervisorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SupervisorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &supervisorServiceClient{
		createAccount: connect_go.NewClient[v1.CreateAccountRequest, v1.CreateAccountResponse](
			httpClient,
			baseURL+"/supervisor.v1.SupervisorService/CreateAccount",
			opts...,
		),
		createProfile: connect_go.NewClient[v1.CreateProfileRequest, v1.CreateProfileResponse](
			httpClient,
			baseURL+"/supervisor.v1.SupervisorService/CreateProfile",
			opts...,
		),
		recordOperation: connect_go.NewClient[v1.RecordOperationRequest, v1.RecordOperationResponse](
			httpClient,
			baseURL+"/supervisor.v1.SupervisorService/RecordOperation",
			opts...,
		),
	}
}

// supervisorServiceClient implements SupervisorServiceClient.
type supervisorServiceClient struct {
	createAccount   *connect_go.Client[v1.CreateAccountRequest, v1.CreateAccountResponse]
	createProfile   *connect_go.Client[v1.CreateProfileRequest, v1.CreateProfileResponse]
	recordOperation *connect_go.Client[v1.RecordOperationRequest, v1.RecordOperationResponse]
}

// CreateAccount calls supervisor.v1.SupervisorService.CreateAccount.
func (c *supervisorServiceClient) CreateAccount(ctx context.Context, req *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error) {
	return c.createAccount.CallUnary(ctx, req)
}

// CreateProfile calls supervisor.v1.SupervisorService.CreateProfile.
func (c *supervisorServiceClient) CreateProfile(ctx context.Context, req *connect_go.Request[v1.CreateProfileRequest]) (*connect_go.Response[v1.CreateProfileResponse], error) {
	return c.createProfile.CallUnary(ctx, req)
}

// RecordOperation calls supervisor.v1.SupervisorService.RecordOperation.
func (c *supervisorServiceClient) RecordOperation(ctx context.Context, req *connect_go.Request[v1.RecordOperationRequest]) (*connect_go.Response[v1.RecordOperationResponse], error) {
	return c.recordOperation.CallUnary(ctx, req)
}

// SupervisorServiceHandler is an implementation of the supervisor.v1.SupervisorService service.
type SupervisorServiceHandler interface {
	CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error)
	CreateProfile(context.Context, *connect_go.Request[v1.CreateProfileRequest]) (*connect_go.Response[v1.CreateProfileResponse], error)
	RecordOperation(context.Context, *connect_go.Request[v1.RecordOperationRequest]) (*connect_go.Response[v1.RecordOperationResponse], error)
}

// NewSupervisorServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSupervisorServiceHandler(svc SupervisorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/supervisor.v1.SupervisorService/CreateAccount", connect_go.NewUnaryHandler(
		"/supervisor.v1.SupervisorService/CreateAccount",
		svc.CreateAccount,
		opts...,
	))
	mux.Handle("/supervisor.v1.SupervisorService/CreateProfile", connect_go.NewUnaryHandler(
		"/supervisor.v1.SupervisorService/CreateProfile",
		svc.CreateProfile,
		opts...,
	))
	mux.Handle("/supervisor.v1.SupervisorService/RecordOperation", connect_go.NewUnaryHandler(
		"/supervisor.v1.SupervisorService/RecordOperation",
		svc.RecordOperation,
		opts...,
	))
	return "/supervisor.v1.SupervisorService/", mux
}

// UnimplementedSupervisorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSupervisorServiceHandler struct{}

func (UnimplementedSupervisorServiceHandler) CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("supervisor.v1.SupervisorService.CreateAccount is not implemented"))
}

func (UnimplementedSupervisorServiceHandler) CreateProfile(context.Context, *connect_go.Request[v1.CreateProfileRequest]) (*connect_go.Response[v1.CreateProfileResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("supervisor.v1.SupervisorService.CreateProfile is not implemented"))
}

func (UnimplementedSupervisorServiceHandler) RecordOperation(context.Context, *connect_go.Request[v1.RecordOperationRequest]) (*connect_go.Response[v1.RecordOperationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("supervisor.v1.SupervisorService.RecordOperation is not implemented"))
}
