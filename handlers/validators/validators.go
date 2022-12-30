package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
)

type FieldConstraint struct {
	FieldName         string
	DataType          string
	IsRequired        bool
	IsNegativeAllowed bool
}

func isFieldIsNeeded(field_name string, field_constraints []FieldConstraint) bool {
	for _, v := range field_constraints {
		if v.FieldName == field_name {
			return true
		}
	}
	return false
}

func validateFields(field_constraints []FieldConstraint, request_body map[string]any) error {
	missing_fields := []string{}
	invalid_datatype_fields := []string{}
	negative_fields := []string{}
	invalid_fields := []string{}

	for k := range request_body {
		if !isFieldIsNeeded(k, field_constraints) {
			invalid_fields = append(invalid_fields, k)
		}
	}

	for _, v := range field_constraints {
		request_body_map_field_value := request_body[v.FieldName]

		if request_body_map_field_value != nil && fmt.Sprintf("%T", request_body_map_field_value) != v.DataType {
			invalid_datatype_fields = append(invalid_datatype_fields, v.FieldName)
		}

		if v.IsRequired && request_body_map_field_value == nil {
			missing_fields = append(missing_fields, v.FieldName)
		}

		if request_body_map_field_value != nil && fmt.Sprintf("%T", request_body_map_field_value) == "float64" {
			if !v.IsNegativeAllowed && request_body_map_field_value.(float64) < 0 {
				negative_fields = append(negative_fields, v.FieldName)
			}
		}
	}

	if len(invalid_fields) != 0 {
		return errors.New(responses.InvalidFieldsInRequestBody + strings.Join(invalid_fields, ", "))
	}

	if len(missing_fields) != 0 {
		return errors.New(responses.MissingFields + strings.Join(missing_fields, ", "))
	}

	if len(invalid_datatype_fields) != 0 {
		return errors.New(responses.InvalidDataTypeForFieldInJSONBody + strings.Join(invalid_datatype_fields, ", "))
	}

	if len(negative_fields) != 0 {
		return errors.New(responses.FollowingFieldsCannotBeNegative + strings.Join(negative_fields, ", "))
	}

	return nil
}

func ValidateRequestBody(request_body io.Reader, field_constraints []FieldConstraint) ([]byte, error) {
	request_body_json, err := io.ReadAll(request_body)
	if err != nil {
		return nil, err
	}

	request_body_map := map[string]any{}
	err = json.Unmarshal(request_body_json, &request_body_map)
	if err != nil {
		return nil, errors.New(responses.BadRequestBody)
	}

	err = validateFields(field_constraints, request_body_map)
	if err != nil {
		return nil, err
	}

	return request_body_json, nil
}

func ValidateRequestBodyMap(request_body_map map[string]any, field_constraints []FieldConstraint) error {
	err := validateFields(field_constraints, request_body_map)
	if err != nil {
		return err
	}

	return nil
}
