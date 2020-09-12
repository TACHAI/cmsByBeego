package utils

import (
	"fmt"
	"reflect"
)

func mainUtils()  {
	makeChan()
}

func makeSlice()  {
	mSlice:=make([]string,3)
	mSlice[0]="3"
	mSlice[1]="2"
	mSlice[2]="1"

	fmt.Println(mSlice)
}

// makeMap 创建map
func makeMap()  {
	mMap:=make(map[int]string)
	mMap[10]="dog"
	mMap[100]="cat"
	fmt.Println(mMap)
}

//创建 chan（管道） 后面缓存是3
func makeChan()  {
	mChan:=make(chan int,3)
	close(mChan)

}

func NewMap()  {
	mMap:=new(map[int]string)
	fmt.Println(reflect.TypeOf(mMap))
}