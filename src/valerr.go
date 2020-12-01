package main

import (
	"fmt"
)

//No statics

const (
	//NoWhere mean there is no error
	NoWhere = -1
)

//Result returned by Run & Sequence
type Result struct {
	Value    interface{}
	Err      error
	ErrWhere int
}

//Debug as container
type Debug struct {
	UsualErrorPosition int
}

//Results defineded type that array of Result
type Results []Result

//CreateResultDefault return basic result as it default set
func CreateResultDefault() Result {
	return Result{
		Value:    nil,
		Err:      nil,
		ErrWhere: NoWhere,
	}
}

//NewDebugger return new debugger as it default set
func NewDebugger() Debug {
	return Debug{
		UsualErrorPosition: 1,
	}
}

//Wrap return interface array and pass multiple value return function result.
func Wrap(t ...interface{}) []interface{} {
	return t
}

//Pass pass the result if it suceesses
func (result Result) Pass(handler func(interface{})) Result {
	if result.Err == nil {
		handler(result.Value)
	}
	return result
}

func (result Result) Error(handler func(error)) {
	if result.Err != nil {
		handler(result.Err)
	}
}

//ForEach return each result
func (results Results) ForEach(handler func(int, Result)) {
	for i, r := range results {
		//if err != nil, i == r.ErrWhere
		handler(i, r)
	}
}

//Run runs a function
func (Debug Debug) Run(val interface{}, err error) (r Result) {
	r = CreateResultDefault()
	if err != nil {
		r.ErrWhere = 0
		r.Err = err
	}
	return
}

//Sequence runs same positioned error value functions
func (Debug Debug) Sequence(results ...[]interface{}) (rs Results) {
	rs = make([]Result, 0)
	for i, r := range results {
		rs = append(rs, CreateResultDefault())
		if r[Debug.UsualErrorPosition] != nil { // error
			rs[i].ErrWhere = i
			rs[i].Err = fmt.Errorf("%v", r[Debug.UsualErrorPosition])
			continue
		}

		rs[i].Value = r[0:Debug.UsualErrorPosition]
	}
	return
}
