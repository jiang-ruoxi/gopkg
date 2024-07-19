package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

var c *cron.Cron
var cronOption cron.Option

type Cron interface {
	Spec() string
	Run()
}

func InitOption() cron.Option {
	return cronOption
}

func InitFromViper(cronList []Cron, options []cron.Option) error {
	if !viper.GetBool("cron.switch") {
		return nil
	}
	c = cron.New(options...)
	for _, task := range cronList {
		if _, err := c.AddFunc(task.Spec(), task.Run); err != nil {
			return err
		}
	}
	c.Start()
	return nil
}
