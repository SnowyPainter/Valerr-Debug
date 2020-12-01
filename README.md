# Valerr-Debug
Valerr Debug is something helpful when you test logics and functions.

# Why Valerr?
As you know, there are many ``` val, err := foo() ``` in Go.  
It is big problem that your test quickly.  
Valerr resolves these problems with ``` Run & Sequencing ```  

# Documentation

This simple documentation explains the core functions and how it works.

## Debugger
``` Go
debug := NewDebugger()
```
To use methods before, must call ```NewDebugger``` like above.
```usualErrorPosition``` is important to use debugger.  
Before you use methods, you must check where -th error return is.

## Run
Run a function seems like ignore error but, it return ```Result``` and it has ```Error``` handler.  
At last you can get the error and handle this.

``` Go
debug.Run(Chicken()).Pass(func(v interface{}) {
	fmt.Printf("Success Run :%v\n", v)
}).Error(func(e error) {
	fmt.Println("Failed to Run :", e.Error())
})
```

## Sequence
Call functions and process error by its default or handler passed through.
Sequence let you know what function made error. (by indexing)

***Simple Example with ForEach***
``` Go
debug.Sequence(Wrap(Chicken()), Wrap(Rose())).ForEach(func(i int, r Result) {
	if r.Err != nil {
		fmt.Println("Error At ", r.ErrWhere, ":", r.Err.Error())
		return
	}
	fmt.Println("Success Sequence at", i, ":", r.Value)
})
```

``` Go
func Chicken() (int, error) {
	return 0, errors.New("Error : chicken")
}

func Rose() (string, error) {
	return "I'm a rose", nil
}
```