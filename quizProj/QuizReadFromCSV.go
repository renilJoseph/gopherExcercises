package main

import(	
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
	reader := csv.NewReader(bufio.NewReader(file))
	//fmt.Println(err)
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
	size := len(questionAnswerMap)

	fmt.Println("start the quiz : enter yes or no")
	newreader := bufio.NewReader(os.Stdin)
	text , _ := newreader.ReadString('\n')
	if text == "no\n"{
		fmt.Println("going out")
		return
	}
	fmt.Println(text, " woooo going into")
	scoreCounter := 0

	for i:=0;i < 2; i++{
		ran := rand.Intn(size)	
		ans := questionAnswerMap[keys[ran]]
		fmt.Println(keys[ran])
		text, _ := newreader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		//fmt.Println("answerr check", text, " and key ", ans)
		if ans == text{
			//fmt.Println()
			scoreCounter++	
		}
	}

	fmt.Println("score is:" , scoreCounter)
}


