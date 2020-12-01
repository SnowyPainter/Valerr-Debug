package main

import (
	"errors"
	"fmt"
)

//testing

func Chicken() (int, error) {
	return 0, errors.New("Error : chicken")
}

func Rose() (string, error) {
	return "I'm a rose", nil
}
func main() {
	debug := NewDebugger()

	debug.Run(Chicken()).Pass(func(v interface{}) {
		fmt.Printf("Success Run :%v\n", v)
	}).Error(func(e error) {
		fmt.Println("Failed to Run :", e.Error())
	})

	debug.Sequence(Wrap(Chicken()), Wrap(Rose())).ForEach(func(i int, r Result) {
		if r.Err != nil {
			fmt.Println("Error At ", r.ErrWhere, ":", r.Err.Error())
			return
		}
		fmt.Println("Success Sequence at", i, ":", r.Value)
	})
}
