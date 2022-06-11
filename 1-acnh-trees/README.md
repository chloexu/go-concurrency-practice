## About

This is a example code practicing interface in golang mocking the players' activity on different types of plants in Animal Crossing game. 

There are 3 interface defined, `Shakeable`, `Choppable`, and `Diggable`. Different plant type and size can choose to implement one or more of these interfaces and call the defined methods.

## Rules

All plants of `Tree` plant type can be shakeable, choppable and diggable in that the methods defined in interfaces will be available, however whether there's outcome from the action depends on other factors such as a plant's size other than plant type. For example, a mature tree is choppable and shakeable but not diggable, while a tree that's not mature yet is diggable but not choppable and shakeable. All plants of `Shrub` plant type and `Flower` plant type can not be shakeable or choppable but are diggable regard less of plant size.

## Run the game
Go to `./1-acnh-trees/` directory in local and run:
```
go run main.go
```
