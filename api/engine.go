package api

import (
	"errors"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

type EngineInterface interface {
	SetConfig(config *Config)
	SetActiveRoutine(configId string) error
	Init() error
	Run() error
	Stop() error
	Reset() error
}

type State struct {
	activeRoutine string
	running       bool
	tickers       []*time.Ticker
}

type DefaultEngine struct {
	config *Config
	state  *State
}

func New() *DefaultEngine {
	engine := new(DefaultEngine)
	engine.config = new(Config)
	engine.state = new(State)
	engine.state.running = false
	return engine
}

func (d *DefaultEngine) Run() error {

	if d.state.running {
		d.Stop()
		d.Reset()
	}

	//At this point it should create the active objects
	cfg := d.config.Routines[d.state.activeRoutine]
	for k, v := range cfg.Generators {
		//Set up the generators
		dtime := time.Duration(v.Delay)
		// Minus 1 here because the pulse is hardcoded to be 1ms long
		ticker := time.NewTicker(time.Millisecond * (dtime - 1))
		d.state.tickers = append(d.state.tickers, ticker)
		go func(cfg BooleanGenerator, id int) {
			///defer close(statusChannel)
			//Set up the pin output here
			var tickHandler func(cfg BooleanGenerator, id int, t time.Time)
			if engine.config.Debug {
				tickHandler = func(cfg BooleanGenerator, id int, t time.Time) {
					log.Println("Running Timer", id, "Active step:", cfg.ActiveStep, t)
				}
			} else {
				pin := rpio.Pin(cfg.Pin)
				pin.Output()
				tickHandler = func(cfg BooleanGenerator, id int, t time.Time) {
					pin.Write(rpio.High)
					time.Sleep(time.Millisecond * 1)
					pin.Write(rpio.Low)
				}
			}

			for t := range ticker.C {
				if cfg.ActiveStep < len(cfg.Pattern) {
					cfg.ActiveStep++
				} else {
					cfg.ActiveStep = 0
				}
				tickHandler(cfg, id, t)
			}
		}(v, k)
	}

	return nil
}

func (d *DefaultEngine) Stop() {
	for _, v := range d.state.tickers {
		v.Stop()
	}
}

func (d *DefaultEngine) Init() error {
	if d.config == nil {
		return errors.New("cannot initialize engine without config")
	}
	return nil
}

func (d *DefaultEngine) Reset() error {
	d.Stop()
	for _, v := range d.config.Routines[d.state.activeRoutine].Generators {
		v.ActiveStep = 0
	}
	return nil
}

func (d *DefaultEngine) SetConfig(config *Config) {
	d.config = config
	d.Init()
}

func (d *DefaultEngine) SetActiveRoutine(configId string) error {
	if _, ok := d.config.Routines[configId]; !ok {
		return errors.New("could not find routine")
	}

	d.state.activeRoutine = configId
	if d.state.running {
		d.Run()
	}

	return nil
}
