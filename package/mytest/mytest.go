package mytest

type Resp struct {
	Sum int `json`
}

func Add(a, b int) *Resp {
	var resp = Resp{
		Sum: a + b,
	}
	return &resp
}
