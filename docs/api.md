# API documentation

## Robot JSON schema
JSON schema for a robot. Used in various responses.

### Properties

- robot_id (string) - Player's id.
- robot_index (number) - Player's index.
- shield_level (number) - Shield level (0-2).
- health (number) - Robot's health.
- max_health (string) - Robot's max health.
- capacity (number) - Capacity used.
- max_capacity (number) - Max capacity used.
- x (number) - X-position.
- y (number) - Y-position.
- engine_level (number) - Shield level (0-2).
- max_moves (number) - Maximum amount of moves.
- moves (number) - Moves left on this round.
- weapon_level (number) - Weapon level (0-2).
- weapon_range (number) - Weapon range.
- weapon_power (number) - Weapon power (amount of health it reduces after hit).
- weapon_ammo (number) - Ammo left on ongoing rund.

## REST endpoints

Available REST API endpoints.

* [General](#join-a-game)
* [New game](#join-a-game)
* [Join a game](#join-a-game)
* [Start turn](#start-turn)
* [Radar scan](#radar-scan)
* [Movement](#movement)
* [Shoot](#shoot)
* [End turn](#start-turn)




- - - -

### General

* Method is always `GET`
* No authorization required
* Success HTTP code is 200
* Error HTTP code is 400 CLIENT ERROR
* In case of error JSON is returned: `{ error: "ERROR_CODE", message: "Potential error message"}`





- - - -

### New game

Resets the current game and starts a new game. Handy functions if it seems that your robot will lose :)

**URL:** `/new`





- - - -

## Join a game

Join the game. Returns a robot.

**URL:** `/join`

**Parameters:**
- engineLevel: Engine level, 0-2
- shieldLevel: Shield level, 0-2
- weaponLevel: Weapon level, 0-2

**Response:** Robot object.

**Errors:**

| Code       | Reason                       |
|------------|------------------------------|
| GAME_FULL  | No game available.           |
| ERROR      | Generic error.               |





- - - -

## Status

This endpoint returns a status of the game. This should be polled while waiting for own turn.

**URL** : `/status`

**Content example**

```json
{
    "your_turn": 1,
    "current_robot_id": "[Current robots id]",
    "current_robot_status": "[waiting, playing]"
}
```

**Errors**

| Key               | Reason                       |
|-------------------|------------------------------|
| ERROR             | Generic error.               |
| ROBOT_NOT_EXISTS  | Given robot doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Start turn

Starts turn. Should be called before any turn actions (movement, radar or weapon).

**URL** : `/api/{GAME ID}/{ROBOT ID}/turn/start`

**Method** : `GET`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "ERROR",
    "message": "Description."
}
```

**Errors**

| Key               | Reason                       |
|-------------------|------------------------------|
| ERROR             | Generic error.               |
| NOT_YOUR_TURN     | Not your turn.               |
| ROBOT_NOT_EXISTS | Given robot doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Radar scan

Scan robots surroundings. Scanning middlepoint is always in robot's coordinates and the scanned area depends on radars's type. This endpoint returns a list of coordinates and their contents. Content could be:

* `empty`- empty cell
* `enemy` - enemy
* `obstable` - any obstacle, robot can't move to this cell
* `outside`- outside game area

**URL** : `/api/{GAME ID}/{ROBOT ID}/radar`

**Method** : `POST`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

**Content example**

```
{
    { "x": 1, "y": 1, "object": "empty" },
    { "x": 2, "y": 1, "object": "enemy" },
    { "x": 3, "y": 1, "object": "obstacle" },
    ...
}
```

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "ERROR",
    "message": "Description."
}
```

**Errors**

| Key               | Reason                       |
|-------------------|------------------------------|
| ERROR             | Generic error.               |
| NOT_YOUR_TURN     | Not your turn.               |
| ROBOT_NOT_EXISTS | Given robot doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Movement

Move robot around. Robot can be move only one step up, down, left or right. Diagonal movement is not possible. If you want to move multiple steps you need to call this endpoint multiple times.

**URL** : `/api/{GAME ID}/{ROBOT ID}/move`

**Method** : `POST`

**Auth required** : NO

**Content example**

```json
{
    "x": "[Number, coordinate]",
    "y": "[Number, coordinate]",
}
```

### Success Response

**Code** : `200 OK`

**Content example**

```
{
    "moves_left": "[Number, amount of moves left]",
    "robot": "[Object, robot object]"
}
```

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "ERROR",
    "message": "Description."
}
```

**Errors**

| Key               | Reason                                                         |
|-------------------|----------------------------------------------------------------|
| ERROR             | Generic error.                                                 |
| OUT_OF_BOUNDS     | Out of bounds.                                                 |
| INVALID_MOVE      | Illegal move. Only one step up, down, left, right allowed.     |
| NOT_YOUR_TURN     | Not your turn.                                                 |
| ROBOT_NOT_EXISTS | Given robot doesn't exists                                    |
| GAME_NOT_EXISTS   | Given game doesn't exists                                      |





- - - -

## Shoot

Shoot to given coordinates. Gun can be launched on each turn. Cannon can be launched every second turn.

**URL** : `/api/{GAME ID}/{ROBOT ID}/shoot`

**Method** : `POST`

**Auth required** : NO

**Content example**

```json
{
    "x": "[Number, coordinate]",
    "y": "[Number, coordinate]",
}
```

### Success Response

**Code** : `200 OK`

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "ERROR",
    "message": "Description."
}
```

**Errors**

| Key                 | Reason                                                         |
|---------------------|----------------------------------------------------------------|
| ERROR               | Generic error.                                                 |
| OUT_OF_BOUNDS       | Out of bounds.                                                 |
| ILLEGAL_COORDINATES | Illegal move. Only one step up, down, left, right allowed.     |
| NOT_YOUR_TURN       | Not your turn.                                                 |
| ROBOT_NOT_EXISTS   | Given robot doesn't exists                                    |
| GAME_NOT_EXISTS     | Given game doesn't exists                                      |





- - - -

## End turn

Ends robot's turn.

**URL** : `/api/{GAME ID}/{ROBOT ID}/turn/end`

**Method** : `GET`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "ERROR",
    "message": "Description."
}
```

**Errors**

| Key               | Reason                       |
|-------------------|------------------------------|
| ERROR             | Generic error.               |
| NOT_YOUR_TURN     | Not your turn.               |
| ROBOT_NOT_EXISTS | Given robot doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |

