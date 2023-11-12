package example

import "fmt"

type SaveLogInterface interface {
	saveLog()
}

func saveLog(st SaveLogInterface) {
	st.saveLog()
}

type TransferBBL struct {
	name string
}

func (tBBL *TransferBBL) saveLog() {
	fmt.Println("save to database", tBBL.name)
}

type TransferKTB struct {
	name string
}

func (tKTB *TransferKTB) saveLog() {
	fmt.Println("save to database", tKTB.name)
}

func main2() {
	transA := TransferBBL{name: "BBl"}
	transB := TransferKTB{name: "KTB"}
	saveLog(&transA)
	saveLog(&transB)
}
