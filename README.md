# livematches-service

## Endpoints:
###### Clients:
http://iafoosball.aau.dk/ws/users/{id} <br />
Substitude `{id}` with the ID of user.

###### Table:
http://iafoosball.aau.dk/ws/tables/{id} <br />
Substitude `{id}` with the ID of table.

###### Matches:
http://iafoosball.aau.dk/matches <br />
Returns all open lobbies and matches. <br />
I am thinking about rewriting this so you get a list of tables connected
liveMatches and their status.

### Admin
The first user to connect to a match/lobby is the admin (as before).

## API

### Normal User:
###### joinMatch (tested)
###### leaveMatch (tested)
###### setUsername (tested)

### Admin User:

### Table:
###### createMatch (tested)
###### addGoal (not tested)




