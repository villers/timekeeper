package models

import "fmt"

type ProcessInfo struct {
	Pid         int32
	AppName     string
	WindowName  string
	ProcessName string
	Url         string
}

func (p ProcessInfo) ToString() string {
	return fmt.Sprintf("%v\n%s\n%s\n%s\n%s\n", p.Pid, p.WindowName, p.AppName, p.ProcessName, p.Url)
}
