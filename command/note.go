package command

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/hhkbp2/go-logging"
	"github.com/xconstruct/go-pushbullet"
)

// logging format
var (
	log         = logging.GetLogger("pbc.note")
	handler     = logging.NewStdoutHandler()
	format      = "[%(levelname)s]:%(filename)s:%(lineno)d %(message)s"
	date_format = "%Y/%M%/D-%H:%m:%S"
	formatter   = logging.NewStandardFormatter(format, date_format)
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
		log.Infof("pushing note for %s.", device.Model)
		// err = pb.PushNote(device.Iden, "TEST", "this is test") if err != nil {
		// 	log.Error(err)
		// } else {
		// 	log.Infof("pushing note for %s.\n", device.Model)
		// }
	}

}
