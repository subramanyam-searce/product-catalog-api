package ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/subramanyam-searce/product-catalog-go/handlers/validators"
)

func GetPositiveFloatFromConsole(message string) (float64, error) {
	var input_float float64
	var err error

	for input_float <= 0 {
		fmt.Print(message)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_string := scanner.Text()

		input_float, err = strconv.ParseFloat(input_string, 64)

		if err != nil {
			fmt.Println(EnterValidFloat)
			fmt.Println()
			continue
		}

		if input_float <= 0 {
			fmt.Println(EnterPositiveValue)
			fmt.Println()
		}
	}

	if err != nil {
		return input_float, err
	}

	return input_float, nil
}

func FormatPrintStruct(v any) {
	s := structs.New(v)

	fmt.Println(Divider)
	for _, v := range s.Fields() {

		if v.Kind().String() == "slice" {
			continue
		}

		fmt.Printf("%v: %v\n", LeftFormat(v.Name()), s.Field(v.Name()).Value())
	}
	fmt.Println(Divider)
}

func LeftFormat(str string) string {
	spaces := ""
	for i := 0; i < SizeOfOutputStructFields; i++ {
		spaces += " "
	}

	return (str + spaces)[:SizeOfOutputStructFields]
}

func GetFieldNameWithJSONTag(s *structs.Struct, tag string) string {
	for _, v := range s.Fields() {
		if v.Tag("json") == tag {
			return v.Name()
		}
	}

	return ""
}

func ScanField(field_name string) (string, error) {
	var user_input string
	fmt.Print(LeftFormat(field_name) + Colon)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user_input = scanner.Text()

	return user_input, nil
}

func GetInputBody(s *structs.Struct, field_constraints []validators.FieldConstraint) error {
	for _, v := range field_constraints {
		field_name_in_go := GetFieldNameWithJSONTag(s, v.FieldName)
		if v.DataType == "float64" {
			user_input, err := GetPositiveFloatFromConsole(LeftFormat(field_name_in_go) + Colon)
			if err != nil {
				return err
			}

			if s.Field(field_name_in_go).Kind().String() == "int" {
				err = s.Field(field_name_in_go).Set(int(user_input))
			} else {
				err = s.Field(field_name_in_go).Set(user_input)
			}

			if err != nil {
				return err
			}

			continue
		}

		user_input, err := ScanField(field_name_in_go)
		if err != nil {
			return err
		}

		if v.DataType == "map[string]interface {}" {
			user_input = strings.ReplaceAll(user_input, " ", "")
			for !json.Valid([]byte(user_input)) {
				fmt.Println(EnterValidJson)
				user_input, err = ScanField(field_name_in_go)
				if err != nil {
					fmt.Println(err)
				}
			}

			json_item := map[string]any{}
			err := json.Unmarshal([]byte(user_input), &json_item)
			if err != nil {
				return err
			}

			err = s.Field(field_name_in_go).Set(json_item)
			if err != nil {
				return err
			}

			continue
		}

		err = s.Field(field_name_in_go).Set(user_input)
		if err != nil {
			return err
		}

	}

	return nil
}
