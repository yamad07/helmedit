package main

import (
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/yamad07/helmedit/pkg/chart"
)

func edit(c *cli.Context) error {
	fname := c.String("file")
	valfile, err := chart.NewValueFile(fname)
	if err != nil {
		return err
	}

	vals := strings.Split(c.String("values"), ",")
	for _, val := range vals {
		kv := strings.Split(val, "=")
		key := kv[0]
		val := kv[1]

		err := valfile.Set(key, val)
		if err != nil {
			return err
		}
	}
	err = valfile.Save()
	if err != nil {
		return err
	}

	return nil

}
