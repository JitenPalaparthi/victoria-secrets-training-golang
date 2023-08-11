package tostring

import "fmt"

type Cint int

func (c Cint) ToString() string {
	return fmt.Sprint(c)
}

type Cfloat float32

func (c Cfloat) ToString() string {
	return fmt.Sprint(c)
}
