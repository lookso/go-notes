package utiltest

type Resp struct {
	Sum int `json:"sum"`
}

func Add(a, b int) *Resp {
	var resp = Resp{
		Sum: a + b,
	}
	return &resp
}
