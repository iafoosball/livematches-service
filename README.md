# livematches-service

## TODO:
End game at 10 Goals <br />
Logic for MaxTime to end game <br />
Multiple IDs on one Table <br />
Multiple IDs on mu.tiple Tables, kick older connection <br />
Multiple IDs on mutiple Table (kick from table which he does not connect <br />


## Future:
Add database and synch every important event. <br />

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
// { "command": "settings", "values": { "twoOnTwo": true }}

twoOnOne        = "twoOnOne"
// { "command": "settings", "values": { "oneOnTwo": true }}

oneOnOne        = "oneOnOne"
// { "command": "settings", "values": { "oneOnOne": true }}

switchPositions = "switchPositions"
// { "command": "settings", "values": { "switchPositions": true }}

bet             = "bet"
// { "command": "settings", "values": { "bet": true }}

maxGoals        = "maxGoals"
// { "command": "settings", "values": { "maxGoals": 10 }}

tournament      = "tournament"
// { "command": "settings", "values": { "tournament": true }}

startMatch      = "startMatch"
// { "command": "settings", "values": { }}

drunk           = "drunk"
// { "command": "settings", "values": { "drunk": true }}

freeGame        = "freeGame"
// { "command": "settings", "values": { "freeGame": true }}

payed           = "payed"
// { "command": "settings", "values": { "payed": true }}

maxTime         = "maxTime"
// { "command": "settings", "values": { "maxTime": 600 }}

rated           = "rated"
// { "command": "settings", "values": { "rated": true }}

cancelMatch     = "cancelMatch"
// { "command": "cancelMatch", "values": { }}

kickUser        = "kickUser"
// { "command": "settings", "values": { "kickUser": "userID" }}
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




Cert info:

IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/iafoosball.me/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/iafoosball.me/privkey.pem
   Your cert will expire on 2019-03-26. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot
   again. To non-interactively renew *all* of your certificates, run
   "certbot renew"
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le

