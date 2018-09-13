package main

import "github.com/fatih/color"

type kitchen struct {
	room
}

func newKitchen() *kitchen {
	return &kitchen{room: room{
		short: "You are in a clean, white '40s era kitchen with modern appliances. " +
			"The refrigerator is making groaning noises. To the north is the living room.",
		desc:    "The kitchen has many pots & pans, pantry items, cooking utensils, serving utensils, and cleaning supplies.",
		objects: map[string]*object{},
	}}
}

func (r *kitchen) data() *room {
	return &r.room
}

func (r *kitchen) onEnter() {
	if r.state == "" {
		writeLn(color.HiBlackString("You wonder where Tim could have gone. " +
			"You hear a faint noise coming from somewhere in the back of the " +
			"house. Probably nothing."))
		r.state = "visited"
	} else {
	}
}

func (r *kitchen) onExit() {}
func (r *kitchen) onLook() {}
