package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	//input and output flags

	//here we find link of directory with in.txt
	src := flag.String("src", "", "in.txt location")
	//here we choose a link of output directory
	out := flag.String("out", "", "folder out location")
	flag.Parse()

	err := parserURL(src, out)
	if err != nil {
		fmt.Println("Can't open file", err)
		os.Exit(1)
	}

}

//parsing with flags use
func parserURL(src *string, out *string) error {

	file, err := ioutil.ReadFile(*src) //choosing a text file for reading
	if err != nil {
		return err
	}

	getLocation := *out //choosing directory for file recording
	if getLocation == "" {
		fmt.Println("Can't choose a directory")
	}
	myLinks := string(file) //transfer text file from []byte to string
	if myLinks == "" {
		fmt.Println("Can't create myLinks")
	}

	parserErr := parser(myLinks, getLocation)
	if parserErr != nil {
		fmt.Println("Can't open file", parserErr)
		os.Exit(2)
	}

	return nil

}

// Text parsing
func parser(myLinks string, getLocation string) error {
	//split strings to []string
	link := strings.Split(myLinks, "\r\n")
	if len(link) == 0 {
		fmt.Println("Can't split file")
	}

	for i := range link {
		//link for the output text file
		url := link[i]

		err := parseText(url, getLocation)
		if err != nil {
			fmt.Println("Can't open parseText function", err)
			os.Exit(2)
		}

	}
	return nil
}

//HTTP Get, recording and writing output file. We will get html-content
func parseText(url string, getLocation string) error {

	//name of file, domain.name without "https://www."
	fileName := url

	//HTTP request data processing to string
	resp, err := http.Get("https://www." + url)
	if err != nil {
		fmt.Println("Can't get URLs", err)
		os.Exit(3)
	}
	body, err := ioutil.ReadAll(resp.Body) //transfer from *Response to []byte
	if err != nil {
		fmt.Println("Can't read responses")
	}
	bodyString := string(body) //transfer from []byte to string
	if bodyString == "" {
		fmt.Println("can't provide bodystring")
	}
	//creating and recording files

	errWrite := writeFiles(fileName, bodyString, getLocation)
	if err != nil {
		fmt.Println("Can't write files", errWrite)
		os.Exit(5)
	}
	return nil

}

//creating output
func writeFiles(fileName string, bodyString string, getLocation string) error {
	//making output files
	newFiles, err := os.Create(getLocation + "/" + fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(4)
	}
	defer newFiles.Close()
	// recording files
	newFiles.WriteString(bodyString)

	return nil
}
