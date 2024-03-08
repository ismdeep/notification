package core

import "fmt"

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func WarnIf(err error) {
	if err != nil {
		fmt.Println("[WARN]", err.Error())
	}
}
func IfErr(flag bool, err error) error {
	if flag {
		return err
	}
	return nil
}
