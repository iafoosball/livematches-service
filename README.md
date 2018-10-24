# livematches-service

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

### Normal User:
###### leaveMatch (tested)

### Admin User:

### Table:
###### addGoal (not tested)

### JSON
```json
[
  {
    "tableID": "table-2",
    "started": false,
    "users": {
      "user-1": {
        "username": "",
        "admin": false,
        "position": "attack",
        "bet": 100,
        "color": "#ffffff"
      },
      "user-2": {
        "username": "",
        "admin": false,
        "position": "attack",
        "bet": 0,
        "color": "#ffffff"
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
      "twoVtwo": true,
      "oneVone": true,
      "switchPositions": true,
      "twoVone": true,
      "bet": true,
      "maxGoals": true,
      "tournament": true,
      "drunk": true,
      "freeGame": true,
      "payed": true,
      "maxTime": ""
    },
    "goals" :[
      {"speed":20.0,"position":"attack/defense","side":"blue/red"},
      {"speed":20.0,"position":"attack/defense","side":"blue/red"}
    ]
  }
]
```


