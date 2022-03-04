package world

type World struct {
	chunks [][]*Chunk
}

func (w *World) GetChunk(x int, y int) *Chunk {
	x_idx := x / chunkSizeX
	y_idx := y / chunkSizeY
	return w.chunks[x_idx][y_idx]
}

func (w *World) GetVoxel(x int, y int, z int) Voxel {
	chunk := w.GetChunk(x, y)
	localX, localY, localZ := GetLocalCoords(x, y, z)
	return chunk.GetVoxel(localX, localY, localZ)
}
