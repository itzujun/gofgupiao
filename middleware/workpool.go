package middleware
// 2019.02.22
type WorkPool struct {
	//
}

func NewWorkPool() *WorkPool {
	return &WorkPool{}
}

func (this *WorkPool) Pool(num int, work func()) {
	for i := 0; i < num; i++ {
		go work()
	}
}
