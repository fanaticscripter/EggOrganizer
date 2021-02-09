package config

import (
	"github.com/pkg/errors"
)

type Config struct {
	Players []Player `mapstructure:"players"`
}

type Player struct {
	UserId   string `mapstructure:"user_id"`
	Nickname string `mapstructure:"nickname"`
}

func (c *Config) Validate() error {
	if len(c.Players) == 0 {
		return errors.Errorf("no players set")
	}
	for i, player := range c.Players {
		if player.UserId == "" {
			return errors.Errorf("user_id not set for %dth player", i+1)
		}
	}
	return nil
}
