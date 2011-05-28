package tripit

import (
	"json"
	"testing"
	"strings"
	"log"
	"os"
	"fmt"
	"time"
)

func TestWarning(t *testing.T) {
	x := func() os.Error { return &Warning{"Something went wrong", "trip", "2011-05-27T13:38:33"} }()
	log.Print(x)
}

func TestError(t *testing.T) {
	x := func() os.Error { return &Error{500, nil, "Something else went wrong", "trip", "2011-05-27T13:38:34"} }()
	log.Print(x)
}

func TestJsonWrite(t *testing.T) {
	var r Response
	log.Print("Marshal JSON")
	r.Error = make([]Error, 1)
	r.Error[0].Code = 304
	r.Error[0].Description = "WTF"
	b, _ := json.Marshal(r)
	os.Stdout.Write(b)
	fmt.Fprintf(os.Stdout, "\n")
}

func TestJsonRead(t *testing.T) {
	s := `
{
"Error": [
	{
		"code":502,
		"description":"wtf",
		"timestamp":"2011-05-26T23:44:33"
	},
	{
		"code":503,
		"description":"WTF Happened?",
		"timestamp":"2011-05-26T23:44:34"
	}
]
}`
	log.Print("Unmarshal JSON")
	var r Response
	b := make([]uint8, 300)
	l, err := strings.NewReader(s).Read(b)
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(b[0:l], &r)
	if err != nil {
		log.Print(err)
	}
	log.Print(r)
	log.Print("Marshal it back")
	b2, _ := json.Marshal(r)
	os.Stdout.Write(b2)
	fmt.Fprintf(os.Stdout, "\n")
}

func TestDateTime(t *testing.T) {

	d := &DateTime{"2009-11-10", "14:00:00", "America/Los_Angeles", "-08:00"}
	s, err := d.Time()
	log.Print("Parsed time: ", s, " err: ", err)

	log.Print(time.LocalTime().Format(time.RFC3339))
	d.SetTime(time.LocalTime())

	log.Print("Assigned time: ", d)
}
