package errorx

import (
	"errors"
	"fmt"
)

type ExpectError struct {
	Code uint32
	Msg  string
	err  error
}

func (e *ExpectError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%d:%s:%s", e.Code, e.Msg, e.err.Error())
	}
	return fmt.Sprintf("%d:%s", e.Code, e.Msg)
}

func (e *ExpectError) Unwrap() error {
	return e.err
}

func (e *ExpectError) DeepestExpectError() *ExpectError {
	if e.err != nil {
		var expect *ExpectError
		if errors.As(e.err, &expect) {
			return expect.DeepestExpectError()
		}
	}
	return e
}

// NewExpertError 构造Error
func NewExpectError(code uint32, msg string) *ExpectError {
	return &ExpectError{
		Code: code,
		Msg:  msg,
	}
}

// Wrap 拷贝并包装错误
func (e *ExpectError) WrapErr(err error) *ExpectError {
	temp := *e
	dup := &temp
	dup.err = err
	return dup
}

// Wrap 拷贝并包装错误信息
func (e *ExpectError) Wrap(err string) *ExpectError {
	temp := *e
	dup := &temp
	dup.err = errors.New(err)
	return dup
}
