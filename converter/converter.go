package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var inDirFlag *string = flag.String("i", "", "input directory to check")
var outDirFlag *string = flag.String("o", "", "output directory")

func main() {
	fmt.Println("converter started")
	flag.Parse() // Scan the arguments list

	if *inDirFlag == "" || *outDirFlag == "" {
		fmt.Println("Please provide input directory AND output directory")
	} else {
		fmt.Println("input directory: ", *inDirFlag)
		fmt.Println("output directory: ", *outDirFlag)

		ProcessData(*inDirFlag, *outDirFlag)
	}
	fmt.Println("converter finished")
}

func ProcessData(inDir string, outDir string) {

	for {
		files, err := ioutil.ReadDir(inDir)
		if err != nil {
			fmt.Println(err)
		} else if len(files) > 0 {

			for _, file := range files {
				oldFqfn := inDir + string(os.PathSeparator) + file.Name()
				newFqfn := outDir + string(os.PathSeparator) + file.Name()
				fmt.Println("move ", oldFqfn, " -> ", newFqfn)
				err2 := os.Rename(oldFqfn, newFqfn)
				if err2 != nil {
					fmt.Println(err)
				}
			}
		}
		fmt.Println("sleep for 5 seconds")
		time.Sleep(5 * time.Second)
	}
}
