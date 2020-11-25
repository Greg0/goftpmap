module ftpmap

go 1.15

require (
	ftp v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.10.0
	github.com/jlaffaye/ftp v0.0.0-20201112195030-9aae4d151126
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	golang.org/x/text v0.3.4
)

replace ftp => ../ftp
