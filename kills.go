package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can`t open pid file (is server runnig?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can`t remove pid file - %s", err)
	}
}
