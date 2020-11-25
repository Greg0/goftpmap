package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jlaffaye/ftp"
)

type Connection struct {
	Host     string
	Login    string
	Password string
	Root     string
	Output   string
}

func main() {

	conn := Connection{
		Host:     "",
		Login:    "",
		Password: "",
		Root:     "/",
		Output:   "output.csv",
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

	client, err := ftp.Dial(conn.Host, ftp.DialWithDisabledUTF8(true))
	ErrorCheck(err)

	err = client.Login(conn.Login, conn.Password)
	ErrorCheck(err)

	fmt.Println("Connected to server")

	file, err := os.Create(conn.Output)
	ErrorCheck(err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"dir", "filename", "path"})
	count := ReadDir(client, conn.Root, writer)

	fmt.Println("Saved " + strconv.Itoa(count) + " to " + conn.Output)
}

func ReadDir(client *ftp.ServerConn, dir string, writer *csv.Writer) int {
	count := 0
	entries, _ := client.List(dir)
	for _, entry := range entries {
		relativePath := strings.Trim(dir+"/"+entry.Name, "/")
		if entry.Type == ftp.EntryTypeFolder {
			count += ReadDir(client, relativePath, writer)
		} else {
			writer.Write([]string{dir, entry.Name, relativePath})
			// fmt.Println(relativePath)
			count++
		}
	}

	return count
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
