package core

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func IfErr(flag bool, err error) error {
	if flag {
		return err
	}
	return nil
}
