package api

import (
	"bufio"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	"strings"
)

func RunLoop(cfg *Config) {
	engine.SetConfig(cfg)
	log.SetFlags(0)
	log.Println("--- Pulse Trainer Interactive Console ---")
	consolein := bufio.NewReader(os.Stdin)
	originalArgs := os.Args
	if !cfg.Debug {
		err := rpio.Open()
		if err != nil {
			log.Fatal(err)
		}
	}
	for {
		fmt.Print("> ")
		input, err := consolein.ReadString('\n')
		handleError(err)
		cmdline := strings.ToLower(input)
		words := strings.Fields(cmdline)
		if len(words) == 0 {
			//if there is no entry do nothing, normally passing nothing to a command prints the help
			//but this is not helpful in an interactive console setting
			continue
		}
		words = append([]string{"pulsetrainer"}, words...)
		os.Args = words
		handleError(PtRoot.Execute())
		os.Args = originalArgs
	}
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

type Config struct {
	Version  string             `yaml:"version,omitempty"`
	Debug    bool               `yaml:"debug,omitempty"`
	Routines map[string]Routine `yaml:"routines,flow"`
}

type Routine struct {
	Generators []BooleanGenerator `yaml:"generators,flow"`
	Receivers  []BooleanReceiver  `yaml:"receivers,flow"`
}

type BooleanGenerator struct {
	Pattern    []bool `yaml:"pattern,flow"`
	Pin        int    `yaml:"pin"`
	Delay      int    `yaml:"delay,omitempty"`
	ActiveStep int
	Active     bool
}

type BooleanReceiver struct {
	Pattern    []bool `yaml:"pattern,flow"`
	ActiveStep int
}
