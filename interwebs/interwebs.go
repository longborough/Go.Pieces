// interwebs makes a basic web server
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Well hello, world!\n")
}

func DataServer(w http.ResponseWriter, req *http.Request) {
	dat, err := ioutil.ReadFile("D:/BLiss/Turkish.Airlines.2012/Load.History/LMOD.History/critical.lmod.log")
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
	} else {
		io.WriteString(w, fmt.Sprintf("%s\n",dat))
	}
}

func ProgServer(w http.ResponseWriter, req *http.Request) {
	path := strings.ToUpper(req.URL.Path[len("/prog/"):])
	dat, err := ioutil.ReadFile("D:/BLiss/Datalex/Lan.Backup/Troya.Endevor.Mirror/" + path)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Oops! Couldn't find %s\n",path))
	} else {
		io.WriteString(w, fmt.Sprintf("%s\n",dat))
	}
}

func XrefServer(w http.ResponseWriter, req *http.Request) {
	var found bool = false
	var output string = ""
	var line string = ""
	var star byte = '*' 
	var i int
	var x string
	var lastx string = ""
	xref := strings.ToUpper(req.URL.Path[len("/xref/"):])
	list := regexp.MustCompile("[ ,]+").Split(xref,-1)
	if len(list) < 1 {
		io.WriteString(w, "Need one or more program names after xref. Try /xref/uio3\n")
	} else {
		for i, x = range list {
			if len(x) > 1 && x[len(x)-1] == star {
				list[i] = x[:len(x)-1]
				lastx = list[i]
			} else if len(x) >= 4 {
				lastx = x[:4]
			} else {
				list[i] = x 					
				if len(x) < len(lastx) {
					list[i] = lastx[:len(lastx)-len(x)] + x 
				}
				lastx = list[i]
			}
		}
		data, err := ioutil.ReadFile("D:/BLiss/Turkish.Airlines.2012/Troya.Xref/troya.refs")
		table := strings.Split(fmt.Sprintf("%s",data),"\n")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
		} else {
			var last string = ""
			for _, titem := range table {
				if len(titem) >8 {
					for _, xitem := range list {
						if xitem == titem[:len(xitem)] || xitem == titem[5:5+len(xitem)] {
							found = true
							buildxref(titem,&last,&line,&output)
						}
					}
				} 
			} 
		}
		if found {
			addout(&output,&line)
			io.WriteString(w, output)
		} else {
			io.WriteString(w, fmt.Sprintf("Sorry, nothing found:\n Query: %s\n",list))
		}
	}
}

func buildxref(xitem string, last *string, line *string, output *string) {
	to := xitem[:4]
	from := xitem[5:9]
	if to != *last {
		addout(output,line)
		*line = to + " <== "
		*last = to
	}
	*line += from + " "
	if len(*line) > 80 {
		addout(output,line)
		*line = "         "
	}
}

func addout(output *string, line *string) {
	if len(*line) > 5 {
		*output = *output + "\n" + *line
	}
}

func ExitServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Farewell, cruel world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/data", DataServer)
	http.HandleFunc("/prog/", ProgServer)
	http.HandleFunc("/xref/", XrefServer)
	http.HandleFunc("/bye", ExitServer)
	log.Fatal(http.ListenAndServe(":11080", nil))
}
