package thingiverseio_test

import (
	"log"
	"testing"

	"github.com/ThingiverseIO/thingiverseio"
)

const descriptor = `
func Benchmark() ()
`

func BenchmarkNilCall(b *testing.B) {

	in, err := thingiverseio.NewInput(descriptor)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Remove()


	o, err := thingiverseio.NewOutput(descriptor)
	if err != nil {
		log.Fatal(err)
	}

	defer o.Remove()
	rc := o.Requests().AsChan()

	in.Run()
	o.Run()

	// for i := 0; i < b.N; i++ {
		c := in.Call("SayHello", struct{}{}).AsChan()

		o.Reply(<-rc, struct{}{})
		<-c
	// }
}
