package worker

type Job interface{
	Filter() ([]string ,error)
	ExitChan() chan error
}

func Worker(jobs chan Job ,result chan []string)  {
	j := <-jobs
	emails , err  := j.Filter()

	if err != nil {
		j.ExitChan() <- err
	}
	result <- emails
}

