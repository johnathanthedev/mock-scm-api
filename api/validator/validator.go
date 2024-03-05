// Custom validator package provides functionality for validating struct fields based on struct tags.

package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func Init() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// "Validate" validates the fields of a struct based on struct tags.
// It returns an error if any field fails validation.
func (cv *CustomValidator) Validate(i interface{}) error {
	// Get the value of the interface
	value := reflect.ValueOf(i)

	// If the provided interface is a pointer, get the value it points to
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Get the type of the struct
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)

		if tag, ok := field.Tag.Lookup("validate"); ok {
			rules := strings.Split(tag, ",")

			for _, rule := range rules {
				parts := strings.Split(rule, "=")
				ruleName := parts[0]
				var ruleValue string
				if len(parts) > 1 {
					ruleValue = parts[1]
				}

				switch ruleName {
				case "required":
					fieldValue := value.Field(i)

					if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
						return fmt.Errorf("'%s' is required", field.Name)
					}
				case "min":
					fieldValue := value.Field(i)

					minLen, err := strconv.Atoi(ruleValue)
					if err != nil {
						return fmt.Errorf("invalid min value for field '%s'", field.Name)
					}

					if len(fieldValue.String()) < minLen {
						return fmt.Errorf("'%s' must be at least %s characters long", field.Name, ruleValue)
					}
				case "max":
					fieldValue := value.Field(i)

					maxLen, err := strconv.Atoi(ruleValue)
					if err != nil {
						return fmt.Errorf("invalid max value for field '%s'", field.Name)
					}

					if len(fieldValue.String()) > maxLen {
						return fmt.Errorf("'%s' must be at most %s characters long", field.Name, ruleValue)
					}
				}
			}
		}
	}

	// If no validation errors were found, return nil
	return nil
}
