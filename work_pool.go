package main

import (
	"fmt"
)
/* work pool 个数 */
const NUMBER = 10

/* work task */
type task struct{
	begin int
	end int
	result chan<-int
}
/* task 添加do method */
func(t *task) do(){
	sum:=0
	for i:=t.begin;i<=t.end;i++{
		sum+=i
	}
	t.result<-sum
}
/* 计算和 */
func sum(){
	/* 创建任务通道 */
	taskChan:=make(chan task,10)
	/* 创建结果通道 */
	resultChan:=make(chan int,10)
	/* 创建worker 信号通道 */
	done:= make(chan struct{}, 10)
	/* 初始化task goroutine */
	go InitTask(taskChan,resultChan,10000)
	/* 每个task启动一个goroutine */
	workers:=NUMBER
	go DistributeTask(taskChan,workers,done)
	/* 获取各个goroutine 处理完任务的通知，并关闭结采通道 */
	go CloseResult(done,resultChan,workers)
	/* 通过结果通道获取结果并汇总 */
	sum:=ProcessResult(resultChan)
	fmt.Println("sum=",sum)
}

/* 构建task并写入task chan */
func InitTask(taskChan chan<- task,resultChan chan int,p int){
	qu:=p/10	//构建10个task分配给10个goroutine
	mod :=p%10
	hign:=qu*10
	for j:=0;j<qu;j++{
		b:=10*j+1
		e:=10*(j+1)
		tsk:=task{
			begin:b,
			end:e,
			result:resultChan,
		}
		taskChan <- tsk
	}
	// 余数
	if mod !=0{
		tsk:=task{
			begin:hign+1,
			end:p,
			result:resultChan,
		}
		taskChan<- tsk
	}
	close(taskChan)
}
/* 读取task chan,并分发到worker goroutine 处理，总的数量是workers*/
func DistributeTask(taskChan <-chan task, workers int,done chan struct{}){
	for i:=0;i<workers;i++{
		go ProcessTask(taskChan,done)
	}
}
/* goroutine task */
func ProcessTask(taskChan <-chan task,done chan struct{}){
	for t:=range taskChan{
		t.do()
	}
	done<-struct{}{}	//通知已经done
}
/* 通过done channel 同步等待所有工作goroutine 的结束，然后关闭结果chan */
func CloseResult(done chan struct{},resultChan chan int,workers int){
	for i:=0;i<workers;i++{
		<-done
	}
	close(done)
	close(resultChan)
}
/* goroutine汇总结果 */
func ProcessResult(resultChan chan int)int{
	sum :=0
	i:=0;
	for r:=range resultChan {
		sum+=r
		i+=1
	}
	fmt.Println(i)
	return sum
}

func main(){
	sum()
}