package unittest

func Add(x, y int) int {
	return x + y
}

func Divide(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}

func UnitTest() {
	// Addを選択して右クリック→Go: Generate unit test for Functionをクリック
	// unitTest_test.goが作成される。
	// cd unit-testディレクトリに移動してから、go test -v .ターミナルで実行

	// テストコードがどれだけソースコードをカバーできているか確かめる。 Divide関数
	// go test -v -cover -coverprofile=coverage.out . を実行
	// coverageが100%ではない場合は、go tool cover -html=coverage.out を実行

}
