// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: talk/v1/talk.proto

package talkv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/submaline/operation-service/gen/talk/v1"
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
	// TalkServiceName is the fully-qualified name of the TalkService service.
	TalkServiceName = "talk.v1.TalkService"
)

// TalkServiceClient is a client for the talk.v1.TalkService service.
type TalkServiceClient interface {
	SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error)
	SendReadReceipt(context.Context, *connect_go.Request[v1.SendReadReceiptRequest]) (*connect_go.Response[v1.SendReadReceiptResponse], error)
}

// NewTalkServiceClient constructs a client for the talk.v1.TalkService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTalkServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TalkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &talkServiceClient{
		sendMessage: connect_go.NewClient[v1.SendMessageRequest, v1.SendMessageResponse](
			httpClient,
			baseURL+"/talk.v1.TalkService/SendMessage",
			opts...,
		),
		sendReadReceipt: connect_go.NewClient[v1.SendReadReceiptRequest, v1.SendReadReceiptResponse](
			httpClient,
			baseURL+"/talk.v1.TalkService/SendReadReceipt",
			opts...,
		),
	}
}

// talkServiceClient implements TalkServiceClient.
type talkServiceClient struct {
	sendMessage     *connect_go.Client[v1.SendMessageRequest, v1.SendMessageResponse]
	sendReadReceipt *connect_go.Client[v1.SendReadReceiptRequest, v1.SendReadReceiptResponse]
}

// SendMessage calls talk.v1.TalkService.SendMessage.
func (c *talkServiceClient) SendMessage(ctx context.Context, req *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error) {
	return c.sendMessage.CallUnary(ctx, req)
}

// SendReadReceipt calls talk.v1.TalkService.SendReadReceipt.
func (c *talkServiceClient) SendReadReceipt(ctx context.Context, req *connect_go.Request[v1.SendReadReceiptRequest]) (*connect_go.Response[v1.SendReadReceiptResponse], error) {
	return c.sendReadReceipt.CallUnary(ctx, req)
}

// TalkServiceHandler is an implementation of the talk.v1.TalkService service.
type TalkServiceHandler interface {
	SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error)
	SendReadReceipt(context.Context, *connect_go.Request[v1.SendReadReceiptRequest]) (*connect_go.Response[v1.SendReadReceiptResponse], error)
}

// NewTalkServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTalkServiceHandler(svc TalkServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/talk.v1.TalkService/SendMessage", connect_go.NewUnaryHandler(
		"/talk.v1.TalkService/SendMessage",
		svc.SendMessage,
		opts...,
	))
	mux.Handle("/talk.v1.TalkService/SendReadReceipt", connect_go.NewUnaryHandler(
		"/talk.v1.TalkService/SendReadReceipt",
		svc.SendReadReceipt,
		opts...,
	))
	return "/talk.v1.TalkService/", mux
}

// UnimplementedTalkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTalkServiceHandler struct{}

func (UnimplementedTalkServiceHandler) SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("talk.v1.TalkService.SendMessage is not implemented"))
}

func (UnimplementedTalkServiceHandler) SendReadReceipt(context.Context, *connect_go.Request[v1.SendReadReceiptRequest]) (*connect_go.Response[v1.SendReadReceiptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("talk.v1.TalkService.SendReadReceipt is not implemented"))
}
