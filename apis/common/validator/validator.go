package validator

import (
	"basic/apis/common/rds"
	"basic/module/cache"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Validator 验证对象
type Validator struct {
	c         *fiber.Ctx
	validator *validator.Validate
}

// NewValidator 创建验证对象
func NewValidator(c *fiber.Ctx) *Validator {
	return &Validator{
		c:         c,
		validator: validate,
	}
}

// Validate 验证
func (_Validator *Validator) Validate(data interface{}) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	originParseUrl, _ := url.Parse(_Validator.c.Get("Origin"))
	lang := _Validator.c.Get("Accept-Language")
	adminId := rds.RedisFindDomainToAdminId(rdsConn, originParseUrl.Host)

	errs := _Validator.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("%v %s %v", err.Field(), rds.RedisFindTranslateField(rdsConn, adminId, lang, err.Tag()), err.Param()))
		}
	}
	return nil
}
