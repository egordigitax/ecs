package ecs

import "reflect"

type ComponentType string

func RegisterComponent[T any]() ComponentType {
	return ComponentType(reflect.TypeOf(*new(T)).String())
}
