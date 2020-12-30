package main

import (
	"fmt"
	"sync"
)
/* task work */
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
	/* 同步等待任务执行 */
	wait:=&sync.WaitGroup{}
	/* 初始化task goroutine */
	go InitTask(taskChan,resultChan,1000)
	/* 每个task启动一个goroutine */
	go DistributeTask(taskChan,wait,resultChan)
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
/* 读取task chan,每个task启动一个goroutine处理*/
func DistributeTask(taskChan <-chan task,wait *sync.WaitGroup,result chan int){
	for v:=range taskChan{
		wait.Add(1) //计数器加1
		go ProcessTask(v,wait)
	}
	wait.Wait() //等待结果
	close(result)
}
/* goroutine task */
func ProcessTask(t task,wait *sync.WaitGroup){
	t.do()
	wait.Done() //计数器-1
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
