package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ErrorFlag string
type ErrorMSg string

const (
	ErrorFlagUncategorized ErrorFlag = "UNCATEGORIZED"
	ErrorFlagCanNotDelete  ErrorFlag = "CAN_NOT_DELETE"
	ErrorFlagCanNotUpdate  ErrorFlag = "CAN_NOT_UPDATE"
	ErrorFlagCanNotDisable ErrorFlag = "CAN_NOT_DISABLE"
	ErrorFlagCanNotEnable  ErrorFlag = "CAN_NOT_ENABLE"
	ErrorFlagCanNotCreate  ErrorFlag = "CAN_NOT_CREATE"
	ErrorFlagUpdateFail    ErrorFlag = "UPDATE_FAIL"
	ErrorFlagCreateFail    ErrorFlag = "CREATE_FAIL"
	ErrorFlagWarningUpdate ErrorFlag = "WARNING_UPDATE"
	NotFound               ErrorFlag = "RECORD_NOT_FOUND"
	ErrorFlagQueryFail     ErrorFlag = "FAIL_TO_QUERY"
	FieldHasAlreadyExisted string    = "%s has already existed"
	TotalSalaryMustBeEqual string    = "The total amount of salary components must equal the amount of gross salary"
	OnlySalaryBase         string    = "You can add only one component which is the salary base"
	RequireInput           string    = "Please enter %s"
	RecordNotFound         string    = "%s with id %s Not Found"
	CannotDeleteField      string    = "Cannot delete %s with id %s"
	FailToCreate           string    = "Fail to create %s"
	FailToUpdate           string    = "Fail to update %s"
	FailToDelete           string    = "Fail to delete %s"
	InvalidInput           string    = "Invalid input %s"
	CannotUpdateRecord     string    = "Cannot update %s with id %s"
	InvalidKeyword         string    = "%s 's keyword must contain only lowercase letters and underscore (_) only"
)

func WrapGQLError(ctx context.Context, message string, code int, errorFlag ErrorFlag) *gqlerror.Error {
	e := &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code":       code,
			"error_flag": errorFlag,
		},
	}
	if ctx != nil {
		e.Path = graphql.GetPath(ctx)
	}

	return e
}

func WrapGQLInternalError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Internal Server Error", http.StatusInternalServerError, ErrorFlagUncategorized)
}

func WrapGQLBadRequestError(ctx context.Context, format string, args ...interface{}) *gqlerror.Error {
	return WrapGQLError(ctx, fmt.Sprintf(format, args...), http.StatusBadRequest, ErrorFlagUncategorized)
}

func WrapGQLUnauthorizedError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Unauthorized Request", http.StatusUnauthorized, ErrorFlagUncategorized)
}

func WrapGQLNotFoundError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Not Found", http.StatusNotFound, NotFound)
}
