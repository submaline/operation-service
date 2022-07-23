package operation_service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	operationv1 "github.com/submaline/operation-service/gen/operation/v1"
)

type OperationService struct{}

func (os *OperationService) FetchOperations(
	_ context.Context, request *connect.Request[operationv1.FetchOperationsRequest],
	_ *connect.ServerStream[operationv1.FetchOperationsResponse]) error {
	return connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
