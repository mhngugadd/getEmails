package main

import (
	"github.com/mhngugadd/getEmail/file"
	"github.com/mhngugadd/getEmail/worker"
	"fmt"
)

func main() {
	maxWorkers := 100

	// 获取文件列表
	var DIR  = file.GetCurrentDirectory()
	// 获取文件列表名称
	list , _:= file.GetAllFile(DIR)
	// 初始化工作任务，指定最大列队数
	work := file.NewReadJob(maxWorkers,len(list))
	// 初始化工作任务通道
	jobs := make(chan worker.Job)
	// 生成工作任务
	jobs <- file.ReadFileContent(list , work)
	result := make(chan []string,len(list))
	for i := 0; i < 10 ; i++  {
		go worker.Worker(jobs,result)
	}
	for i := 0; i < len(result) ; i++{
		fmt.Println(<-result)
	}
}