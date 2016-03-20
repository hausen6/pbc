package command

import (
	"fmt"
	"testing"

	"github.com/codegangsta/cli"
)

func TestCmdNote(t *testing.T) {
	// Write your code here
	fmt.Println("starting TestCmdNote")
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "note",
			Action: CmdNote,
		},
	}
	app.Run([]string{"note", "note", "test", "this is test."})
}
