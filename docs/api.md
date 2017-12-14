# JSON schemas

## Robot
JSON schema for a robot. Used in various responses.

### Properties

- player_id: 1 (string) - Player's id.
- health: 10 (number) - Robot's health.
- max_health: 10 (string) - Robot's max health.
- capacity: 2 (number) - Capacity used.
- max_capacity: 10 (number) - Max capacity used.

# REST endpoints

Available REST API endpoints.

* [Join a game](#join-a-game)
* [Add a feature to robot](#add-a-feature-to-robot)
* [Finish deployment](#finish-deployment)
* [Start turn](#start-turn)
* [Radar scan](#radar-scan)
* [Movement](#movement)
* [Shoot](#shoot)
* [End turn](#start-turn)





- - - -

## Join a game

Join the game. 

**URL** : `/api/join`

**Method** : `GET`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "game_id": "123456",
    "user_id": "123456"
}
```

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "GAME_FULL",
    "message": "Description."
}
```

**Errors**

| Key        | Reason                       |
|------------|------------------------------|
| GAME_FULL  | No game available.           |
| ERROR      | Generic error.               |





- - - -

## Add a feature to robot

Add a single feature (radar, shield, engine or weapon) to a robot. One feature type can be added only once. If feature already exists the existing will be deleted.

**URL** : `/api/{GAME ID}/{PLAYER ID}/feature/add`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "feature": "[feature name: radar, shield, engine, weapon]",
    "type": "[feature type]"
}
```

**Data example**

```json
{
    "username": "engine",
    "password": "engine1"
}
```

### Success Response

**Code** : `200 OK`

**Content example**

Robot object.

### Error Response

**Code** : `400 CLIENT ERROR`

**Content** :

```json
{
    "error": "OUT_OF_CAPACITY",
    "message": "Description."
}
```

**Errors**

| Key              | Reason                       |
|------------------|------------------------------|
| OUT_OF_CAPACITY  | Robot's capacity exceeded.   |
| ERROR            | Generic error.               |





- - - -

## Finish deployment

This enpoint is called when robot deployment is finished.

**URL** : `/api/{GAME ID}/{PLAYER ID}/feature/end`

**Method** : `GET`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

**Content example**

Robot object.

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

| Key              | Reason                       |
|------------------|------------------------------|
| ERROR            | Generic error.               |





- - - -

## Status

This endpoint returns a status of the game. This should be polled while waiting for own turn.

**URL** : `/api/{GAME ID}/{PLAYER ID}/status`

**Method** : `GET`

**Auth required** : NO

### Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "your_turn": 1,
    "current_player_id": "[Current players id]",
    "current_player_status": "[waiting, playing]"
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
| PLAYER_NOT_EXISTS | Given player doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Start turn

Starts turn. Should be called before any turn actions (movement, radar or weapon).

**URL** : `/api/{GAME ID}/{PLAYER ID}/turn/start`

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
| PLAYER_NOT_EXISTS | Given player doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Radar scan

Scan robots surroundings. Scanning middlepoint is always in robot's coordinates and the scanned area depends on radars's type. This endpoint returns a list of coordinates and their contents. Content could be:

* `empty`- empty cell
* `enemy` - enemy
* `obstable` - any obstacle, robot can't move to this cell
* `outside`- outside game area

**URL** : `/api/{GAME ID}/{PLAYER ID}/radar`

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
| PLAYER_NOT_EXISTS | Given player doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |





- - - -

## Movement

Move robot around. Robot can be move only one step up, down, left or right. Diagonal movement is not possible. If you want to move multiple steps you need to call this endpoint multiple times.

**URL** : `/api/{GAME ID}/{PLAYER ID}/move`

**Method** : `POST`

**Auth required** : NO

**Content example**

```json
{
    "x": "[Number, coordinate]",
    "x": "[Number, coordinate]",
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
| ILLEGAL_MOVE      | Illegal move. Only one step up, down, left, right allowed.     |
| NOT_YOUR_TURN     | Not your turn.                                                 |
| PLAYER_NOT_EXISTS | Given player doesn't exists                                    |
| GAME_NOT_EXISTS   | Given game doesn't exists                                      |





- - - -

## Shoot WIP

Shoot weapon

**URL** : `/api/{GAME ID}/{PLAYER ID}/move`

**Method** : `POST`

**Auth required** : NO

**Content example**

```json
{
    "x": "[Number, coordinate]",
    "x": "[Number, coordinate]",
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
| ILLEGAL_MOVE      | Illegal move. Only one step up, down, left, right allowed.     |
| NOT_YOUR_TURN     | Not your turn.                                                 |
| PLAYER_NOT_EXISTS | Given player doesn't exists                                    |
| GAME_NOT_EXISTS   | Given game doesn't exists                                      |





- - - -

## End turn

Ends robot's turn.

**URL** : `/api/{GAME ID}/{PLAYER ID}/turn/end`

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
| PLAYER_NOT_EXISTS | Given player doesn't exists  |
| GAME_NOT_EXISTS   | Given game doesn't exists    |

