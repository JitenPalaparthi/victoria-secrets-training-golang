package main

import (
	"errors"
	"fmt"
)

func main() {

	mmap := make(MyMap, 3)

	mmap["522001"] = "Guntur"
	mmap["560034"] = "Bangalore south"
	mmap["560086"] = "Bangalare west"

	if ok, err := mmap.Delete("522002"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ok)
	}

	if ok, err := mmap.Delete("522001"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ok, "key successfully deleted")
	}

	var mmap2 MyMap // This is nil

	if ok, err := mmap2.Delete("522002"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ok)
	}

	map1 := make(map[string]any) // Normal map

	map1["522001"] = "Guntur"
	map1["560034"] = "Bangalore south"
	map1["560086"] = "Bangalare west"

	if ok, err := MyMap(map1).Delete("522002"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ok)
	}

}

var (
	ERRNIlMAP = errors.New("nil map")
	ERRNOKEY  = errors.New("no key in the map")
)

type MyMap map[string]any

func (m MyMap) Delete(key string) (b bool, err error) {
	if m == nil {
		return false, ERRNIlMAP
	}

	_, ok := m[key]
	if !ok {
		return false, ERRNOKEY
	}
	delete(m, key)

	return true, nil

}
