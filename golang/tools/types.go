package tools

import "image"

type UniquePoints map[image.Point]struct{}

func (u UniquePoints) Add(point image.Point) {
	u[point] = struct{}{}
}
