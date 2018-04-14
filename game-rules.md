# Game rules

## Game area

Robots are fighting in a 10 x 10 grid. Each cell can contain:

- Nothing
- An obstable (like wall, tree etc.)
- A robot

The grid is surrounded by a wall and robots can't leave the game area.

## Robot initialisation

### Capacity

Robot has capacity of 10 slots. Player can decide how the robot is built and different features (weapons, shields etc) takes space. Players can add features until capacity is full.

### Health

Robots has 10 HP by default. Adding shields can increase HP. When HP goes to zero robot is destroyed.

## Robot features

Robot can have features: engine, radar, shield, weapon and extras. Player can install aither one or zero of these to his robot. Multiple similar type of features can't be installed. E.g. player can install only one weapon.

### Engine / movement

Engine moves the robot around.

| Engine        | Capacity      | Moves per turn    |
| ------------- |-------------- |-------------------|
| engine1       | 2             | 2                 |
| engine2       | 4             | 4                 |
| engine3       | 6             | 6                 |

Engine defines how many steps a robot can take in a turn. 

* Robot can move up, down, left and right.
* Diagonal movement is not possible.
* Robots can't move over obstacles.
* Robots can't move over robots.
* Robots can't move outside game area.

If player requests illegal move the request is ignored.

### Radar

Radar is the "eyes of the robot".

| Radar         | Capacity      | Scan area         |
| ------------- |-------------- |-------------------|
| radar1        | 2             | 3x3               |
| radar2        | 4             | 6x6               |
| radar3        | 6             | 12x12             |

With radar robot can scan it's surroundings. Radar scan area is a square. Scan result contains information about each cell:

* Empty cells
* Osbtables
* Another robots
* Area outsite the grid

Radar scan can be done multiple times during a turn. You can scan, move the robot and then again.

### Shield

With shields players can increase robot's health.

| Shield        | Capacity      | Additional HP     |
| ------------- |-------------- |-------------------|
| shield1       | 2             | +2                |
| shield2       | 4             | +4                |
| shield3       | 6             | +8                |

### Weapons

Currently two weapons are available: gun and cannon.

| Weapon        | Capacity      | Range             | Strentgh       | Other                  |
| ------------- |-------------- |-------------------|----------------|------------------------|
| gun           | 2             | 6x6               | -3 HP          |                        |
| cannon        | 4             | 3x3               | -8 HP          | Takes one turn to load |

Player can launch a weapon during robot's turn. Weapon is always targeted to given cell. 

## Game preparations

Before the game starts

* Join a game
* Deploy a robot (add features)
* End deployment

## Between turns

Game status can be queried from the server.

## Turn

During single turn following events are possible:

* Start turn (mandatory)
* Radar scan (optional)
* Movement (optional)
* Shoot (optional)
* End turn (mandatory)

Scan, movement and shoot can occur in any order. Movement steps can be separated in to multiple steps. 

## Further ideas

* Multiplayer support
* Communication between robots
* Mines
