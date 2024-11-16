package observable

import "github.com/low-level-design/observer/observer"

type iphone struct{
	stockCnt int

	observer []observer.Observer
}


func NewIphone() *iphone{
	return &iphone{
		stockCnt: 0,
		observer: make([]observer.Observer,0),
	}
}

func (i *iphone) Add(ob observer.Observer){
	i.observer = append(i.observer, ob)
}

func (i *iphone) Remove(ob observer.Observer){
	// remove from the list
}

func (i *iphone) Notify(){
	for _, ob := range i.observer{
		ob.Update()
	}
}

func (i *iphone) SetCount(cnt int){
	i.stockCnt = cnt

	if cnt > 0{
		i.Notify()
	}
}

func (i *iphone) GetData() int{
	return i.stockCnt
}