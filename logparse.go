package main

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)


type Event struct {
    Time        string
    Ip          string
    Description string
}


func Parse(content string) ([]Event, error) {
    // Use standard string manipulations to implement this function.
var Parsed []Event;
temp := strings.Split(content,"\n")  
errorMessage := ""
for i,element := range temp {
  // i is the index where we are
  // element is the element from someSlice for where we are
	if (len(element)>28) {
	spl:=strings.Index(element[27:len(element)-1]," ") // spl is the first space after the IP address
	 	if (spl!=-1) {
	 	Parsed = append(Parsed,Event{element[0:25],element[27:27+spl],element[28+spl:len(element)]})
		/* fmt.Println(element[0:25]) // Date is always the first 25 chars	
	 	fmt.Println(element[27:27+spl]) // IP
	 	fmt.Println(element[28+spl:len(element)]) */
	 	} else {
	 		errorMessage=errorMessage+"Error on line "+strconv.Itoa(i)+": Could not find IP address"+"\n"
	 	} 
	 } else {
	 	errorMessage=errorMessage+"Error on line "+strconv.Itoa(i)+": Line too short"+"\n"
	 }

}
if (errorMessage=="") {
	return Parsed, nil
	} else {
	return Parsed, errors.New(errorMessage)
}
}

/*
func MustParse(content string) []Event {
    // Instead of returning an error here, panic!
}

func ParseRegex(content string) ([]Event, error) {
    // Use regular expressions to implement this function.
}

func MustParseRegex(content string) []Event {
    // Instead of returning an error here, panic!
}
*/

func main() {

	// Sample correct log for parsing

	log := `Wed Feb 24 12:03:14 2016 0 94.156.237.146 41280 /home/site/public_html/log/fonts/glyphicons-halflings-regular.ttf b _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 4.156.237.146 23320 /home/site/public_html/log/fonts/glyphicons-halflings-regular.woff b _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 69727 /home/site/public_html/log/css/animate.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 99984 /home/site/public_html/log/css/bootstrap.min.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 127.0.0.1 33233 /home/site/public_html/log/css/font-awesome.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 32.156.237.146 27466 /home/site/public_html/log/css/font-awesome.min.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 11741 /home/site/public_html/log/css/main.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 5839 /home/site/public_html/log/css/responsive.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 3979 /home/site/public_html/log/css/responsive2.css a _ i r site ftp 1 * c`

	// Sample incorrect log for parsing

	log2:= `Wed Feb 24 12:03:14 2016 0 94.156.237.146 41280 /home/site/public_html/log/fonts/glyphicons-halflings-regular.ttf b _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 23320 /home/site/public_html/log/fonts/glyphicons-halflings-regular.woff b _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 69727 /home/site/public_html/log/css/animate.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 99984 /home/site/public_html/log/css/bootstrap.min.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 127.0.0.1 33233 /home/site/public_html/log/css/font-awesome.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 32.156.237.146 27466 /home/site/public_html/log/css/font-awesome.min.css a _ i r site ftp 1 * c

Error: Unable to connect

Wed Feb 24 12:03:14 2016 0 94.156.237.146 33233 /home/site/public_html/log/css/font-awesome.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 27466 /home/site/public_html/log/css/font-awesome.min.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 20164.156.237.146 11741 /home/site/public_html/log/css/main.css a _ i r site ftp 1 * c
Wed Feb 24 12:03:14 2016 0 94.156.237.146 5839 /home/site/public_html/log/css/responsive.css a _ i r site ftp 1 * c
[Finished in 0.2s with exit code 2]
Wed Feb 24 12:03:14 2016 0 94.156.237.146 3979 /home/site/public_html/log/css/responsive2.css a _ i r site ftp 1 * c`
 fmt.Println("Parsing Correct Log...")
	x, err :=Parse(log)
 fmt.Println("Finished Parsing.")
 if (err==nil) { 
 	fmt.Println("No errors.")
 } else {
 fmt.Println("Errors:\n",err)
 }	

 fmt.Println("\n\nParsing Incorrect Log...")	
	x, err = Parse(log2)
 fmt.Println("Finished Parsing.\n\n")
 if (err==nil) { 
 	fmt.Println("No errors.")
 } else {
 fmt.Println("Errors:\n",err)
 }		
 fmt.Println("Sample first row of Struct:")
 fmt.Println(x[1].Time+"\n"+x[1].Ip+"\n"+x[1].Description)
}