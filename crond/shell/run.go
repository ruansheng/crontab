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

	ret := ExecCmd(request)

	*response = *ret

	jtmp, err2 := json.Marshal(response)
	if err2 == nil {
		utils.Write("INFO", string(jtmp))
		fmt.Printf("%s\n", string(jtmp))
	}

	return nil
}

func ExecCmd(request *protocal.Request) *protocal.Response {
	response := new(protocal.Response)
	response.SetId(request.GetId())

	command := exec.Command("ssh", request.Machine, request.Cmd)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		utils.Write("ERROR", err.Error())
		fmt.Println("error:", err.Error())
		response.SetEn("501")
		response.SetEm("fail")
		response.SetData(err.Error())
	} else {
		response.SetEn("200")
		response.SetEm("success")
		response.SetData(out.String())
	}

	utils.Write("INFO", out.String())
	fmt.Println("output:", out.String())

	return response
}
