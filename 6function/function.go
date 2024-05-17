package function

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Function() {
	funcDefer()
	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(files)
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	// 無名関数（即座に実行される）
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)
	// f1を定義した場所で実行される関数
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExtension(f2, "file1")

	f3 := multiply()
	fmt.Println(f3(3))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(4)
		fmt.Println(v)
	}
}

// defer
func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello func defer")
}

// trim
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	fmt.Printf("out: value: %v length: %v capacity: %v\n", out, len(out), cap(out))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

// file checker 返り値はstringとerror
func fileChecker(name string) (string, error) {
	f, err := os.Open(name) // functionファイルにfile.txtがあるとerrになる
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

// .csvを追加 引数には func(file string)でその返り値がstringの関数
func addExtension(f func(file string) string, name string) {
	fmt.Println(f(name))
}

// 返り値に無名関数を返す
func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

// closure
func countUp() func(int) int {
	count := 0 //グローバル変数っぽい、でもグローバル変数ではないので値の変更の心配なし
	return func(n int) int {
		count += n
		return count
	}
}
