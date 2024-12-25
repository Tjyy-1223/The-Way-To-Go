package main

import "fmt"

type Any interface{}
type Evaluator func(Any) (Any, Any)

func BuildLazyEvaluator(evaluator Evaluator, initState Any) func() Any {
	in := make(chan Any)
	loopFunc := func() {
		s := initState
		var res Any
		for {
			_, res = evaluator(s)
			s = Any(s.(uint) + 1)
			in <- res
		}
	}

	retFunc := func() Any {
		return <-in
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evaluator Evaluator, initState Any) func() uint {
	ef := BuildLazyEvaluator(evaluator, initState)
	return func() uint {
		return ef().(uint)
	}
}

func fibonacci(k uint) uint {
	if k <= 1 {
		return 1
	} else {
		return fibonacci(k-1) + fibonacci(k-2)
	}
}

func mainGenerateLazyEvaluation2() {
	fiboFunc := func(state Any) (Any, Any) {
		s := state.(uint)
		return s, fibonacci(s)
	}

	fibonacci := BuildLazyIntEvaluator(fiboFunc, uint(0))
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fibonacci: %v\n", i, fibonacci())
	}
}
