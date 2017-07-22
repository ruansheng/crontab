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

	// get tasks
	results := GetTaskData()

	// push to crond
	for _, result := range results {
		for i := 0; i < 50; i++ {
			wg.Add(1)
			PushToCrond(client, &wg, host, port, result)
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
