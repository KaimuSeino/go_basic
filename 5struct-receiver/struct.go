package structreceiver

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func StructReceiver() {
	task1 := Task{
		Title:    "hello go",
		Estimate: 3,
	}

	task1.Title = "hello gogo"
	fmt.Printf("Task: Type: %[1]T %+[1]v %v\n", task1, task1.Title)

	var task2 Task = task1
	task2.Title = "new Title"
	fmt.Printf("task1: %v task2: %v\n", task1.Title, task2.Title)

	task1_p := &Task{
		Title:    "pointer_task1",
		Estimate: 3,
	}
	fmt.Printf("task1_p: type: %T dereference: %+v size: %v\n", task1_p, *task1_p, unsafe.Sizeof(task1_p))
	task1_p.Title = "changed"
	fmt.Printf("task1_p: %+v\n", *task1_p)
	var task2_p *Task = task1_p
	task2_p.Title = "Changed by task2"
	fmt.Printf("task1: %+v\n", *task1_p)
	fmt.Printf("task2: %+v\n", *task2_p)

	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
	(&task1).extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)

}

// 値receverを持つメソッドを定義
// 値receiverは受け取った構造体のコピーに対してメソッド内の操作を行うので値が変更されない。
func (task Task) extendEstimate() { // このメソッドが受け取るある型から生成された実態のことをreceiverという
	// 特定の型から生成された実態に対してメソッドを追加できる
	task.Estimate += 10
}

func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
