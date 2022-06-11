## About

This is a card game written in golang.

This game currently takes 5 hardcoded players, there are 10 rounds. For each round, each player plays by drawing a card with number ranging 1-13. Based on the card number, player gets a score for the round. Scores for all 10 rounds will be summed up by each player and that sum will be used to determine who is the winner of the game, whoever has the highest total score will be the winner.

This program is a practice for using goroutine and channel communication in golang. 

In the code, there are 2 channels defined, one for emitting all players' score for each round to a monitor that prints out the result of the round, the other for letting the monitor thread know that all round has finished so it could announce the winner and exit the monitoring.

In the code there are 2 threads running, one of them is a goroutine that starts the game and have all players play each round, other than printing the log on starting the round, this goroutine is purely in charge of organizing the game and coordinate it to proceed, while the other thread the `Monitor` function is in charge of announcing all the game results for each round and the final winner information.

## Run the game
Go to `./0-card-game/` directory in local and run:
```
go run main.go
```

## Next steps
As the next steps of this game, I plan to write a command line tool that lets user enter the number of players and the name of each player before starting the game, instead of using hardcoded players.
