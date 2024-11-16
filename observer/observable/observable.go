package observable

import "github.com/low-level-design/observer/observer"

type Observerable interface{
	Add(ob observer.Observer)
	Remove(ob observer.Observer)
	Notify()
	SetCount(cnt int)
	GetData() int
}