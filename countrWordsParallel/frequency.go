package main

import(
        "time"
        "flag"
        _ "io/ioutil"
        "os"
        "fmt"
        "bufio"
      )

func main() {
        flag.Parse()
        args := flag.Args() 
        
        length := len(args)
        c := make(chan int)
        
        start := time.Now()	
        for _, file := range args{
               go frequency(file, c)
        }

        count := 0
        for{
            a:= <-c
            count+=a
            if count == length{
                break
            }
        }

        elapsed := time.Since(start)
        fmt.Printf("elapsed time %s", elapsed)

        fmt.Println("End of main")
}

func frequency(fileName string, c chan int){ 
        file, err := os.Open(fileName)
        if err != nil {
               panic(err) 
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanWords)
       
        //var m map[string]int will make a variable but nil. 
        countMap := make(map[string]int)        

        for scanner.Scan(){
                s := scanner.Text()
                countMap[s]+= 1            
        }

       // for key, value := range countMap {
       //         fmt.Println("key:", key, " value:", value)
       // }

        c<-1

        fmt.Println("****end of frequency****")
}
