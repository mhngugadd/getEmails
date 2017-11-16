package file

import (
	"path/filepath"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"github.com/mhngugadd/getEmail/worker"

)

type Jobs struct {
	Jobs chan worker.Job
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetAllFile(dirName string) ([]string, error ){
	fileName := []string{}
	list , err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _,file := range list  {
		isTxt := strings.HasSuffix(file.Name(),"txt")
		if ok := file.IsDir(); !ok && isTxt {
			fileName = append(fileName,file.Name())
		}
	}
	return  fileName, err
}

func ReadFileContent(files []string , job *ReadJob) *ReadJob {
	for _ , file := range files {
		content , err  := ioutil.ReadFile(file)
		if err != nil {
			job.exitChan <- err
		}
		job.Content <-content
	}
	return job
}