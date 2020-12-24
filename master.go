package main

import "fmt"

func main(){
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
	panic("panic")
	
	
	


}