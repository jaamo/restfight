package restfight

// Cell represents a single cell in the arena.
type Cell struct {
	// Cell type. 0 = empty, 1 = robot, 2 = obstacle, 3 = powerup.
	Type    int      `json:"type"`
	X       int      `json:"x"`
	Y       int      `json:"y"`
	Robot   *Robot   `json:"robot"`
	Powerup *Powerup `json:"poweup"`
}

// Robot object.
type Robot struct {
	RobotID     int      `json:"robot_id"`
	RobotIndex  int      `json:"robot_index"`
	ShieldLevel int      `json:"shield_level"`
	Health      int      `json:"health"`
	MaxHealth   int      `json:"max_health"`
	Capacity    int      `json:"capacity"`
	MaxCapacity int      `json:"max_capacity"`
	X           int      `json:"x"`
	Y           int      `json:"y"`
	EngineLevel int      `json:"engine_level"`
	MaxMoves    int      `json:"max_moves"`
	Moves       int      `json:"moves"`
	WeaponLevel int      `json:"weapon_level"`
	WeaponRange int      `json:"weapon_range"`
	WeaponPower int      `json:"weapon_power"`
	WeaponAmmo  int      `json:"weapon_ammo"`
	Powerup     *Powerup `json:"powerup"`
}

// Powerup object.
type Powerup struct {
	// See API documentation for available types
	Type               string `json:"type"`
	RobotLifetime      int    `json:"robot_lifetime"`
	ArenaLifetime      int    `json:"arena_lifetime"`
	AdditionalMaxMoves int    `json:"additional_max_moves"`
	X                  int    `json:"x"`
	Y                  int    `json:"y"`
}

// Status data model for the game.
type Status struct {

	// Unique game id.
	GameID int `json:"game_id"`

	// Game status: 0 = waiting for players, 1 = robot deployment, 2 = game is on, 3 = game over
	Status GameStatus `json:"status"`

	// Active robot. 0 or 1.
	ActiveRobot int `json:"active_robot"`

	// Is your turn. 0 or 1.
	IsYourTurn int `json:"is_your_turn"`

	// Player's own robot.
	Robot *Robot `json:"robot"`

	// List of enemies.
	Enemies []*Robot `json:"enemies"`

	// All robots.
	Robots *[]*Robot `json:"robots"`

	// All powerups
	Powerups *[]*Powerup `json:"powerups"`

	Arena *[ArenaSize][ArenaSize]Cell `json:"arena"`
}
