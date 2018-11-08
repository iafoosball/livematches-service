# livematches-service

## TODO:
Add database and synch every important event. <br />
Hanlde erros, don't exit. <br />
Send livematch data to server <br />

## Important:
Position have to omitted when no empty.


## Endpoints:
###### Clients:
ws://iafoosball.aau.dk:9003/users/{table-id}/{user-id} <br />
Substitude `{table-id}` with the ID of table and `{user-id}` with the ID of user.

###### Table:
ws://iafoosball.aau.dk:9003/tables/{table-id} <br />
Substitude `{table-id}` with the ID of table.

###### Matches:
ws://iafoosball.aau.dk/matches <br />
Returns all open lobbies and matches. <br />
I am thinking about rewriting this so you get a list of tables connected
liveMatches and their status.

### Admin
The first user to connect to a match/lobby is the admin (as before).

## API

#### For Users
```
setPosition = "setPosition"
// { "command": "setPosition", "values": { "side": "red", "position": "attack" }}

setColor    = "setColor"
// { "command": "setColor", "values": { "color": "green" }}

setUsername = "setUsername"
// { "command": "setUsername", "values": { "username": "joe" }}

setBet      = "setBet"
// { "command": "setBet", "values": { "bet": 123 }}

ready       = "ready"
// { "command": "ready", "values": { "ready": true }}

leaveMatch = "leaveMatch"
// { \"command\": \"leaveMatch\", \"values\": { }}
```


#### For Admin
```
twoOnTwo        = "twoOnTwo"
// { "command": "twoOnTwo", "values": { "twoOnTwo": true }}

twoOnOne        = "twoOnOne"
// { "command": "oneOnTwo", "values": { "oneOnTwo": true }}

oneOnOne        = "oneOnOne"
// { "command": "oneOnOne", "values": { "oneOnOne": true }}

switchPositions = "switchPositions"
// { "command": "switchPositions", "values": { "switchPositions": true }}

bet             = "bet"
// { "command": "bet", "values": { "bet": true }}

maxGoals        = "maxGoals"
// { "command": "maxGoals", "values": { "maxGoals": 10 }}

tournament      = "tournament"
// { "command": "tournament", "values": { "tournament": true }}

startMatch      = "startMatch"
// { "command": "startMatch", "values": { }}

drunk           = "drunk"
// { "command": "drunk", "values": { "drunk": true }}

freeGame        = "freeGame"
// { "command": "freeGame", "values": { "freeGame": true }}

payed           = "payed"
// { "command": "payed", "values": { "payed": true }}

maxTime         = "maxTime"
// { "command": "maxTime", "values": { "maxTime": 600 }}

rated           = "rated"
// { "command": "rated", "values": { "rated": true }}

cancelMatch     = "cancelMatch"
// { "command": "candelMatch", "values": { }}

kickUser        = "kickUser"
// { "command": "kickUser", "values": { "kickUser": "userID" }}
```

#### For Table, possible by admin as well
```
addGoal    = "addGoal"
// { "command": "addGoal", "values": { "speed": 12, "side": "blue", "position": "attack"  }}

removeGoal = "removeGoal"
// { "command": "removeGoal", "values": { "side": "blue" }}
```

### JSON
```json
[
  {
    "tableID": "table2",
    "started": false,
    "users": {
      "user1": {
        "username": "",
        "admin": false,
        "position": "attack",
        "bet": 100,
        "color": "#ffffff",
        "ready": true
      },
      "user2": {
        "username": "",
        "admin": false,
        "position": "attack",
        "bet": 0,
        "color": "#ffffff",
        "ready": false

      }
    },
    "positions": {
      "blueDefense": "uid",
      "blueAttack": "uid",
      "redDefense": "uid",
      "redAttack": "uid"
    },
    "scoreRed": 0,
    "scoreBlue": 0,
    "settings": {
      "twoOnTwo": true,
      "oneOnOne": true,
      "switchPositions": true,
      "twoOnOne": true,
      "bet": true,
      "maxGoals": true,
      "tournament": true,
      "drunk": true,
      "freeGame": true,
      "payed": true,
      "maxTime": "",
      "rated" : true
    },
    "goals" :[
      {"speed":20.0,"position":"attack/defense","side":"blue/red"},
      {"speed":20.0,"position":"attack/defense","side":"blue/red"}
    ]
  }
]
```


