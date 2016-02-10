package influx

import (
	"github.com/influxdata/telegraf"
)

type InfluxSerializer struct {
}

func (s *InfluxSerializer) Serialize(metric telegraf.Metric) (string, error) {
	return metric.String(), nil
}
