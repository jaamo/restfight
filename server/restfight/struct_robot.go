package restfight

/**
 * Single cell object.
 */
type Robot struct {
	RobotID     int `json:"robot_id,omitempty"`
	Health      int `json:"health,omitempty"`
	MaxHealth   int `json:"max_health,omitempty"`
	Capacity    int `json:"capacity,omitempty"`
	MaxCapacity int `json:"max_capacity,omitempty"`
	X           int `json:"x,omitempty"`
	Y           int `json:"y,omitempty"`
}
