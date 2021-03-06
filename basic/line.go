package basic

import "github.com/terranodo/tegola"

// Line is a basic line type which is made up of two or more points that don't
// interect.
// TODO: We don't really check to make sure the points don't intersect.
type Line []Point

// Just to make basic collection only usable with basic types.
func (Line) basicType() {}

// Subpoints return the points in a line.
func (l *Line) Subpoints() (points []tegola.Point) {
	points = make([]tegola.Point, 0, len(*l))
	for i := range *l {
		points = append(points, tegola.Point(&((*l)[i])))
	}
	return points
}

// MultiLine is a set of lines.
type MultiLine []Line

// Just to make basic collection only usable with basic types.
func (MultiLine) basicType() {}

// Lines are the lines in a Multiline
func (ml *MultiLine) Lines() (lines []tegola.LineString) {
	lines = make([]tegola.LineString, 0, len(*ml))
	for i := range *ml {
		lines = append(lines, tegola.LineString(&(*ml)[i]))
	}
	return lines
}
