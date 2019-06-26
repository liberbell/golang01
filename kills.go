package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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

	strPID := strings.TimeSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process id")
	}

	fmt.Printf("killing server with pid=%d\n", pid)
}
