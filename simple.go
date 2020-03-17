package simple

// simple.go

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/tarm/serial"
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
	for _, s := range s.Input {
		acc.AddFields(s, map[string]interface{}{"value": 1}, nil)
	}
	return nil
}

func init() {
	inputs.Add("simple", func() telegraf.Input { return &Sensor{} })
	c := &serial.Config{Name: "/dev/tty.usbmodem14101", Baud: 9600, ReadTimeout: 1000}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("cannot open serial: %s\n", err)
		return "null"
	}
}
