package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

var c *cron.Cron

type Cron interface {
	Spec() string
	Run()
}

func InitFromViper(cronList []Cron) error {
	if !viper.GetBool("cron.switch") {
		return nil
	}
	c = cron.New()
	for _, task := range cronList {
		if _, err := c.AddFunc(task.Spec(), task.Run); err != nil {
			return err
		}
	}
	c.Start()
	return nil
}
