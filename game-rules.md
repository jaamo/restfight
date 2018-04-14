# Game rules

RESTFight is a turn based robot combat game. Robots can move around and shoot their opponents in an arena. There are also obstacles on the arena which robots should avoid. A robot has three primary properties: engine level, shield level and weapon level. Each property affects how the robot behaves. Robots has maximum load capacity and players need to balance properties to not overload the robot.

## Goal

To destroy enemy robot by shooting it until it's health is zero or less. 

## Game area

Game arena is a grid of square cells. Arena size is 10 x 10. Each cell can contain:

- Nothing
- An obstable
- A robot

Robots can't leave the game area. 

## Game phases

The game has three phases:

- Joining a game. During this phase players join the game and specify their robot's properties.
- Turns. In this phase robots play their turns.
- Game over. The game has ended.

## Turn

During single turn following events are possible:

* Movement
* Shoot
* End turn

## Robot properties and behavior

Players can configure three main properties: engine level, shield level and weapon level. Each of these affects how the robot behaves. These properties are defined when the players joins the game and they can't be changed during the game.

### Capacity

Engine, shields and weapon level has their own "weight" and players need to consider this when configuring the robot. Robot has a maximum capacity of 10 slots which can't be exceed.

### Health

Robot has health. When health goes to zero robot is destroyed. Initial health depends on shield level.

### Engine

Engine moves the robot around.

| Engine        | Capacity      | Moves per turn    |
| ------------- |-------------- |-------------------|
| Level 1       | 2             | 2                 |
| Level 2       | 4             | 4                 |
| Level 3       | 6             | 6                 |

Engine defines how many steps a robot can take in a turn. 

* Robot can move up, down, left and right.
* Diagonal movement is not possible.
* Robots can't move over obstacles.
* Robots can't move over robots.
* Robots can't move outside game area.


### Shield

Shield defines robot's initial health.

| Shield        | Capacity      | Initial health    |
| ------------- |-------------- |-------------------|
| Level 1       | 2             | 6                 |
| Level 2       | 4             | 10                |
| Level 3       | 6             | 14                |

### Weapons

Weapon is used to shoot other players. Range defines how far a robot can shoot. Power defines how much health is reduced when ammo hits the target. Robot can shoot only once in a turn.

| Weapon        | Capacity      | Range             | Power          |
| ------------- |-------------- |-------------------|----------------|
| Level 1       | 2             | 2                 | 2              |
| Level 2       | 4             | 4                 | 4              |
| Level 3       | 6             | 6                 | 6              |
