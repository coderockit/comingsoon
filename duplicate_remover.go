package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Alphbetic order 0-9 then a-z, no caps allowed
var emailList []string

//Struct for creating array from file
type arrayBuilder struct {
	fileContent []byte
	numEmails   int
	firstComma  int
	secondComma int
	email       string
	fileName    string
}

func (fileReader *arrayBuilder) readFile() {
	var err error
	fileReader.fileContent, err = ioutil.ReadFile(fileReader.fileName)
	if err != nil {
		fmt.Print(err)
	}
}

func (fileReader *arrayBuilder) makeArray() {
	for i := 0; i < fileReader.numEmails; i++ {
		fileReader.makeArrayChunk()
		emailList = append(emailList, fileReader.email)
	}
}
func (fileReader *arrayBuilder) countCommas() {
	for i := 0; i < len(fileReader.fileContent); i++ {
		if fileReader.fileContent[i] == ',' {
			fileReader.numEmails++
		}
	}
}
func (fileReader *arrayBuilder) makeArrayChunk() {
	//comma finder
	for i := fileReader.firstComma + 1; fileReader.fileContent[i] != ','; i++ {
		fileReader.secondComma = i + 1
	}
	//end comma finder
	var count int = 0
	var slice []byte
	for i := fileReader.firstComma + 1; i < fileReader.secondComma; i++ {
		//if statement for newlines
		if fileReader.fileContent[count+fileReader.firstComma+1] == 10 {
			count++
		} else {
			slice = append(slice, fileReader.fileContent[count+fileReader.firstComma+1])
			count++
		}
	}
	fileReader.email = string(slice[:])
	fileReader.firstComma = fileReader.secondComma
}

//Struct for sorting algorithm
type bubbleSort struct {
	whichCharS1 int
	whichCharS2 int
	stopSpot    int
	currentInd  int
	tempString  string
}

func (sorter bubbleSort) alphVal(str1, str2 string) int {
	if str1[sorter.whichCharS1] == 64 || str1[sorter.whichCharS1] == 46 {
		sorter.whichCharS1++
	}
	if str2[sorter.whichCharS2] == 64 || str2[sorter.whichCharS2] == 46 {
		sorter.whichCharS2++
	}
	if str1[sorter.whichCharS1] == str2[sorter.whichCharS2] && len(str1)-1 > sorter.whichCharS1 && len(str2)-1 > sorter.whichCharS2 {
		sorter.whichCharS1++
		sorter.whichCharS2++
		return (sorter.alphVal(str1, str2))
	} else if str2[sorter.whichCharS2] < str1[sorter.whichCharS1] {
		sorter.whichCharS1, sorter.whichCharS2 = 0, 0
		return (1)
	} else if str1[sorter.whichCharS1] < str2[sorter.whichCharS2] {
		sorter.whichCharS1, sorter.whichCharS2 = 0, 0
		return (2)
	}
	return (3)
	//value 3 return cases: (ison , isond) (identical string)
}
func (sorter bubbleSort) sortElem() {
	if sorter.stopSpot > -1 {
		for i := 0; i < sorter.stopSpot; i++ {
			var alphValRet = sorter.alphVal(emailList[sorter.currentInd], emailList[sorter.currentInd+1])
			if alphValRet == 1 {
				sorter.tempString = emailList[sorter.currentInd+1]
				emailList[sorter.currentInd+1] = emailList[sorter.currentInd]
				emailList[sorter.currentInd] = sorter.tempString
				sorter.currentInd++

			} else if alphValRet == 2 {
				sorter.currentInd++
			} else if alphValRet == 3 {
				sorter.currentInd++
			}
		}
		sorter.currentInd = 0
		sorter.stopSpot--
		sorter.sortElem()
	} else {

	}
}

//Struct for duplicate removal
//Removal cases (identical strings) (ipson, ipso)
type duplicateRemover struct {
	numEmails int
	whichChar int
	toRemove  []int
}

func (remover *duplicateRemover) worker() {
	for i := 0; i < remover.numEmails-1; i++ {
		if emailList[i] == emailList[i+1] {
			remover.toRemove = append(remover.toRemove, i+1)
		}
	}
	for i := len(remover.toRemove) - 1; i > -1; i-- {
		if len(emailList)-1 == remover.toRemove[i] {
			emailList = emailList[:len(emailList)-1]
		} else {
			remover.removeIndex(emailList, remover.toRemove[i])
		}
	}
}

func (remover *duplicateRemover) removeIndex(s []string, i int) {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	emailList = s[:len(s)-1]
}

//Struct for file writing
type fileWriter struct {
	fileName string
}

func (makeFile *fileWriter) myWriter() {
	f, err := os.Create(makeFile.fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < len(emailList); i++ {
		f.WriteString(",")
		f.WriteString(emailList[i])
	}
	f.WriteString(",")
}

func main() {
	fileReader := arrayBuilder{firstComma: 0, numEmails: -1, fileName: "emails.txt"}
	fileReader.readFile()
	fileReader.countCommas()
	fileReader.makeArray()
	sorter := bubbleSort{whichCharS1: 0, whichCharS2: 0, stopSpot: len(emailList) - 1, currentInd: 0, tempString: ""}
	sorter.sortElem()
	remover := duplicateRemover{whichChar: 0, numEmails: fileReader.numEmails}
	remover.worker()
	makeFile := fileWriter{fileName: "emails.txt"}
	makeFile.myWriter()
}
