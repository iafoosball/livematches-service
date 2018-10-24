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




