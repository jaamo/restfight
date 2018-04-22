package restfight

// ArenaTypeEmpty is constant for empty cell.
const ArenaTypeEmpty = 0

// ArenaTypeRobot is constant for a cell with a robot.
const ArenaTypeRobot = 1

// ArenaTypeRobot is constant for a cell with a robot.
const ArenaTypeObstacle = 2

// Status
var status Status

// ArenaSize defines the size for the game arena. Arena is always square.
const ArenaSize = 10

// Arena
var arena [ArenaSize][ArenaSize]Cell

// Robots
var robots []*Robot

// Timestamp when the game ended.
var gameOverTime int64
