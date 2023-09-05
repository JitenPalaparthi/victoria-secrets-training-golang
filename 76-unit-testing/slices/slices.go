// package slices is to create a new user defined package called slice
// this is a int slice with additional methods

package slices

import (
	"errors"
	"math/rand"
)

type SliceInt []int // SliceInt is a user defined type

// Fill fills the slice with random numbers
func (si SliceInt) Fill(len int) SliceInt {

	if si == nil {
		si = make(SliceInt, len)
	}

	for i := range si {
		si[i] = rand.Int()
	}
	return si

}

// InferDefault func infers the slice with a given value.
// normal slice by default it is infered with 0, provided by slice is int.
func (si SliceInt) InferDefaults(val int) SliceInt {

	if si == nil {
		return nil
	}

	for i := range si {
		si[i] = val
	}

	return si

}

// Max func returns the biggest value of the slice.
// if slice is nil then infer returns 0 and error
// if length of the slice is 1 then it returns the 0th element
func (si SliceInt) Max() (int, error) {

	if si == nil {
		return 0, errors.New("nil slice")
	}

	if len(si) == 1 {
		return si[0], nil
	}

	max := si[0]

	for i := 1; i < len(si); i++ {
		if max < si[i] {
			max = si[i]
		}

	}

	return max, nil
}
