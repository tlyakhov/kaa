package main

type livingroom struct {
	room
}

func newLivingroom() *livingroom {
	return &livingroom{room: room{
		short: "You are in a large living room with a fancy home theater system and a black leather couch.",
		desc: "The room is large, but quite cozy. An unused fireplace has a giant log crammed into it. " +
			"Many colorful books line a row of shelves behind the couch. An obscenely high-end turntable sits on top of them. " +
			"You feel a little afraid of it.",
		objects: map[string]*object{},
	}}
}

func (r *livingroom) data() *room {
	return &r.room
}

func (r *livingroom) onEnter() {
	if r.state == "" {
		r.state = "visited"
	} else {
	}
}

func (r *livingroom) onExit() {}
func (r *livingroom) onLook() {}
