package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
)

type PushingData struct {
	Pushing struct {
		SequenceID string `json:"sequence_id"`
		Command    string `json:"command"`
		Version    int    `json:"version"`
		PushTarget int    `json:"push_target"`
	} `json:"pushing"`
}

func newPushingData() PushingData {
	p := PushingData{}
	p.Pushing.SequenceID = "0"
	p.Pushing.Command = "pushall"
	p.Pushing.Version = 1
	p.Pushing.PushTarget = 1
	return p
}

// Send pushing.pushall request to broker
func (c *Client) PublishPushAll(ctx context.Context) error {
	data := newPushingData()
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Printf("mqtt client published, cmd=%s\n", "pushall")
	return c.publish(ctx, b)
}
