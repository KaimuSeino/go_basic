package main

import structreceiver "go-basics/struct-receiver"

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

type Task struct {
	Title    string
	Estimate int
}

func main() {
	// slicemap.Slice()
	// slicemap.Map()

	structreceiver.StructReceiver()

	// task1 := Task{
	// 	Title:    "hello go",
	// 	Estimate: 3,
	// }

	// task1.Title = "see you go"
	// fmt.Printf("Task: Type: %[1]T %+[1]v %v\n", task1, task1.Title)

	// var task2 Task = task1
	// task2.Title = "new title" // 別のメモリ領域にtask2の値が入っているのがわかる
	// fmt.Printf("task1_title: %v task2_title: %v\n", task1.Title, task2.Title)

	// task1_p := &Task{
	// 	Title:    "pointer_task1",
	// 	Estimate: 2,
	// }
	// fmt.Printf("task1_p: type: %T dereference_task1_p: %+v size: %v\n", task1_p, *task1_p, unsafe.Sizeof(task1_p))
	// task1_p.Title = "changed"
	// fmt.Printf("task1_p: %+v\n", *task1_p)
	// var task2_p *Task = task1_p
	// task2_p.Title = "Changed by task2"
	// fmt.Printf("task1: %+v\n", *task1_p)
	// fmt.Printf("task2: %+v\n", *task2_p)
}

// 一括に値を定義している
// var (
// 	i int
// 	s string
// 	b bool
// )

// fmt.Println(calculator.Offset)
// fmt.Println(calculator.Sum(2, 3))
// fmt.Println("Multiply: ", calculator.Multiply(3, 4))

// fmt.Println("hello world")
// sl := []int{1, 2, 3}
// if len(sl) > 0 {
// 	fmt.Println("unreachable code")
// }

// godotenv.Load()
// fmt.Println(os.Getenv("GO_ENV"))
