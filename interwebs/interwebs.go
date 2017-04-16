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

func XrefServer(w http.ResponseWriter, req *http.Request) {
	var found bool = false
	var output string = ""
	var line string = ""
	xref := strings.ToUpper(req.URL.Path[len("/xref/"):])
	list := regexp.MustCompile("[ ,]+").Split(xref,-1)
	if len(list) < 1 {
		io.WriteString(w, "Need one or more program names after xref. Try /xref/uio3\n")
	} else {
		data, err := ioutil.ReadFile("D:/BLiss/Turkish.Airlines.2012/Troya.Xref/troya.refs")
		table := strings.Split(fmt.Sprintf("%s",data),"\n")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
		} else {
			var last string = ""
			for _, titem := range table {
				if len(titem) >8 {
					for _, xitem := range list {
						if xitem == titem[:4] || xitem == titem[5:9] {
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
	http.HandleFunc("/xref/", XrefServer)
	http.HandleFunc("/bye", ExitServer)
	log.Fatal(http.ListenAndServe(":11080", nil))
}
