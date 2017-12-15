# RESTFight

RESTFight is a turn based robot AI programming game. Players controls their robots via REST API which means any language can be used to write the AI. This repo provides a game server.

The game also offers a viewer app to follow the robot fight in real time.

## Available documentation

* [Game rules](docs/game-rules.md)
* [API documentation](docs/api.json)
* [Server documentation](server/README.md)
* [Client documentation](client/README.md)
* [Example robot](docs/example-robot.md)

## Roadmap

**Game is currently higly WIP.** And as all side projects this project won't never be finished :) However the first step will be a super simple MVP version wich only supports robot movement. In the MVP version the idea is to try to avoid obstacles and find an exit door.



| Version       | Deliverable                                                                                |
| ------------- |--------------------------------------------------------------------------------------------|
| v0.1          | Documentation                                                                              |
| v0.2          | Initial prototype to create one REST endpoint.                                             |
| v0.3          | Unit tests for endpoints: join game, add feature, status, start turn, movement, end turn.  |
| v0.4          | Endpoints: join game, add feature, status, start turn, movement, end turn.                 |
| v0.5          | MVP client & backend communication.                                                        |
