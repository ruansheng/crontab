package model

import (
	"cron/protocal"
	//"encoding/json"
	"fmt"
	"net"
	//"net/rpc"
	"cron/utils"
	"net/rpc/jsonrpc"
	"sync"
)

func connectCrond(host string, port string) *net.Conn {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &conn
}

func PushToCrond(wg *sync.WaitGroup, host string, port string, taskinfo TaskInfo) bool {
	defer wg.Done()
	if false == Parse(taskinfo.Exectime) {
		return false
	}

	conn := connectCrond(host, port)
	defer (*conn).Close()
	client := jsonrpc.NewClient(*conn)
	defer client.Close()

	request := new(protocal.Request)
	request.SetId(utils.UniqueId())
	request.SetUser("root")
	request.SetMachine("rs")
	request.SetCmd(taskinfo.Cmd)
	request.SetTime("1384534678")

	var response protocal.Response

	err := client.Call("ExecCron.RunShell", request, &response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Id: %s  En: %s  En: %s Data: %s\n", response.GetId(), response.GetEn(), response.GetEm(), response.GetData())
	return true
}
