package file

import (
	"regexp"
	"errors"
)

type ReadJob struct {
	Content chan []byte
	exitChan chan error
}

func NewReadJob(maxWorkers int , allJobs int) *ReadJob  {
	return &ReadJob{
		Content  : make(chan []byte , maxWorkers),
		exitChan : make(chan error , allJobs),
	}
}

func (r *ReadJob)Filter() ([]string , error) {
	var err error
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")

	email := regEmail.FindAllString(string(<-r.Content),-1)
	if email == nil {
		err = errors.New("此文件中没有找到Email")
	}
	return email , err
}

func (r *ReadJob)ExitChan()chan error  {
	return r.exitChan
}