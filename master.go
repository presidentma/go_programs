package main

import ("fmt"
"runtime"
		)

func main(){
	var arrBuf ArrayType=[]int{5,6,7,9}
	/* 指针pointer */
	a:="string";
	ptr:=&a;

	fmt.Printf("%s\n",*ptr);
	/* 结构体struct */
	type user struct{
		name string;
		age int;
	}
	tom:=user{
		name:"tom",
		age:22,
	}
	p:=&tom;
	fmt.Printf("%d\n",p.age);
	/* 数组array */
	arr:=[3]int{1,2,3};
	array:=[...]float32{1.0,2.0,3.0};
	arrayIndex:=[...]string{0:"first",2:"third"};
	fmt.Printf("%d,%f,%s\n",arr[0],array[0],arrayIndex[2]);
	for i,v:=range arr{
		fmt.Printf("index:%d,value:%d\n",i,v)
	}
	/* 切片slice */
	slice1:=arr[0:1]
	fmt.Printf("%v\n",slice1)
	slice2:=make([]int,2,3)
	slice2=append(slice2,5,6,7,8)
	copy(slice2,slice1)
	fmt.Printf("%v,len:%d,cap:%d\n",slice2,len(slice2),cap(slice2))
	str:="hello,go"
	strSliceByte:=[]byte(str)
	strSliceRune:=[]rune(str)
	fmt.Printf("%v\n%v\n",strSliceByte,strSliceRune)
	/* 字典map */
	map1:=map[string]int{"a":int('a'),"b":int('b')}
	fmt.Printf("%v\n",map1)
	fmt.Println(map1["a"])
	fmt.Println(map1["b"])
	map2:=make(map[string]int,5)
	map2["tom"]=22;
	map2["jerry"]=12;
	fmt.Println(map2)
	for k,v:=range map2{
		fmt.Printf("key:%s,value:%d\n",k,v)
	}
	delete(map2,"jerry")
	fmt.Println(map2)
	/* for */
	for i:=0;i<5;i++{
		// fmt.Println(i)n
	}
	for _,value:=range array{
		fmt.Println(value)
	}
	limt:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if j==1 {
				break limt
			}			
		}
	}
	/* func */
	swapNum :=func(a,b int)(int,int){
		return b,a
	}
	_=swapNum
	/* 不定参func */
	sum := func(arr ...int)(sum int){
		for _,v:=range arr{
			sum+=v
		}
		return
	}
	fmt.Println(sum(slice2...))
	/* func type */
	type cycle func(int,int)int
	
	pow:=func(a int,b int)int{
		c:=a
		for ;b>1;b--{
			a*=c
		}
		return a
	}
	init:=func(f cycle,a,b int) int {
		return f(a,b)
	}
	_=pow
	_=init
	fmt.Println(nil)
	/* defer */
	defer func(){
		recover()	//recover panic
		fmt.Println("exec over!")
	}()
	
	// panic("panic")
	/* 方法methods */
	/* func (t TypeName) MethodName (ParamList ) (Returnlist) {
		//method body
	} */
	arrBuf.foreach()
	var number1 Int = 10
	number1.set(5)	//方法值调用
	(*Int).set(&number1,15) //表达式调用
	fmt.Println("number1",number1)
	/*	接口interface */
	smtpInstance:=Smtp{account:"admin",key:"jugdejbggpre"}
	smtpInstance.Send()
	var mailer Mail=smtpInstance
	mailer.Recv()
	/* 检测实例类型是否是Smtp */
	if _,ok:=mailer.(Smtp);ok{
		fmt.Println("ok",ok)
	}
	/* 查询接口type */
	switch v:=mailer.(type){
	case nil:
		fmt.Println("type",nil)
	case Mail:
		fmt.Println("type","Mail")
	default:
		fmt.Println("type",v)
	}
	/* go例程goroutine */
	go func(){
		sum :=0
		for i:=0;i<100;i++{
			sum+=1
		}
		println("--------goroutine----------")
		println(sum)
	}()
	/* 管道chan */
	ch := make(chan struct{})
	ci := make(chan int,100)
	go func(ch chan struct{},ci chan int){
		println("-----------chan------------")
		sum :=0
		for i:=0;i<10;i++{
			ci<-i
		}
		close(ci)
		println(sum)
		ch <- struct{}{}
	}(ch,ci)
	<-ch	//读通道ch，通过通道进行同步等待,读通道后其匿名函数启动的goroutine会退出NumGoroutine
	println ("NumGoroutine=", runtime.NumGoroutine())
	/* 多路io select */
	chSelect:=make(chan int,1)
	go func(chan int){
		for{
			select{
			case chSelect<- 0:
			case chSelect<- 1:
			}
		}
	}(chSelect)
	for i:=0;i<5;i++{
		println(<-chSelect)
	}
	
}
/* 方法methods */
type ArrayType []int
func (array ArrayType) foreach() int{
	fmt.Println("foreach")
	for _,v:=range array{
		fmt.Println(v)
	}
	return 1
}
type Int int
func (a *Int)set(b Int)Int{
	*a=b
	return b
}
/*	接口interface */
type Mail interface{
	Send() bool
	Recv() bool
}
type Smtp struct{
	account string
	key string
}
func (s Smtp) Send() bool{
	fmt.Println("send mail")
	return true
}
func (s Smtp) Recv() bool{
	fmt.Println("received mail")
	return true
}



	