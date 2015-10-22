/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"log"
	"time"
)

type Channel struct {
	ID        int
	Name      string
	BotID     int  `sql:"index"`
	Enabled   bool `sql:"default:true"`
	Plugins   []*Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Channel) LoadPlugins(b *Bot) {
	for _, plugin := range c.Plugins {
		plugin.Lua = NewLuaState(b, c, plugin)

		if plugin.File {
			if err := plugin.Lua.Lua.DoFile(plugin.Path); err != nil {
				log.Println("Erorr running plugin from file: ", err)
			}
		} else {
			if err := plugin.Lua.Lua.DoString(plugin.Text); err != nil {
				log.Println("Error running plugin from string: ", err)
			}
		}
	}
}
