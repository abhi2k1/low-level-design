package main

import (
	"github.com/low-level-design/observer/observable"
	"github.com/low-level-design/observer/observer/observer"
)


func main(){
	iphone := observable.NewIphone()


	emailObserver := observer.NewEmailObserver(iphone)

	mobileObserver := observer.NewMobileObserver(iphone)

	iphone.Add(emailObserver)
	iphone.Add(mobileObserver)

	iphone.Notify()

	// setting stock count

	iphone.SetCount(10)
}