// interwebs makes a basic web server
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Well hello, world!\n")
}

func GrepServer(w http.ResponseWriter, req *http.Request) {
	rex := req.URL.Path[len("/grep/"):]
	if len(rex) < 2 {
		io.WriteString(w, fmt.Sprintf("Need at least 2 characters to search\n"))
	} else {
		cmd:= exec.Command("grep", "-ERia", "/(?=^[^*]).*" + + "/", "*.EASM")
		cmd.Dir = "/data/brentl/Troya.Endevor"
		out, err := cmd.Output()
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
		} else {
			io.WriteString(w, fmt.Sprintf("%s\n",out))
		}
	}
}

func LmodServer(w http.ResponseWriter, req *http.Request) {
	cmd:= exec.Command("git", "diff", "-U9999", "HEAD^")
	cmd.Dir = "/data/brentl/LMOD.History"
	out, err := cmd.Output()
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
	} else {
		io.WriteString(w, fmt.Sprintf("%s\n",out))
	}
}

func ProgServer(w http.ResponseWriter, req *http.Request) {
	path := strings.ToUpper(req.URL.Path[len("/prog/"):])
	dat, err := ioutil.ReadFile("/data/brentl/Troya.Endevor/" + path)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Oops! Couldn't find %s\n",path))
	} else {
		io.WriteString(w, fmt.Sprintf("%s\n",dat))
	}
}

func ProcServer(w http.ResponseWriter, req *http.Request) {
	path := strings.ToUpper(req.URL.Path[len("/proc/"):])
	if path == "PROD" {
		path = "/PLEX.PROCLIB/ALCSPROD"
	} else {
		path = "/THY14.PROCLIB/ALCS" + path
	}
	dat, err := ioutil.ReadFile("/data/brentl/THY.Support.Code" + path)
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
		data, err := ioutil.ReadFile("/data/brentl/Troya.Xref/troya.refs")
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
							buildxref(titem[:4],titem[5:9],&last,&line,&output,9)
						}
					}
				} 
			} 
		}
		if found {
			addout(&output,&line,9)
			io.WriteString(w, output)
		} else {
			io.WriteString(w, fmt.Sprintf("Sorry, nothing found:\n Query: %s\n",list))
		}
	}
}

func SerrServer(w http.ResponseWriter, req *http.Request) {
	var found bool = false
	var output string = ""
	var line string = ""
	var star byte = '*' 
	var i int
	var x string
	var lastx string = ""
	xref := strings.ToUpper(req.URL.Path[len("/serrc/"):])
	list := regexp.MustCompile("[ ,]+").Split(xref,-1)
	if len(list) < 1 {
		io.WriteString(w, "Need one or more program names after xref. Try /xref/uio3\n")
	} else {
		for i, x = range list {
			if len(x) > 1 && x[len(x)-1] == star {
				list[i] = x[:len(x)-1]
				lastx = list[i]
			} else if len(x) >= 6 {
				lastx = x[:6]
			} else {
				list[i] = x 					
				if len(x) < len(lastx) {
					list[i] = lastx[:len(lastx)-len(x)] + x 
				}
				lastx = list[i]
			}
		}
		data, err := ioutil.ReadFile("/data/brentl/Troya.Xref/troya.serrc")
		table := strings.Split(fmt.Sprintf("%s",data),"\n")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Oops! %s\n",err))
		} else {
			var last string = ""
			for _, titem := range table {
				if len(titem) > 8 {
					for _, xitem := range list {
						if xitem == titem[:len(xitem)] || xitem == titem[5:5+len(xitem)] {
							found = true
							buildxref(titem[:8],titem[9:13],&last,&line,&output,13)
						}
					}
				} 
			} 
		}
		if found {
			addout(&output,&line,13)
			io.WriteString(w, output)
		} else {
			io.WriteString(w, fmt.Sprintf("Sorry, nothing found:\n Query: %s\n",list))
		}
	}
}

func buildxref(to string, from string, last *string, line *string, output *string, minlen int) {
	if to != *last {
		addout(output, line, minlen)
		*line = to + " <== "
		*last = to
	}
	*line += from + " "
	if len(*line) > 80 {
		addout(output,line,minlen)
		*line = "     " + "             "[:len(to)]
	}
}

func addout(output *string, line *string, minlen int) {
	if len(*line) > minlen {
		*output = *output + "\n" + *line
	}
}

func ExitServer(w http.ResponseWriter, req *http.Request) {
	if req.Host != "127.0.0.1:11080" {
		io.WriteString(w, fmt.Sprintf("Sorry, nice try though\n"))
	} else {
		pid := os.Getpid()
		myself, _ := os.FindProcess(pid)
		_ = myself.Kill()
	}
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/lmod", LmodServer)
	http.HandleFunc("/proc/", ProcServer)
	http.HandleFunc("/prog/", ProgServer)
	http.HandleFunc("/xref/", XrefServer)
	http.HandleFunc("/serrc/", SerrServer)
	http.HandleFunc("/shutdown/now", ExitServer)
	log.Fatal(http.ListenAndServe(":11080", nil))
}
