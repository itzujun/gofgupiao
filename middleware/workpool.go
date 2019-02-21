package middleware

type WorkPool struct {
	//
}

func (this *WorkPool) Pool(num int, work func()) {
	for i := 0; i < num; i++ {
		go work()
	}
}
