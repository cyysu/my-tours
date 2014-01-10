// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
        "bytes"
        "encoding/json"
        "go/ast"
        "go/parser"
        "go/printer"
        "go/token"
        "net/http"
        "os/exec"
        "os"
        "io"
	"time"
	"errors"
)

const (
	TYPE_PYTHON int = 0
	TYPE_RUBY int = 1
	TYPE_JAVA int = 2
	TYPE_C int = 3
)

func init() {
        http.HandleFunc("/fmt", fmtHandler)
        http.HandleFunc("/fmt_ruby", fmtRubyHandler)
        http.HandleFunc("/fmt_java", fmtJavaHandler)
        http.HandleFunc("/fmt_c", fmtCHandler)
}

type fmtResponse struct {
        Body  string
        Error string
}

func fmtHandler(w http.ResponseWriter, r *http.Request) {
        resp := new(fmtResponse)
	body, err := execute_code(TYPE_PYTHON, string(r.FormValue("body")[:]))
        if err != nil {
                resp.Error = string(body[:])
        } else {
		resp.Body = string(body[:])
        }
        json.NewEncoder(w).Encode(resp)
}

func fmtRubyHandler(w http.ResponseWriter, r *http.Request) {
        resp := new(fmtResponse)
	body, err := execute_code(TYPE_RUBY, string(r.FormValue("body")[:]))
        if err != nil {
                resp.Error = string(body[:])
        } else {
		resp.Body = string(body[:])
        }
        json.NewEncoder(w).Encode(resp)
}

func fmtJavaHandler(w http.ResponseWriter, r *http.Request) {
        resp := new(fmtResponse)
	body, err := execute_code(TYPE_JAVA, string(r.FormValue("body")[:]))
        if err != nil {
                resp.Error = string(body[:])
        } else {
		resp.Body = string(body[:])
        }
        json.NewEncoder(w).Encode(resp)
}

func fmtCHandler(w http.ResponseWriter, r *http.Request) {
        resp := new(fmtResponse)
	body, err := execute_code(TYPE_C, string(r.FormValue("body")[:]))
        if err != nil {
                resp.Error = string(body[:])
        } else {
		resp.Body = string(body[:])
        }
        json.NewEncoder(w).Encode(resp)
}

func gofmt(body string) (string, error) {
        fset := token.NewFileSet()
        f, err := parser.ParseFile(fset, "prog.go", body, parser.ParseComments)
        if err != nil {
                return "", err
        }
        ast.SortImports(fset, f)
        var buf bytes.Buffer
        config := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
        err = config.Fprint(&buf, fset, f)
        if err != nil {
                return "", err
        }
        return buf.String(), nil
}

func execute_code(lang int, pyStr string) ([]byte, error)  {
        //Write python program to a file
	var ext string
	if lang == TYPE_PYTHON {
		ext = ".py"
	} else if lang == TYPE_RUBY {
		ext = ".rb"
	} else if lang == TYPE_C {
		ext = ".c"
	} else {
	  	ext = ".java"
	}
        file, _ := os.Create("test" + ext)
        io.WriteString(file, pyStr)

	var out []byte
	var err error
	
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3e9)
		timeout <- true
	}()
	
	ch := make(chan bool, 1)
	var cmd *exec.Cmd
	go func() {
		//Execute the program
		if lang == TYPE_PYTHON {
			cmd = exec.Command("python", "test" + ext)
		} else if lang == TYPE_RUBY {
			cmd = exec.Command("ruby", "test.rb")                        
		} else if lang == TYPE_C {
			cmd = exec.Command("gcc", "test" + ext)
			out, err = cmd.CombinedOutput()
			cmd = exec.Command("./a.out")
		} else {
			cmd = exec.Command("javac", "test" + ext)
			out, err = cmd.CombinedOutput()
			cmd = exec.Command("java", "test")
		}
		out, err = cmd.CombinedOutput()
		ch <- true
	}()
	
	select {
	case <-ch:
		// Command finished in time.
		return out, err
	case <-timeout:
		// Stop forever loop command.
		cmd.Process.Kill()
		return []byte("人家被你玩坏了啦"), errors.New("人家被你玩坏了啦")
	}
}
