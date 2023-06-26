// Code generated by Kitex v0.6.1. DO NOT EDIT.

package bookservice

import (
	book "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return bookServiceServiceInfo
}

var bookServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BookService"
	handlerType := (*book.BookService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateBook": kitex.NewMethodInfo(createBookHandler, newBookServiceCreateBookArgs, newBookServiceCreateBookResult, false),
		"DeleteBook": kitex.NewMethodInfo(deleteBookHandler, newBookServiceDeleteBookArgs, newBookServiceDeleteBookResult, false),
		"UpdateBook": kitex.NewMethodInfo(updateBookHandler, newBookServiceUpdateBookArgs, newBookServiceUpdateBookResult, false),
		"QueryBook":  kitex.NewMethodInfo(queryBookHandler, newBookServiceQueryBookArgs, newBookServiceQueryBookResult, false),
		"MGetBook":   kitex.NewMethodInfo(mGetBookHandler, newBookServiceMGetBookArgs, newBookServiceMGetBookResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "book",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func createBookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*book.BookServiceCreateBookArgs)
	realResult := result.(*book.BookServiceCreateBookResult)
	success, err := handler.(book.BookService).CreateBook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBookServiceCreateBookArgs() interface{} {
	return book.NewBookServiceCreateBookArgs()
}

func newBookServiceCreateBookResult() interface{} {
	return book.NewBookServiceCreateBookResult()
}

func deleteBookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*book.BookServiceDeleteBookArgs)
	realResult := result.(*book.BookServiceDeleteBookResult)
	success, err := handler.(book.BookService).DeleteBook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBookServiceDeleteBookArgs() interface{} {
	return book.NewBookServiceDeleteBookArgs()
}

func newBookServiceDeleteBookResult() interface{} {
	return book.NewBookServiceDeleteBookResult()
}

func updateBookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*book.BookServiceUpdateBookArgs)
	realResult := result.(*book.BookServiceUpdateBookResult)
	success, err := handler.(book.BookService).UpdateBook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBookServiceUpdateBookArgs() interface{} {
	return book.NewBookServiceUpdateBookArgs()
}

func newBookServiceUpdateBookResult() interface{} {
	return book.NewBookServiceUpdateBookResult()
}

func queryBookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*book.BookServiceQueryBookArgs)
	realResult := result.(*book.BookServiceQueryBookResult)
	success, err := handler.(book.BookService).QueryBook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBookServiceQueryBookArgs() interface{} {
	return book.NewBookServiceQueryBookArgs()
}

func newBookServiceQueryBookResult() interface{} {
	return book.NewBookServiceQueryBookResult()
}

func mGetBookHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*book.BookServiceMGetBookArgs)
	realResult := result.(*book.BookServiceMGetBookResult)
	success, err := handler.(book.BookService).MGetBook(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBookServiceMGetBookArgs() interface{} {
	return book.NewBookServiceMGetBookArgs()
}

func newBookServiceMGetBookResult() interface{} {
	return book.NewBookServiceMGetBookResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateBook(ctx context.Context, req *book.CreateBookRequest) (r *book.CreateBookResponse, err error) {
	var _args book.BookServiceCreateBookArgs
	_args.Req = req
	var _result book.BookServiceCreateBookResult
	if err = p.c.Call(ctx, "CreateBook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (r *book.DeleteBookResponse, err error) {
	var _args book.BookServiceDeleteBookArgs
	_args.Req = req
	var _result book.BookServiceDeleteBookResult
	if err = p.c.Call(ctx, "DeleteBook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (r *book.UpdateBookResponse, err error) {
	var _args book.BookServiceUpdateBookArgs
	_args.Req = req
	var _result book.BookServiceUpdateBookResult
	if err = p.c.Call(ctx, "UpdateBook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryBook(ctx context.Context, req *book.QueryBookRequest) (r *book.QueryBookResponse, err error) {
	var _args book.BookServiceQueryBookArgs
	_args.Req = req
	var _result book.BookServiceQueryBookResult
	if err = p.c.Call(ctx, "QueryBook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetBook(ctx context.Context, req *book.MGetBookRequest) (r *book.MGetBookResponse, err error) {
	var _args book.BookServiceMGetBookArgs
	_args.Req = req
	var _result book.BookServiceMGetBookResult
	if err = p.c.Call(ctx, "MGetBook", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}