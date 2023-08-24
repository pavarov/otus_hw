package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	res := make(Bi)
	go func() {
		defer close(res)
		for _, stage := range stages {
			in = stage(in)
		}

		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				res <- v
			}
		}
	}()

	return res
}
