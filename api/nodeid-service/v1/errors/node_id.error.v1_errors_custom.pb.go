// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorv1

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	strconv "strconv"
)

var ERROR_http_code = map[string]int{

	"UNKNOWN":                       500,
	"S102_NO_AVAILABLE_ID":          400,
	"S102_RECORD_NOT_FOUNT":         400,
	"S102_RECORD_ALREADY_EXIST":     400,
	"S102_HAS_BEEN_USED":            400,
	"S102_NODE_ID_RENEWAL_FAILED":   400,
	"S102_NODE_ID_INCORRECT":        400,
	"S102_ACCESS_TOKEN_INCORRECT":   400,
	"S102_NODE_ID_STATUS_INCORRECT": 400,
}

func (x ERROR) HTTPCode() int {
	if v, ok := ERROR_http_code[x.String()]; ok {
		return v
	}
	return 500
}

// 未知错误
func DefaultErrorUnknown() *errors.Error {
	e := errors.New(500, ERROR_UNKNOWN.String(), "未知错误")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_UNKNOWN.Number()))}
	return e
}

// 没有可用的节点ID
func DefaultErrorS102NoAvailableId() *errors.Error {
	e := errors.New(400, ERROR_S102_NO_AVAILABLE_ID.String(), "没有可用的节点ID")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_NO_AVAILABLE_ID.Number()))}
	return e
}

// 节点ID不存在
func DefaultErrorS102RecordNotFount() *errors.Error {
	e := errors.New(400, ERROR_S102_RECORD_NOT_FOUNT.String(), "节点ID不存在")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_RECORD_NOT_FOUNT.Number()))}
	return e
}

// 节点ID已被使用
func DefaultErrorS102RecordAlreadyExist() *errors.Error {
	e := errors.New(400, ERROR_S102_RECORD_ALREADY_EXIST.String(), "节点ID已被使用")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_RECORD_ALREADY_EXIST.Number()))}
	return e
}

// 节点ID已被使用
func DefaultErrorS102HasBeenUsed() *errors.Error {
	e := errors.New(400, ERROR_S102_HAS_BEEN_USED.String(), "节点ID已被使用")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_HAS_BEEN_USED.Number()))}
	return e
}

// 节点ID续订失败
func DefaultErrorS102NodeIdRenewalFailed() *errors.Error {
	e := errors.New(400, ERROR_S102_NODE_ID_RENEWAL_FAILED.String(), "节点ID续订失败")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_NODE_ID_RENEWAL_FAILED.Number()))}
	return e
}

// 节点ID信息不正确
func DefaultErrorS102NodeIdIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S102_NODE_ID_INCORRECT.String(), "节点ID信息不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_NODE_ID_INCORRECT.Number()))}
	return e
}

// 访问令牌不正确
func DefaultErrorS102AccessTokenIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S102_ACCESS_TOKEN_INCORRECT.String(), "访问令牌不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_ACCESS_TOKEN_INCORRECT.Number()))}
	return e
}

// 节点状态不正确
func DefaultErrorS102NodeIdStatusIncorrect() *errors.Error {
	e := errors.New(400, ERROR_S102_NODE_ID_STATUS_INCORRECT.String(), "节点状态不正确")
	e.Metadata = map[string]string{"reason": strconv.Itoa(int(ERROR_S102_NODE_ID_STATUS_INCORRECT.Number()))}
	return e
}
