package command

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/hhkbp2/go-logging"
	"github.com/xconstruct/go-pushbullet"
)

// logging format
var (
	log        = logging.GetLogger("pbc.note")
	handler    = logging.NewStdoutHandler()
	format     = "[%(levelname)s]:%(filename)s:%(lineno)d %(message)s"
	dateFormat = "%Y/%M%/D-%H:%m:%S"
	formatter  = logging.NewStandardFormatter(format, dateFormat)
)

func getToken() string {
	token := os.Getenv("PBC_TOKEN")
	if token == "" {
		log.Fatal("PBC_TOKEN (environment variable) is not found.")
	}
	return token
}

// CmdNote is a function
func CmdNote(c *cli.Context) {
	// handler setting
	handler.SetFormatter(formatter)
	log.AddHandler(handler)
	log.SetLevel(logging.LevelDebug)

	// check the arguments
	args := c.Args()
	if len(args) < 2 {
		log.Fatal("number of arguments must be 2.")
		os.Exit(1)
	}

	// get token
	token := getToken()
	pb := pushbullet.New(token)

	// get devices
	devices, err := pb.Devices()
	if err != nil {
		log.Fatal(err)
	}
	// push note
	for _, device := range devices {
		err = pb.PushNote(device.Iden, args[0], args[1])
		if err != nil {
			log.Error(err)
		} else {
			log.Infof("pushing note for %s.", device.Model)
		}
	}

}
