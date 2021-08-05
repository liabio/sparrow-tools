package parallelize

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type DoWorkFunc func(s string)

func TestParaller(t *testing.T) {
	newpr := []string{"a", "b", "c", "d", "e"}
	//length := len(newpr)

	ctx, cancel := context.WithCancel(context.Background())

	var errs []error
	restorJob := func(pvName string) {
		err := dojob1(pvName)
		if err != nil {
			cancel()
			errs = append(errs, err)
			t.Log("执行报错", err)
		}
		t.Log("执行完毕", pvName)
	}

	ParallelizeDoJob(ctx, 2, newpr, restorJob)
}
func dojob1(s string) error {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(20)
	fmt.Printf("%v--pv，sleep时间为%v\n", s, result)
	if result > 10 {
		return errors.New("大于5秒超时，返回错误")
	}
	time.Sleep(time.Second * time.Duration(result))
	return nil
}

func ParallelizeDoJob(ctx context.Context, works int, pr []string, do DoWorkFunc) {
	var stop <-chan struct{}
	if ctx != nil {
		stop = ctx.Done()
	}
	length := len(pr)
	toProcess := make(chan string, length)
	for _, volume := range pr {
		toProcess <- volume
	}
	close(toProcess)

	wg := sync.WaitGroup{}
	wg.Add(works)
	for i := 0; i < works; i++ {
		go func() {
			defer wg.Done()
			for pvName := range toProcess {
				select {
				case <-stop:
					return
				default:
					do(pvName)
				}
			}
		}()
	}
	wg.Wait()
}
