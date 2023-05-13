package main

import "fmt"


func Worker(id int, jobs <-chan int, results chan<- int){
	for job := range jobs{
		fmt.Printf("Worker %d started fibonacci for %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Workder %d ended fibonacci for %d with result: %d\n", id, job, fib)
		results <- fib
	}
}
func Fibonacci(n int) int{
	if n <= 2{
		return 1
	}
	return Fibonacci(n - 1) + Fibonacci(n - 2)
}
func WorkerPool(){
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks{
		jobs <- value
	}
	close(jobs)
	for i:= 0; i < len(tasks); i++{
		<- results
	}
}