package model

import (
	"cron/protocal"
	"cron/utils"
	"fmt"
	"net/rpc"
	"sync"
	"time"
)

func PushToCrond(client *rpc.Client, wg *sync.WaitGroup, host string, port string, taskinfo TaskInfo) bool {
	defer wg.Done()
	if false == Parse(taskinfo.Exectime) {
		return false
	}

	request := new(protocal.Request)
	request.SetId(utils.UniqueId())
	request.SetUser("root")
	request.SetMachine("rs")
	request.SetCmd(taskinfo.Cmd)
	request.SetTime(time.Now().Unix())

	var response protocal.Response

	err := client.Call("ExecCron.RunShell", request, &response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Id: %s  En: %d  En: %s\n", response.GetId(), response.GetEn(), response.GetEm())
	return true
}
