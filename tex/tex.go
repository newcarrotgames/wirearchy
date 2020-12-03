package tex

type Texture int

const Air Texture = 0
const Grass Texture = 1
const PlanksSpruce Texture = 0
const PlanksOak Texture = 1
const PlanksBigOak Texture = 0
const PlanksAcacia Texture = 1
const Cobblestone Texture = 0
const NetherBrick Texture = 1
const StoneSlabTop Texture = 0
const IronBlock Texture = 1

type Material struct {
	Health  float64
	Texture Texture
	Mass    float64
}
