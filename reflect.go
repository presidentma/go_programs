package main

import (
	"reflect"
	"fmt"
)

type Student struct{
	Name string "学生姓名"
	Age int `tom:"15"jerry:"16"`
}

func main(){
	rt:=reflect.TypeOf(Student{})
	fieldName,ok:=rt.FieldByName("Name")
	if ok{
		fmt.Println(fieldName.Tag)
	}
	fieldAge,okey:=rt.FieldByName("Age")
	if okey{
		fmt.Println(fieldAge.Tag.Get("tom"))
	}
	fmt.Println("type_Name",rt.Name())
	fmt.Println("type_NumField",rt.NumField())
	fmt.Println("type_PkgPath",rt.PkgPath())
	fmt.Println("type_String",rt.String())
	fmt.Println("type.Kind.String",rt.Kind().String())
	fmt.Println("type.String()",rt.String())
	/* struct字段名称 */
	for i := 0 ; i< rt.NumField();i++ {
		fmt.Printf("type.Field[%d].Name:=%v\n",i,rt.Field(i).Name)
	}
	fmt.Println("------------------")
	rv:=reflect.ValueOf(Student{"tom",15})
	t:=rv.Type()
	fmt.Println("t.Name",t.Name())
	for i:= 0;i<t.NumField();i++{
		field:=t.Field(i)
		value:=rv.Field(i).Interface()
		switch value:=value.(type){
		case int:
			fmt.Printf("%6s:%v=%d\n",field.Name,field.Type,value)
		case string:
			fmt.Printf("%6s:%v=%s\n",field.Name,field.Type,value)
		default:
			fmt.Printf("%6s:%v=%s\n",field.Name,field.Type,value)
		}
	}
}
