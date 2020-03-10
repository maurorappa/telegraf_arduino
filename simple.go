package simple

// simple.go

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
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

func (s *Simple) Gather(acc telegraf.Accumulator) error {
    for _,s := s.Input.Sensor() {
        acc.AddFields("state", map[string]interface{}{"value": s}, nil)
    }
    return nil
}

func init() {
    inputs.Add("simple", func() telegraf.Input { return &Sensor{} })
}
