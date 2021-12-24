package pkg

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"time"
//)
//
//func Timeout(f func() gin.Context) gin.Context {
//	timeoutChan := make(chan bool, 1)
//
//	go func() {
//		f()
//		timeoutChan <- false
//	}()
//
//	go func() {
//		time.Sleep(1 * time.Second)
//		timeoutChan <- true
//	}()
//
//	if <-timeoutChan {
//		fmt.Println("Timeout")
//		return _
//	}
//	return _
//
//}
