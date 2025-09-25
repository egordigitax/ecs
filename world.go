package ecs

type World struct {
	nextID   ID
	storages map[ComponentType]any
	systems  []System
}

func NewWorld() *World {
	return &World{
		storages: make(map[ComponentType]any),
	}
}

func (w *World) Update() {
	for _, sys := range w.systems {
		sys.Update(w)
	}
}

func (w *World) NewEntity() Entity {
	id := w.nextID
	w.nextID++
	return Entity{id: id}
}

func QueryN(w *World, types []ComponentType, cb func(e Entity, comps []any)) {
	if len(types) == 0 {
		return
	}

	storages := make([]map[ID]any, len(types))
	for i, t := range types {
		if s, ok := w.storages[t]; ok {
			storages[i] = s.(interface{ Raw() map[ID]any }).Raw()
		} else {
			return
		}
	}

	minIdx := 0
	for i := 1; i < len(storages); i++ {
		if len(storages[i]) < len(storages[minIdx]) {
			minIdx = i
		}
	}

	for id := range storages[minIdx] {
		comps := make([]any, len(storages))
		match := true
		for i, st := range storages {
			if v, ok := st[id]; ok {
				comps[i] = v
			} else {
				match = false
				break
			}
		}
		if match {
			cb(Entity{id: id}, comps)
		}
	}
}
