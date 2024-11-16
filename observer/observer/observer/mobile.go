package observer

import (
	"fmt"

	"github.com/low-level-design/observer/observable"
)

type mobileObserver struct{
	observable observable.Observerable
}

func NewMobileObserver(ob observable.Observerable) *emailObserver{
	return &emailObserver{
		observable: ob,
	}
}

func (m *mobileObserver) Update(){
	cnt := m.observable.GetData()

	fmt.Println("Sending notification on a mobile, stockCnt: ", cnt)
}