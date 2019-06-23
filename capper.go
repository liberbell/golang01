package main

import "io"

type Capper struct {
	wtr io.Writer
}
