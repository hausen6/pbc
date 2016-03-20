package command

import (
	"fmt"
	"testing"

	"github.com/codegangsta/cli"
)

func TestCmdNote(t *testing.T) {
	// Write your code here
	fmt.Println("starting TestCmdNote")
	var c *cli.Context
	CmdNote(c)
}
