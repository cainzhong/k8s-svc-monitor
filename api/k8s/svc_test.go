package k8s

import (
	"fmt"
	"testing"
	"time"
)

func TestGetK8sSvc(t *testing.T) {
	namespace := "itsma1"
	GetAllSvc(namespace)
}

func TestTimer(t *testing.T)  {
	go  timer1()
}

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()

}

func timer1() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}