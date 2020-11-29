package adcal2020

var WorkerChannel = make(chan chan task)

type Worker struct {
	WorkerChannel chan chan task // used to communicate between dispatcher and workers
	Channel       chan task
	End           chan bool
}

// start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case t := <-w.Channel: // worker has received job
				t()
			case <-w.End:
				return
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	w.End <- true
}

type Collector struct {
	Work chan task // receives jobs to send to workers
	End  chan bool // when receives bool stops workers
}

func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan task) // channel to recieve work
	end := make(chan bool)   // channel to spin down workers
	collector := Collector{Work: input, End: end}

	for i < workerCount {
		i++
		worker := Worker{
			Channel:       make(chan task),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool)}
		worker.Start()
		workers = append(workers, worker) // stores worker
	}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop() // stop worker
				}
				return
			case t := <-input:
				worker := <-WorkerChannel // wait for available channel
				worker <- t               // dispatch work to worker
			}
		}
	}()

	return collector
}
