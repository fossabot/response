/*
 *   Copyright (c) 2020 
 *   All rights reserved.
 */
package main

import (
	"log"
	"os"

	"github.com/responserms/response/command"
)

func main() {
	c := command.New()
	if err := c.Run(os.Args); err != nil {
		log.Fatalf("could not execute command, received error: %s", err.Error())
	}
}
