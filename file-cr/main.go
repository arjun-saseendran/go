package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
createData(50.5)
var data  = readData()
fmt.Println(data)

}

func createData(num float64){
	numData := fmt.Sprint(num)
	os.WriteFile("text.txt",[]byte(numData),0644)
}

func readData()float64{
	data, _ := os.ReadFile("text.txt")
	stringData := string(data)
	floatData,_ :=strconv.ParseFloat(stringData,64)
	return floatData
	
}