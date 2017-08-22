//
// run tests:                            go test
// run tests with coverage:              ???
//
// run load test for 5 seconds:          go-wrk -d 5 http://localhost:8080/tcsiwula@gmail.com
//
// run benchmarks:                       go test -bench .
//
// profiling:                            go test -bench . -cpuprofile prof.cpu
//                                       && go-torch --binaryname cli.test -b profl.cpu
//                                       && open torch.svg
//

package main

import (
	"testing"
	"time"
)

// benchmark examples
func BenchmarkHandler(b *testing.B) {

	for i := 0; i < b.N; i++ {

		timeModI := time.Now().YearDay() + i
		//fmt.Print(timeModI)

		if timeModI == 0 {
			b.Error("error discovered...")
			b.Fatalf("Benchmarking error: %v", timeModI)
		}

	}

}

// test example
func TestHandler(t *testing.T) {

	// catch errors here
	// req, err
	// if err != nil {
	//
	// }

	// testing with coverage -- use go package test?
	//      see https://youtu.be/uBjoTxosSys?t=24m34s

	// cases := []struct {
	// 	in, out string
	// }{
	// 	{"tcsiwula@gmail.com", "gmail tcsiwula"},
	// 	{"ifWordOne", "thenWordTwo"},
	// }

	// no iterate over above cases to test them
	//    for _, c := range cases {
	// catch errors here
	// req, err
	//
	//    c.in
	//
	// if err != nil {
	//
	// }
	//}

	//      profiling****** go tool pprof -.....
	//       https://youtu.be/uBjoTxosSys?t=27m58s
	// find out more about what code is going on inside the code
	//  run your code and then while it is running you run pprof
	//  when in pprof, run web to see the call graph!!!!!!
	//
	// then we instrument the code
	// import _ "net/http/pprof"
	// then run, then view localhost:8080/debug/pprof
	//

}
