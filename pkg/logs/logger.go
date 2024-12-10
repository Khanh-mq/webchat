package logs

//
//import (
//	"errors"
//	"fmt"
//	"net/http"
//	"strings"
//)
//
//type AppError struct {
//	StatusCode int    `json:"status_code"`
//	RootErr    error  `json:"_"`
//	Message    string `json:"message"`
//	Log        string `json:"log"`
//	Key        string `json:"key"`
//}
//
//func NewAppError(status int, root error, mgs, log, key string) *AppError {
//	return &AppError{
//		StatusCode: status,
//		RootErr:    root,
//		Message:    mgs,
//		Log:        log,
//		Key:        key,
//	}
//}
//
//func NewErrorResponse(root error, mgs, log, key string) *AppError {
//	return &AppError{
//		StatusCode: http.StatusBadRequest,
//		RootErr:    root,
//		Message:    mgs,
//		Log:        log,
//		Key:        key,
//	}
//}
//
////	func (e *AppError) Error() string {
////		return e.RootErr.Error()
////	}
////
////	func (e *AppError) RootError() error {
////		if err, ok := e.RootErr.(*AppError); ok {
////			//de quy
////			return err.RootError()
////		}
////		return e.RootErr
////	}
//func NewCustomError(root error, msg string, key string) *AppError {
//	if root != nil {
//		return NewErrorResponse(root, msg, root.Error(), key)
//
//	}
//	return NewErrorResponse(errors.New(msg), msg, msg, key)
//}
//func ErrDB(err error) *AppError {
//	return NewAppError(http.StatusInternalServerError, nil, "something went wrong with DB ", err.Error(), "DB_ERROR")
//}
//func ErrInvalidRequest(err error) *AppError {
//	return NewAppError(http.StatusInternalServerError, err, "something went wrong with invalid request ", err.Error(), "INVALID_REQUEST")
//
//}
//func ErrInternal(err error) *AppError {
//	return NewAppError(http.StatusInternalServerError, err, "something went wrong with internal ", err.Error(), "INTERNAL_ERROR")
//
//}
//func ErrorCannotListEntity(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("cannot list %s", strings.ToLower(entity)),
//		fmt.Sprintf("something went wrong with cannot list %s", entity),
//	)
//}
//func ErrorCannotDeletedEntity(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("cannot deleted %s", strings.ToLower(entity)),
//		fmt.Sprintf("something went wrong with cannot delete %s", entity),
//	)
//
//}
//func ErrCannotCreateEntity(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
//		fmt.Sprintf("ErrorCannotCreate%s", entity),
//	)
//}
//func ErrCannotUpdateEntity(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
//		fmt.Sprintf("ErrorUpdateCreate%s", entity),
//	)
//}
//func ErrCannotGetEntity(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("cannot get %s", strings.ToLower(entity)),
//		fmt.Sprintf("Error cannot get %s", entity),
//	)
//}
//func ErrEntityNotFound(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("%s not found", strings.ToLower(entity)),
//		fmt.Sprintf("Error%sNotFound", entity),
//	)
//}
//func ErrEntityExisted(entity string, err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("%s Existed", strings.ToLower(entity)),
//		fmt.Sprintf("Error%sExisted", entity),
//	)
//}
//func ErrNoPermission(err error) *AppError {
//	return NewCustomError(
//		err,
//		fmt.Sprintf("you have no permission"),
//		fmt.Sprintf("ErrorNoPermission"),
//	)
//}
//
//var RecordNotFound = errors.New("record not found")
