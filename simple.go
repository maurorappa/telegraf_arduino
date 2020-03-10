package simple

// simple.go

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
    "math/rand"
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
    for _,s := range s.Input {
        acc.AddFields(s, map[string]interface{}{"value": rand.Intn(100)}, nil)
    }
    return nil
}

func init() {
    rand.Seed(80)
    inputs.Add("simple", func() telegraf.Input { return &Sensor{} })
}
