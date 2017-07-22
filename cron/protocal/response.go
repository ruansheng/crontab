package protocal

type Response struct {
	Id   string      `json:"id"`
	En   int         `json:"en"`
	Em   string      `json:"em"`
	Data interface{} `json:"data"`
}

func (resp *Response) SetId(id string) {
	resp.Id = id
}

func (resp *Response) GetId() string {
	return resp.Id
}

func (resp *Response) SetEn(en int) {
	resp.En = en
}

func (resp *Response) GetEn() int {
	return resp.En
}

func (resp *Response) SetEm(em string) {
	resp.Em = em
}

func (resp *Response) GetEm() string {
	return resp.Em
}

func (resp *Response) SetData(data interface{}) {
	resp.Data = data
}

func (resp *Response) GetData() interface{} {
	return resp.Data
}
