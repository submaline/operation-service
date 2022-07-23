// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/submaline/operation-service/gen/user/v1"
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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	// アカウント
	GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error)
	UpdateAccount(context.Context, *connect_go.Request[v1.UpdateAccountRequest]) (*connect_go.Response[v1.UpdateAccountResponse], error)
	// プロフィール
	GetProfile(context.Context, *connect_go.Request[v1.GetProfileRequest]) (*connect_go.Response[v1.GetProfileResponse], error)
	UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.UpdateProfileResponse], error)
	// 友達
	AddFriend(context.Context, *connect_go.Request[v1.AddFriendRequest]) (*connect_go.Response[v1.AddFriendResponse], error)
	GetFriends(context.Context, *connect_go.Request[v1.GetFriendsRequest]) (*connect_go.Response[v1.GetFriendsResponse], error)
	// ブロック
	BlockUser(context.Context, *connect_go.Request[v1.BlockUserRequest]) (*connect_go.Response[v1.BlockUserResponse], error)
	GetBlockingUsers(context.Context, *connect_go.Request[v1.GetBlockingUsersRequest]) (*connect_go.Response[v1.GetBlockingUsersResponse], error)
	UnBlockUser(context.Context, *connect_go.Request[v1.UnBlockUserRequest]) (*connect_go.Response[v1.UnBlockUserResponse], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getAccount: connect_go.NewClient[v1.GetAccountRequest, v1.GetAccountResponse](
			httpClient,
			baseURL+"/user.v1.UserService/GetAccount",
			opts...,
		),
		updateAccount: connect_go.NewClient[v1.UpdateAccountRequest, v1.UpdateAccountResponse](
			httpClient,
			baseURL+"/user.v1.UserService/UpdateAccount",
			opts...,
		),
		getProfile: connect_go.NewClient[v1.GetProfileRequest, v1.GetProfileResponse](
			httpClient,
			baseURL+"/user.v1.UserService/GetProfile",
			opts...,
		),
		updateProfile: connect_go.NewClient[v1.UpdateProfileRequest, v1.UpdateProfileResponse](
			httpClient,
			baseURL+"/user.v1.UserService/UpdateProfile",
			opts...,
		),
		addFriend: connect_go.NewClient[v1.AddFriendRequest, v1.AddFriendResponse](
			httpClient,
			baseURL+"/user.v1.UserService/AddFriend",
			opts...,
		),
		getFriends: connect_go.NewClient[v1.GetFriendsRequest, v1.GetFriendsResponse](
			httpClient,
			baseURL+"/user.v1.UserService/GetFriends",
			opts...,
		),
		blockUser: connect_go.NewClient[v1.BlockUserRequest, v1.BlockUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/BlockUser",
			opts...,
		),
		getBlockingUsers: connect_go.NewClient[v1.GetBlockingUsersRequest, v1.GetBlockingUsersResponse](
			httpClient,
			baseURL+"/user.v1.UserService/GetBlockingUsers",
			opts...,
		),
		unBlockUser: connect_go.NewClient[v1.UnBlockUserRequest, v1.UnBlockUserResponse](
			httpClient,
			baseURL+"/user.v1.UserService/UnBlockUser",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getAccount       *connect_go.Client[v1.GetAccountRequest, v1.GetAccountResponse]
	updateAccount    *connect_go.Client[v1.UpdateAccountRequest, v1.UpdateAccountResponse]
	getProfile       *connect_go.Client[v1.GetProfileRequest, v1.GetProfileResponse]
	updateProfile    *connect_go.Client[v1.UpdateProfileRequest, v1.UpdateProfileResponse]
	addFriend        *connect_go.Client[v1.AddFriendRequest, v1.AddFriendResponse]
	getFriends       *connect_go.Client[v1.GetFriendsRequest, v1.GetFriendsResponse]
	blockUser        *connect_go.Client[v1.BlockUserRequest, v1.BlockUserResponse]
	getBlockingUsers *connect_go.Client[v1.GetBlockingUsersRequest, v1.GetBlockingUsersResponse]
	unBlockUser      *connect_go.Client[v1.UnBlockUserRequest, v1.UnBlockUserResponse]
}

// GetAccount calls user.v1.UserService.GetAccount.
func (c *userServiceClient) GetAccount(ctx context.Context, req *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error) {
	return c.getAccount.CallUnary(ctx, req)
}

// UpdateAccount calls user.v1.UserService.UpdateAccount.
func (c *userServiceClient) UpdateAccount(ctx context.Context, req *connect_go.Request[v1.UpdateAccountRequest]) (*connect_go.Response[v1.UpdateAccountResponse], error) {
	return c.updateAccount.CallUnary(ctx, req)
}

// GetProfile calls user.v1.UserService.GetProfile.
func (c *userServiceClient) GetProfile(ctx context.Context, req *connect_go.Request[v1.GetProfileRequest]) (*connect_go.Response[v1.GetProfileResponse], error) {
	return c.getProfile.CallUnary(ctx, req)
}

// UpdateProfile calls user.v1.UserService.UpdateProfile.
func (c *userServiceClient) UpdateProfile(ctx context.Context, req *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.UpdateProfileResponse], error) {
	return c.updateProfile.CallUnary(ctx, req)
}

// AddFriend calls user.v1.UserService.AddFriend.
func (c *userServiceClient) AddFriend(ctx context.Context, req *connect_go.Request[v1.AddFriendRequest]) (*connect_go.Response[v1.AddFriendResponse], error) {
	return c.addFriend.CallUnary(ctx, req)
}

// GetFriends calls user.v1.UserService.GetFriends.
func (c *userServiceClient) GetFriends(ctx context.Context, req *connect_go.Request[v1.GetFriendsRequest]) (*connect_go.Response[v1.GetFriendsResponse], error) {
	return c.getFriends.CallUnary(ctx, req)
}

// BlockUser calls user.v1.UserService.BlockUser.
func (c *userServiceClient) BlockUser(ctx context.Context, req *connect_go.Request[v1.BlockUserRequest]) (*connect_go.Response[v1.BlockUserResponse], error) {
	return c.blockUser.CallUnary(ctx, req)
}

// GetBlockingUsers calls user.v1.UserService.GetBlockingUsers.
func (c *userServiceClient) GetBlockingUsers(ctx context.Context, req *connect_go.Request[v1.GetBlockingUsersRequest]) (*connect_go.Response[v1.GetBlockingUsersResponse], error) {
	return c.getBlockingUsers.CallUnary(ctx, req)
}

// UnBlockUser calls user.v1.UserService.UnBlockUser.
func (c *userServiceClient) UnBlockUser(ctx context.Context, req *connect_go.Request[v1.UnBlockUserRequest]) (*connect_go.Response[v1.UnBlockUserResponse], error) {
	return c.unBlockUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	// アカウント
	GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error)
	UpdateAccount(context.Context, *connect_go.Request[v1.UpdateAccountRequest]) (*connect_go.Response[v1.UpdateAccountResponse], error)
	// プロフィール
	GetProfile(context.Context, *connect_go.Request[v1.GetProfileRequest]) (*connect_go.Response[v1.GetProfileResponse], error)
	UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.UpdateProfileResponse], error)
	// 友達
	AddFriend(context.Context, *connect_go.Request[v1.AddFriendRequest]) (*connect_go.Response[v1.AddFriendResponse], error)
	GetFriends(context.Context, *connect_go.Request[v1.GetFriendsRequest]) (*connect_go.Response[v1.GetFriendsResponse], error)
	// ブロック
	BlockUser(context.Context, *connect_go.Request[v1.BlockUserRequest]) (*connect_go.Response[v1.BlockUserResponse], error)
	GetBlockingUsers(context.Context, *connect_go.Request[v1.GetBlockingUsersRequest]) (*connect_go.Response[v1.GetBlockingUsersResponse], error)
	UnBlockUser(context.Context, *connect_go.Request[v1.UnBlockUserRequest]) (*connect_go.Response[v1.UnBlockUserResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/user.v1.UserService/GetAccount", connect_go.NewUnaryHandler(
		"/user.v1.UserService/GetAccount",
		svc.GetAccount,
		opts...,
	))
	mux.Handle("/user.v1.UserService/UpdateAccount", connect_go.NewUnaryHandler(
		"/user.v1.UserService/UpdateAccount",
		svc.UpdateAccount,
		opts...,
	))
	mux.Handle("/user.v1.UserService/GetProfile", connect_go.NewUnaryHandler(
		"/user.v1.UserService/GetProfile",
		svc.GetProfile,
		opts...,
	))
	mux.Handle("/user.v1.UserService/UpdateProfile", connect_go.NewUnaryHandler(
		"/user.v1.UserService/UpdateProfile",
		svc.UpdateProfile,
		opts...,
	))
	mux.Handle("/user.v1.UserService/AddFriend", connect_go.NewUnaryHandler(
		"/user.v1.UserService/AddFriend",
		svc.AddFriend,
		opts...,
	))
	mux.Handle("/user.v1.UserService/GetFriends", connect_go.NewUnaryHandler(
		"/user.v1.UserService/GetFriends",
		svc.GetFriends,
		opts...,
	))
	mux.Handle("/user.v1.UserService/BlockUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/BlockUser",
		svc.BlockUser,
		opts...,
	))
	mux.Handle("/user.v1.UserService/GetBlockingUsers", connect_go.NewUnaryHandler(
		"/user.v1.UserService/GetBlockingUsers",
		svc.GetBlockingUsers,
		opts...,
	))
	mux.Handle("/user.v1.UserService/UnBlockUser", connect_go.NewUnaryHandler(
		"/user.v1.UserService/UnBlockUser",
		svc.UnBlockUser,
		opts...,
	))
	return "/user.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.GetAccount is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateAccount(context.Context, *connect_go.Request[v1.UpdateAccountRequest]) (*connect_go.Response[v1.UpdateAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.UpdateAccount is not implemented"))
}

func (UnimplementedUserServiceHandler) GetProfile(context.Context, *connect_go.Request[v1.GetProfileRequest]) (*connect_go.Response[v1.GetProfileResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.GetProfile is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.UpdateProfileResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.UpdateProfile is not implemented"))
}

func (UnimplementedUserServiceHandler) AddFriend(context.Context, *connect_go.Request[v1.AddFriendRequest]) (*connect_go.Response[v1.AddFriendResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.AddFriend is not implemented"))
}

func (UnimplementedUserServiceHandler) GetFriends(context.Context, *connect_go.Request[v1.GetFriendsRequest]) (*connect_go.Response[v1.GetFriendsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.GetFriends is not implemented"))
}

func (UnimplementedUserServiceHandler) BlockUser(context.Context, *connect_go.Request[v1.BlockUserRequest]) (*connect_go.Response[v1.BlockUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.BlockUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetBlockingUsers(context.Context, *connect_go.Request[v1.GetBlockingUsersRequest]) (*connect_go.Response[v1.GetBlockingUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.GetBlockingUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) UnBlockUser(context.Context, *connect_go.Request[v1.UnBlockUserRequest]) (*connect_go.Response[v1.UnBlockUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("user.v1.UserService.UnBlockUser is not implemented"))
}