package shell

import (
	"bytes"
	"cron/protocal"
	"cron/utils"
	"crond/model"
	"encoding/json"
	"os/exec"
	"time"
)

type ExecCron int

func (t *ExecCron) RunShell(request *protocal.Request, response *protocal.Response) error {
	tmp, err1 := json.Marshal(request)
	if err1 == nil {
		utils.Write("INFO", string(tmp))
	}

	go ExecCmd(request)

	response.SetId(request.GetId())
	response.SetEn(200)
	response.SetEm("success")

	return nil
}

func ExecCmd(request *protocal.Request) {
	command := exec.Command("ssh", request.Machine, request.Cmd)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	var flag int
	var output string
	if err != nil {
		flag = -1
		output = err.Error()
		utils.Write("ERROR", output)
	} else {
		flag = 1
		output = out.String()
		utils.Write("INFO", output)
	}

	// add runlog
	model.AddRunLog(request.GetId(), flag, time.Now().Unix(), output)
}
