package main

import (
	"dumdum_client/collectors/rabbit"
	"encoding/json"
	"fmt"
	"time"

	"github.com/prometheus/procfs"
)

func main() {

	fs, err := procfs.NewFS("/proc")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		procs, err := fs.CPUInfo()
		if err != nil {
			fmt.Println(err.Error())
		}
		prc, err := json.Marshal(&procs)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			rabbit.SendData(string(prc))
			fmt.Println("Data sent!")
		}()

		time.Sleep(time.Minute)

	}
}
