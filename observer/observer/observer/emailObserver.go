package observer

import (
	"fmt"

	"github.com/low-level-design/observer/observable"
)

type emailObserver struct{
	observable observable.Observerable
}

func NewEmailObserver(ob observable.Observerable) *emailObserver{
	return &emailObserver{
		observable: ob,
	}
}

func (e *emailObserver) Update(){
	cnt := e.observable.GetData()

	fmt.Println("Sending notification on a mail, stockCnt: ", cnt)
}