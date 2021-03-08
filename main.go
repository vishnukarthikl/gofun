package main

import (
	"fmt"

	"k8s.io/klog"
	"rsc.io/quote"
)

func main() {
	klog.Info("Hi")
	fmt.Println(quote.Hello())
}
