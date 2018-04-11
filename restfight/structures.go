package restfight

// Cell represents a single cell in the arena.
type Cell struct {
	Type  int    `json:"type"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Robot *Robot `json:"robot"`
}

// Robot object.
type Robot struct {
	RobotID     int `json:"robot_id"`
	RobotIndex  int `json:"robot_index"`
	ShieldLevel int `json:"shield_level"`
	Health      int `json:"health"`
	MaxHealth   int `json:"max_health"`
	Capacity    int `json:"capacity"`
	MaxCapacity int `json:"max_capacity"`
	X           int `json:"x"`
	Y           int `json:"y"`
	EngineLevel int `json:"engine_level"`
	MaxMoves    int `json:"max_moves"`
	Moves       int `json:"moves"`
	WeaponLevel int `json:"weapon_level"`
	WeaponRange int `json:"weapon_range"`
	WeaponPower int `json:"weapon_power"`
	WeaponAmmo  int `json:"weapon_ammo"`
}

// Radar object.
type Radar struct {
	Range int
}

/**
 * Status data model for the game.
 */
type Status struct {

	// Unique game id.
	GameID int `json:"game_id"`

	// Game status: 0 = waiting for players, 1 = robot deployment, 2 = game is on, 3 = game over
	Status GameStatus `json:"status"`

	// Active robot. 0 or 1.
	ActiveRobot int `json:"active_robot"`

	Robots *[]*Robot `json:"robots"`

	Arena *[ArenaSize][ArenaSize]Cell `json:"arena"`

	// Active robot status. 0 = waiting, 1 = turn started
	ActiveRobotStatus ActiveRobotStatus `json:"active_robot_status"`
}

/**
 * Datatype defining game status.
 */
type GameStatus int

/**
 * Game status constannts.
 */
const (
	GameStatusWaitingForPlayers GameStatus = iota
	GameStatusRunning
	GameStatusGameOver
)

/**
 * Labels for game statuses.
 */
var GameStatusLabels = [...]string{
	"WAITING_FOR_PLAYERS",
	"DEPLOYMENT",
	"RUNNING",
	"GAME_OVER",
}

/**
 * Active robot status.
 */
type ActiveRobotStatus int

/**
 * Game status constannts.
 */
const (
	ActiveRobotStatusWaiting ActiveRobotStatus = iota
	ActiveRobotStatusPlaying
)

/**
 * Labels for game statuses.
 */
var ActiveRobotStatusLabels = [...]string{
	"WAITING",
	"PLAYING",
}
