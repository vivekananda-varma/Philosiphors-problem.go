package main

import (
	"fmt"
	"sync"
)

var forks [6]sync.Mutex
var exit int
var wg sync.WaitGroup

func philosiphor(id int) {
	if id%2==0{
    forks[id].Lock()
    fmt.Printf("philosiphor[%v] picked up the right fork\n", id)
    forks[(id+1)%6].Lock()
		fmt.Printf("philosiphor[%v] picked up the left fork\n", id)
  } else{
    forks[(id+1)%6].Lock()
      fmt.Printf("philosiphor[%v] picked up the left fork\n", id)
      forks[id].Lock()
			fmt.Printf("philosiphor[%v] picked up the right fork\n", id)
  }
	fmt.Printf("philosiphor[%v] ate the food successfully\n", id)
	forks[id].Unlock()
	forks[(id+1)%6].Unlock()
	exit++
	wg.Done()

}

func main() {
	fmt.Println("Dining philosiphors problem")

	for i := 0; i < 6; i++ {

		wg.Add(1)
		go philosiphor(i)
		wg.Wait()
	}
	for exit != 6{

	}

}
