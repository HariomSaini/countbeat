package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/HariomSaini/countbeat/config"
)

// Countbeat configuration.
type Countbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of countbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Countbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

// Run starts countbeat.
func (bt *Countbeat) Run(b *beat.Beat) error {
	logp.Info("countbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
			select {
			case <-bt.done:
					return nil
			case <-ticker.C:
			}

			event := common.MapStr{ 
					"@timestamp": common.Time(time.Now()), 
					"type":       b.Name,
					"counter":    counter,
			}
			bt.client.PublishEvent(event) 
			logp.Info("Event sent")
			counter++
	}
}}

// Stop stops countbeat.
func (bt *Countbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
