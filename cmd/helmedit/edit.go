package main

import (
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/yamad07/helmedit/pkg/chart"
)

func edit(c *cli.Context) error {
	// "image.tag=latest,image.repositoryUrl=latest"
	fname := c.Args().Get(0)
	valf, err := chart.NewValueFile(fname)
	if err != nil {
		return err
	}

	args := strings.Split(c.Args().Get(1), ",")
	for _, arg := range args {
		kv := strings.Split(arg, "=")
		key := kv[0]
		val := kv[1]

		err := valf.Set(key, val)
		if err != nil {
			return err
		}
	}
	err = valf.Save()
	if err != nil {
		return err
	}

	return nil

}
