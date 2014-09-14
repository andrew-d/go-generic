package chans

import (
	"sync"

	"github.com/joeshaw/gengen/generic"
)

// Merge any number of input channels into a single output channel.
func Merge(cs ...<-chan generic.T) <-chan generic.T {
	dummy := make(chan struct{})
	return MergeWithDone(dummy, cs...)
}

// Merge any number of input channels into a single output channel.  Allows
// the caller of this function to explicitly cancel this channel.
func MergeWithDone(done chan struct{}, cs ...<-chan generic.T) <-chan generic.T {
	var wg sync.WaitGroup
	out := make(chan generic.T)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan generic.T) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
