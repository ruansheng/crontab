package model

import (
	"cron/protocal"
	"cron/utils"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

func DoTask(host string, port string) {
	var wg sync.WaitGroup

	// jsonrpc client
	conn := connectCrond(host, port)
	defer (*conn).Close()
	client := jsonrpc.NewClient(*conn)
	defer client.Close()

	// get machines
	machines := getMachinesMap()

	// get tasks
	taskinfos := GetTaskData()

	// push to crond
	for _, taskinfo := range taskinfos {
		if checkIsRun(taskinfo, machines) {
			wg.Add(1)
			PushToCrond(client, &wg, host, port, taskinfo, machines)
		}
	}
	wg.Wait()
}

func connectCrond(host string, port string) *net.Conn {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &conn
}

func checkIsRun(taskinfo TaskInfo, machines map[string]MachineInfo) bool {
	// check is run
	return true
}

func PushToCrond(client *rpc.Client, wg *sync.WaitGroup, host string, port string, taskinfo TaskInfo, machines map[string]MachineInfo) bool {
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
