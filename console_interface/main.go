package main

import (
	"errors"
	"fmt"

	"github.com/subramanyam-searce/product-catalog-go/console_interface/ops"
)

func main() {
	var quit_menu bool

	for !quit_menu {
		var err error
		err, quit_menu = MainMenu()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
	}
}

func MainMenu() (error, bool) {
	var user_choice uint
	var quit_menu bool

	for i, v := range ops.HomePageOptions {
		fmt.Printf("%v.%v\n", i+1, v.DisplayName)
	}

	user_choice_float, err := ops.GetPositiveFloatFromConsole(ops.ChoiceInput)
	user_choice = uint(user_choice_float)

	if err != nil {
		return err, quit_menu
	}

	index := user_choice - 1

	if index == 0 {
		quit_menu = true
		return nil, quit_menu
	}

	if index >= uint(len(ops.HomePageOptions)) {
		return errors.New(ops.InvalidChoice), quit_menu
	}

	ops.HomePageOptions[index].Handler()
	return nil, quit_menu
}
