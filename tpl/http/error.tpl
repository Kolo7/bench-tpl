package http

import (
	"encoding/json"
	"fmt"

	"git.imgo.tv/ft/go-ceres/pkg/net/http"
	"git.imgo.tv/ft/go-lib2/ecode"
	v10 "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
    "{{.fqdn}}/pkg/validator"
)

// 错误信息翻译
func JSONRequestError(r *http.Gin, err error) {

	tagErr := &v10.ValidationErrors{}
	if errors.As(err, &tagErr) {
		r.JSON(nil, ecode.Transient(ecode.Code(ecode.RequestErr.Code()), validator.JoinMap(tagErr.Translate(validator.Trans))))
		return
	}
	if errors.As(err, tagErr) {
		r.JSON(nil, ecode.Transient(ecode.Code(ecode.RequestErr.Code()), validator.JoinMap(tagErr.Translate(validator.Trans))))
		return
	}
	tagErr2 := &json.UnmarshalTypeError{}
	if errors.As(err, &tagErr2) {
		r.JSON(nil, ecode.Transient(ecode.Code(ecode.RequestErr.Code()), fmt.Sprintf("字段[%s]类型错误,请使用[%s]", tagErr2.Field, tagErr2.Type)))
		return
	}

	r.JSON(nil, ecode.RequestErr)
}
