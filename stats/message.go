package stats

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func WriteMessage(c client.Client, db string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s", // message use s
	})
	if err != nil {
		return err
	}
	pt, err := client.NewPoint("message", tags, fields, t)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return c.Write(bp)
}
