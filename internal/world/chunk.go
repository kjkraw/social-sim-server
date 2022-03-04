package world

type Voxel int // Temporary type (make it easier to migrate in the future
// In my model, voxels aren't like Minecraft "blocks" where they can have functionality and store info.
// My voxels are just small units of material.

const (
	chunkSizeX     = 64 // North-South axis
	chunkSizeY     = 64 // East-West axis
	chunkSizeZ     = 64 // Vertical axis
	horizontalArea = 4096
	chunkVolume    = 262144
)

func GetLocalCoords(globalX int, globalY int, globalZ int) (int, int, int) {
	return globalX % chunkSizeX, globalY % chunkSizeY, globalZ
}

type Chunk struct {
	voxels [chunkVolume]Voxel
}

func (c Chunk) getVoxelIdx(x int, y int, z int) int {
	return x + y + horizontalArea*z
}

func (c *Chunk) GetVoxel(x int, y int, z int) Voxel {
	return c.voxels[c.getVoxelIdx(x, y, z)]
}

func (c *Chunk) SetVoxel(x int, y int, z int, voxel Voxel) {
	c.voxels[c.getVoxelIdx(x, y, z)] = voxel
}
