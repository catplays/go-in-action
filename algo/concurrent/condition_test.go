package concurrent

import "testing"

func TestCondition(t *testing.T) {
	for i:=0;i<10;i++ {
		go func() {
			SendMail()
		}()
	}
	for i:=0;i<10;i++ {
		 RecvMail()
	}
}
