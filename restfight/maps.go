package restfight

func randomMap(arena *[ArenaSize][ArenaSize]Cell) {

	arena[3][1].Type = ArenaTypeObstacle
	arena[3][2].Type = ArenaTypeObstacle
	arena[3][3].Type = ArenaTypeObstacle

	arena[6][6].Type = ArenaTypeObstacle
	arena[6][7].Type = ArenaTypeObstacle
	arena[6][8].Type = ArenaTypeObstacle

}
