package protocal

type Request struct {
	Id      string `json:"id"`
	Machine string `json:"machine"`
	User    string `json:"user"`
	Cmd     string `json:"cmd"`
	Time    string `json:"time"`
}

func (resq *Request) SetId(id string) {
	resq.Id = id
}

func (resq *Request) GetId() string {
	return resq.Id
}

func (resq *Request) SetMachine(machine string) {
	resq.Machine = machine
}

func (resq *Request) GetMachine() string {
	return resq.Machine
}

func (resq *Request) SetUser(user string) {
	resq.User = user
}

func (resq *Request) GetUser() string {
	return resq.User
}

func (resq *Request) SetCmd(cmd string) {
	resq.Cmd = cmd
}

func (resq *Request) GetCmd() string {
	return resq.Cmd
}

func (resq *Request) SetTime(time string) {
	resq.Time = time
}

func (resq *Request) GetTime() string {
	return resq.Time
}
