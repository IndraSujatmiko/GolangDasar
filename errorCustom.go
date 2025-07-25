package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (x *notFoundError) Error() string {
	return x.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "Indra" {
		return &notFoundError{"data not found"}
	}
	// ok
	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		//terjadi error
		
			if validationErr, ok := err.(*validationError); ok {
				fmt.Println("validation error : ", validationErr.Error())
			} else if notFoundErr, ok := err.(*notFoundError); ok {
				fmt.Println("not found error : ", notFoundErr.Error())
			} else {
				fmt.Println("unknown error : ", err.Error())
			}
		

		// switch finalError := err.(type) {
		// case *validationError:
		// 	fmt.Println("validation error : ", finalError.Error())
		// case *notFoundError:
		// 	fmt.Println("not found error : ", finalError.Error())
		// default:
		// 	fmt.Println("unknown error : ", finalError.Error())
		// }
	} else {
		//sukses
		fmt.Println("Sukses")
	}

}
