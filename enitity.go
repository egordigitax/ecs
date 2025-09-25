package ecs

type Entity struct {
	id ID
}

func (e Entity) ID() ID { return e.id }
