package simple

// simple.go

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/tarm/serial"
	"log"
)

type Sensor struct {
	Input []string `toml:"sensors"`
}

func (s *Sensor) Description() string {
	return "Arduino plugin"
}

func (s *Sensor) SampleConfig() string {
	return `
  ## an array of sensors
  sensors = ["T","H"]
`
}

func (s *Sensor) Init() error {
	return nil
}

func (s *Sensor) Gather(acc telegraf.Accumulator) error {
	c := &serial.Config{Name: "/dev/tty.usbmodem14101", Baud: 9600, ReadTimeout: 1000}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("cannot open serial: %s\n", err)
		return "null"
	}
	for _, s := range s.Input {
		_, err = s.Write([]byte(sensor))
		if err != nil {
			log.Printf("cannot write to serial: %s\n", err)
		}
		nbytes, failed := s.Read(buf)
		whole_reply := string(buf)
		log.Printf("Got %d bytes:, took %f", nbytes, whole_reply)
		acc.AddFields(s, map[string]interface{}{"value": 1}, nil)
	}
	return nil
}

func init() {
	inputs.Add("simple", func() telegraf.Input { return &Sensor{} })
}
