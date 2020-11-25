package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/jlaffaye/ftp"
)

type Connection struct {
	Host     string
	Login    string
	Password string
	Root     string
	Output	 string
}

func main() {

	conn := Connection{
		Host:     "",
		Login:    "",
		Password: "",
		Root:     "/",
		Output:   "",
	}

	flag.StringVar(&conn.Host, "h", "", "host[:port]")
	flag.StringVar(&conn.Login, "u", "", "User")
	flag.StringVar(&conn.Password, "p", "", "Password")
	flag.StringVar(&conn.Root, "r", "/", "Root directory")
	flag.StringVar(&conn.Root, "o", "/", "Output file")
	flag.Parse()

	if len(conn.Host) == 0 || len(conn.Login) == 0 || len(conn.Password) == 0 {
		ProgramError("Missing arguments")
	}

	client, err := ftp.Dial(conn.Host)
	ErrorCheck(err)

	err = client.Login(conn.Login, conn.Password)
	ErrorCheck(err)

	fmt.Println("Connected to server")

	ReadDir(client, conn.Root)
}

func ReadDir(client *ftp.ServerConn, dir string) {
	entries, _ := client.List(dir)

	for _, entry := range entries {
		relativePath := strings.Trim(dir+"/"+entry.Name, "/")
		if entry.Type == ftp.EntryTypeFolder {
			ReadDir(client, relativePath)
		} else {
			fmt.Println(relativePath)
		}
	}
}

func ErrorCheck(err error) {
	if err != nil {
		color.Set(color.FgHiRed)
		panic(err)
		color.Unset()
	}
}

func ProgramError(message string) {
	panic(fmt.Sprintf(message))
}
