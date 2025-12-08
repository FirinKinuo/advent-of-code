package tools

import "image"

type UniquePoints map[image.Point]struct{}

func (u UniquePoints) Add(point image.Point) {
	u[point] = struct{}{}
}

type UniqueAny map[any]struct{}

func (u UniqueAny) Add(element any) {
	u[element] = struct{}{}
}

func (u UniqueAny) Exists(element any) bool {
	_, ok := u[element]
	return ok
}

type Point3D struct {
	X int
	Y int
	Z int
}
