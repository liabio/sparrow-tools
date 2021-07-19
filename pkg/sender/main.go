package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"k8s.io/klog/v2"
)

func main() {
	start := time.Now()

	defer func() {
		spend := time.Since(start).Seconds()
		klog.Infof("spend time: %.2f seconds", spend)
	}()
	rsp, err := http.Post("http://127.0.0.1:10000/inter/timeout/simulate", "", nil)
	if err != nil {
		klog.Fatal(err)
	}

	all, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info(string(all))

}

