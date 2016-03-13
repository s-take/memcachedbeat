package beater

import (
	"bufio"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"../config"
)

type Memcachedbeat struct {
	Configuration *config.Config
	period        time.Duration
	done          chan struct{}
}

// Creates beater
func New() *Memcachedbeat {
	return &Memcachedbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Memcachedbeat) Config(b *beat.Beat) error {

	err := cfgfile.Read(&bt.Configuration, "")
	if err != nil {
		logp.Err("Error reading config file: %v", err)
		return err
	}

	return nil
}

func (bt *Memcachedbeat) Setup(b *beat.Beat) error {
	// Setting default period if not set
	if bt.Configuration.Memcachedbeat.Period == "" {
		bt.Configuration.Memcachedbeat.Period = "10s"
	}

	var err error
	bt.period, err = time.ParseDuration(bt.Configuration.Memcachedbeat.Period)
	if err != nil {
		return err
	}
	return nil
}

func (bt *Memcachedbeat) Run(b *beat.Beat) error {

	logp.Info("memcachedbeat is running! Hit CTRL-C to stop it.")

	var err error
	var blanks = regexp.MustCompile(`\s+`)

	ticker := time.NewTicker(bt.period)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		cmd := exec.Command("memcached-tool", "localhost:11211", "stats")
		stdout, err := cmd.StdoutPipe()

		if err != nil {
			logp.Err("Error get stdout pipe: %v", err)
			return err
		}

		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		lines := 0
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
		}

		for scanner.Scan() {
			if lines != 0 {
				line := scanner.Text()
				items := blanks.Split(line, -1)
				event[items[1]] = toFloat(items[2])
			}
			lines += 1
		}

		b.Events.PublishEvent(event)
		logp.Info("Event sent")
	}

	return err
}

func toFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		logp.Err("Cannot parser to float. Ignore this value: %v", err)
		return 0
	}

	return value
}

func (bt *Memcachedbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Memcachedbeat) Stop() {
	close(bt.done)
}
