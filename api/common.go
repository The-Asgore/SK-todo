package api

import (
	"SK-todo/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("%v", e.Field())
			tag := fmt.Sprintf("Valid.%sv", e.Tag())
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON type error",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: 40001,
		Msg:    "Invalid parameters",
		Error:  fmt.Sprint(err),
	}
}
