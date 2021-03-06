package mvt

import (
	"fmt"

	"github.com/terranodo/tegola/mvt/vector_tile"
)

//Tile describes a tile.
type Tile struct {
	layers []Layer
}

//AddLayer adds a Layer to the tile
func (t *Tile) AddLayer(nl *Layer) error {
	// Need to make sure that all layer names are unique.
	for i, l := range t.layers {
		if l.Name == nl.Name {
			return fmt.Errorf("Layer %v, already is named %v, new layer not added.", i, l.Name)
		}
	}
	t.layers = append(t.layers, *nl)
	return nil
}

// Layers returns a copy of the layers in this tile.
func (t *Tile) Layers() (l []Layer) {
	l = append(l, t.layers...)
	return l
}

// VTile returns a tile object according to the Google Protobuff def. This function
// does the hard work of converting everthing to the standard.
func (t *Tile) VTile() (vt *vectorTile.Tile, err error) {
	vt = new(vectorTile.Tile)
	for _, l := range t.layers {
		vtl, err := l.VTileLayer()
		if err != nil {
			return nil, err
		}
		vt.Layers = append(vt.Layers, vtl)
	}
	return vt, nil
}

//TileFromVTile will return a Tile object from the given vectorTile Tile object
func TileFromVTile(t *vectorTile.Tile) (*Tile, error) {
	return nil, nil
}
