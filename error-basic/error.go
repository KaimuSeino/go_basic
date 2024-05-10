package errorbasic

import (
	"errors"
	"fmt"
	"os"
)

// センチネルエラー
var ErrCustom = errors.New("not found")

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}

func Error() {
	err01 := errors.New("something wrong")
	err02 := errors.New("something worng")
	// エラーの型
	fmt.Printf("error pointer: %[1]p type: %[1]T value: %[1]v\n", err01)
	fmt.Println(err01)
	fmt.Println(err01 == err02)

	// エラーのラップ
	err0 := fmt.Errorf("add info: %w", errors.New("original error"))
	fmt.Printf("err0 pointer: %[1]p type: %[1]T value: %[1]v\n", err0)
	fmt.Println(errors.Unwrap(err0))

	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)
	fmt.Printf("err1 type: %T\n", err1)
	fmt.Println(errors.Unwrap(err1))

	// エラーのカスタム
	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	// err2の重ねがけ
	err2 = fmt.Errorf("in repository layer: %w", err2)
	fmt.Println(err2)

	// errors.Isでboolが帰ってくる
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}

}
