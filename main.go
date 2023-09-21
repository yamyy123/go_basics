package main
// import (
//     "fmt"
//     "time"
// )

// func main() {
//     go func() {
//         time.Sleep(2 * time.Second)
//         fmt.Println("Hello from Goroutine")
//     }()
//     time.Sleep(1 * time.Second)
//     fmt.Println("Hello from Main")
// }

/*  
    a.) Hello from Goroutine
        Hello from Main

     b.)Hello from Main

     c.)No Output

     d.)Error
*/

//<------------------------------------------------------------------------------------------------------------------------------------------------------->




// import (
//     "fmt"
//     "sync"
// )

// func workerA(wg *sync.WaitGroup) {  
//     defer wg.Done()   
//     fmt.Println("Worker A is working")
// }

// func main() {
//     var wg sync.WaitGroup
//     wg.Add(2)
//     go workerA(&wg)
// 	//go workerA(&wg)
//     wg.Wait()
//     fmt.Println("Main function has completed")
// }

   /* 
      a.) Worker A is working and main function will wait for infinite time

      b.) Worker A is Working
          Main function has completed 

	  c.) No Output


	  d.) Error
 
   */


// <------------------------------------------------------------------------------------------------------------------------------------------->



// import (
//     "fmt"
//     "sync"
// )

// func workerA(wg *sync.WaitGroup) {
//     defer wg.Done()
    
//     fmt.Println("A is working")
//     //wg.Add(1)
//     go workerB(wg)
	
// }

// func workerB(wg *sync.WaitGroup) {
//     defer wg.Done()
//     fmt.Println("Worker B is working")
//     fmt.Println("Worker B is completed")
// }

// func main() {
//     var wg sync.WaitGroup

//     wg.Add(1)

//     go workerA(&wg)

//     wg.Wait()

//     fmt.Println("Main function has completed")
// }


/*
a.)No Output
b.)A is working
   Main function has completed
c.)Error
d.)A is working
   Worker B is working
   Worker B is completed
   Main function has completed
*/


//<------------------------------------------------------------------------------------------------------------------------------------------->

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int) 
    
    go func() {
        fmt.Println("Sending data into the channel")
        ch <-1
        ch <-2
        fmt.Println("Data sent")
    }()
    time.Sleep(1 * time.Second)
    <-ch
	<-ch
    fmt.Println("Main function has completed")
}

