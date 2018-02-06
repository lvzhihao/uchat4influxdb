package stats

import (
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

func WriteMemberJoin(c client.Client, db string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s", // message use s
	})
	if err != nil {
		return err
	}
	pt, err := client.NewPoint("member.join", tags, fields, t)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return c.Write(bp)
}

func WriteMemberQuit(c client.Client, db string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s", // message use s
	})
	if err != nil {
		return err
	}
	pt, err := client.NewPoint("member.quit", tags, fields, t)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return c.Write(bp)
}
