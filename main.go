package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mitchellh/colorstring"
	"github.com/phayes/permbits"
)

func main() {
		octal, err := strconv.ParseUint(os.Args[1], 8, 32)
		if err != nil {
				panic(err)
		}

		perms := permbits.PermissionBits(octal)
		if perms.Setuid() {
				colorstring.Print("[bold][red]setuid[reset] ")
		}
		if perms.Setgid() {
				colorstring.Print("[red]setgid[reset] ")
		}
		if perms.Sticky() {
				colorstring.Print("[bold][magenta]sticky[reset] ")
		}
		if perms.Setuid() || perms.Setgid() || perms.Sticky() {
				fmt.Println()
		}
		colorstring.Print("[bold][cyan]owner:[reset] ")
		if perms.UserRead() {
				colorstring.Print("[green]read[reset] ")
		} 
		if perms.UserWrite() {
				colorstring.Print("[green]write[reset] ")
		}
		if perms.UserExecute() {
				colorstring.Print("[green]execute[reset] ")
		}
		fmt.Println()
		colorstring.Print("[bold][magenta]group:[reset] ")
		if perms.GroupRead() {
				colorstring.Print("[green]read[reset] ")
		}
		if perms.GroupWrite() {
				colorstring.Print("[green]write[reset] ")
		}
		if perms.GroupExecute() {
				colorstring.Print("[green]execute[reset] ")
		}
		fmt.Println()
		colorstring.Print("[bold][yellow]other:[reset] ")
		if perms.OtherRead() {
				colorstring.Print("[green]read[reset] ")
		}
		if perms.OtherWrite() {
				colorstring.Print("[green]write[reset] ")
		}
		if perms.OtherExecute() {
				colorstring.Print("[green]execute[reset] ")
		}
		fmt.Println()
}
