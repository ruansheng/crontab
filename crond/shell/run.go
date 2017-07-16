package shell

import (
	"bytes"
	"cron/protocal"
	"cron/utils"
	"encoding/json"
	"fmt"
	"os/exec"
)

type ExecCron int

func (t *ExecCron) RunShell(request *protocal.Request, response *protocal.Response) error {
	tmp, err1 := json.Marshal(request)
	if err1 == nil {
		utils.Write("INFO", string(tmp))
		fmt.Printf("%s\n", string(tmp))
	}

	go ExecCmd(request)

	return nil
}

func ExecCmd(request *protocal.Request) {
	command := exec.Command("ssh", request.Machine, request.Cmd)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		utils.Write("ERROR", err.Error())
		fmt.Println("error output:", err.Error())
	} else {
		fmt.Println("success output:", out.String())
	}
}
