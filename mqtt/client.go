package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	defaultClientId      = "go-bambulab-mqtt"
	defaultLocalUsername = "bblp"
	defaultPort          = 8883
	defaultProtocol      = "ssl"
	defaultQos           = 1
)

type Client struct {
	mqtt     mqtt.Client
	local    bool
	deviceId string
	msgs     chan<- Message
}

func newClient(url, deviceId, username, password, clientId string, local bool) (*Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(clientId)

	opts.SetOrderMatters(false)       // Allow out of order messages (use this option unless in order delivery is essential)
	opts.ConnectTimeout = time.Second // Minimal delays on connect
	opts.WriteTimeout = time.Second   // Minimal delays on writes
	opts.KeepAlive = 10               // Keepalive every 10 seconds so we quickly detect network outages
	opts.PingTimeout = time.Second    // local broker so response should be quick

	// Automate connection management (will keep trying to connect and will reconnect if network drops)
	opts.ConnectRetry = true
	opts.AutoReconnect = true
	opts.SetUsername(username)
	opts.SetPassword(password)

	// Log events
	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		fmt.Println("connection lost")
	}
	opts.OnConnect = func(mqtt.Client) {
		fmt.Println("connection established")
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		fmt.Println("attempting to reconnect")
	}

	c := mqtt.NewClient(opts)

	client := &Client{
		mqtt:     c,
		local:    local,
		deviceId: deviceId,
	}
	return client, nil
}

// NewLocalClient creates a new client connecting to local printer mqtt server
func NewLocalClient(ip, deviceId, accessCode string) (*Client, error) {
	url := fmt.Sprintf("%s://%s:%d", defaultProtocol, ip, defaultPort)
	return newClient(url, deviceId, defaultLocalUsername, accessCode, defaultClientId, true)
}

// NewCloudClient creates a new client connecting to bambulab cloud mqtt server
func NewCloudClient(endpoint, deviceId, username, password string) (*Client, error) {
	url := fmt.Sprintf("%s://%s:%d", defaultProtocol, endpoint, defaultPort)
	return newClient(url, deviceId, username, password, defaultClientId, false)
}

func (c *Client) Connect() error {
	fmt.Println("connecting...")
	if token := c.mqtt.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Disconnect() {
	c.mqtt.Disconnect(1000)
}

func (c *Client) Subscribe(msgs chan<- Message) {
	topic := c.reportTopic()
	fmt.Printf("subscribe %s\n", topic)
	c.msgs = msgs
	c.mqtt.Subscribe(topic, 1, c.handle)
}

func (c *Client) handle(_ mqtt.Client, msg mqtt.Message) {
	var m Message
	if err := json.Unmarshal(msg.Payload(), &m); err != nil {
		fmt.Printf("fail parse msg, err=%s, msg=%s\n", err, msg.Payload())
		return
	}
	if c.msgs == nil {
		fmt.Printf("fail handle msg, chan nil, msg=%s\n", msg.Payload())
		return
	}
	c.msgs <- m
}

func (c *Client) publish(ctx context.Context, msg []byte) error {
	topic := c.requestTopic()
	t := c.mqtt.Publish(topic, defaultQos, false, msg)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.Done():
		return t.Error()
	}
}

func (c *Client) requestTopic() string {
	return fmt.Sprintf("device/%s/request", c.deviceId)
}

func (c *Client) reportTopic() string {
	return fmt.Sprintf("device/%s/report", c.deviceId)
}
