package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)
/* 随机整数生成 */
func GenerateIntA(done chan struct{}) chan int{
	ch:=make(chan int)
	go func(){
		Label:
			for{
				select{
				case ch<- rand.Int():
				case <- done:
					break Label
				}
			}
			close(ch)
	}()
	return ch
}

func GenerateIntB() chan int{
	ch:=make(chan int,10)
	go func(){
			for{
				ch<- rand.Int()
			}
	}()
	return ch
}

func GenerateInt(){
	done:=make(chan struct{})
	ch:=GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	println ("NumGoroutine=", runtime.NumGoroutine())
}
/* chain调用 */
func chain(in chan int) chan int{
	out:=make(chan int)
	go func(){
		for v:=range in{
			out<-1+v
		}
		close(out)
	}()
	return out
}

func ChainCall(){
	in:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			in<-i
		}
		close(in)
	}()
	//输入输出chan类型相同可以形成chain调用
	out:=chain(chain(in))
	for v:=range out{
		fmt.Println(v)
	}
}


func main(){
	// ChainCall()
	wait:=sync.WaitGroup{}
	arr:=[]int{1,2,3,4,5,6,7,8,9,10}
	for _,v:=range arr{
		wait.Add(1)
		go func(v int){
			fmt.Println(v)
			wait.Done()
		}(v)
	}
	wait.Wait()
}