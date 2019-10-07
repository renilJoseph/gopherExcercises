package main

import(
	"flag"
	"time"	
	"fmt"
	"encoding/csv"
	"log"
	"bufio"
	"io"
	"os"
	//"reflect"
	"math/rand"
	"strings"
)


func main(){
	
	file, err := os.Open("problems.csv")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	reader := csv.NewReader(bufio.NewReader(file))
	questionAnswerMap := make(map[string]string)
	keys := make([]string, 0, len(questionAnswerMap))
	for{
		line, err := reader.Read()
		if err == io.EOF{
			break
		}else if err != nil{
			log.Fatal(err)
		}	
	//	fmt.Println(reflect.TypeOf(line[0]), " ", reflect.TypeOf(line[1]))
		fmt.Println(line[0], " next", line[1])
		questionAnswerMap[line[0]]= line[1]
		keys = append(keys, line[0])
	}
	//create map and and array with questions

	scoreCounter := 0

	//PHASE 2 - adding timer to quiz		
	var inputPointer = flag.Int("sec", 10, "number of seconds for each user")
	timer1  := time.NewTimer(time.Duration(*inputPointer)*time.Second)	

	for{
		msg := quiz(questionAnswerMap, keys, scoreCounter, timer1)
		if msg == "exit"{
			return
		}
	}
}

func quiz(questionAnswerMap map[string]string, keys []string, scoreCounter int, timer1 *time.Timer) (string){
	
	outchan := make(chan string)
	
	size := len(questionAnswerMap)
	fmt.Println("start the quiz : enter yes or no")
	newreader := bufio.NewReader(os.Stdin)
	text , _ := newreader.ReadString('\n')
	if text == "no\n"{
		fmt.Println("going out")
		return "exit"
	}
	fmt.Println(text, " woooo going inside")

	var i =0 
	go check(outchan, timer1)

	for i=0;i < size; i++{
		select{
			case <-outchan:
				fmt.Println("TIMEOUT:: total score: ", scoreCounter, " and total questions answered:", i-1)
				return ""
			default:
				ran := rand.Intn(size)	
				ans := questionAnswerMap[keys[ran]]
				fmt.Println(keys[ran])	
				newreader.ReadString('\n')
				text = strings.TrimSuffix(text, "\n")
				if ans == text{
					//fmt.Println("***********")
					scoreCounter++
				}	
		}
	}

	fmt.Println("score is:" , scoreCounter)

	return ""
}

func check(outchan chan string, timer1 *time.Timer){
	
	<-timer1.C
	outchan <- "exit"	
}
