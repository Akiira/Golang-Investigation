package Version2

type empty struct {}

type Semaphore struct{
	resources chan empty
}

func New(resources int) *Semaphore {
	gSem := Semaphore{ resources: make(chan empty, resources) }
	for i:= 0; i < resources; i++ {
		gSem.resources <- empty{}
	}
	
	return &gSem
}

func (sem *Semaphore) Wait() {
	<- sem.resources
}

func (sem *Semaphore) Signal() {
	sem.resources <- empty{}
}