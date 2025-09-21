# Traffic Simulation in Go

A simple 2D traffic simulation written in Go that models cars moving through a road with traffic lights using very simple heuristic decision-making rules.

## Overview

This project simulates traffic flow on a multi-lane road with traffic lights. Cars spawn at the left side of the road, move eastward, and must obey traffic lights and avoid collisions with other vehicles.

## Features

-
- **Multi-lane road simulation**: Configurable width and height road grid
- **Traffic light system**: Red/green lights that switch automatically
- **Car spawning**: Random car generation at configurable rates
- **Collision avoidance**: Cars look ahead to avoid collisions
- **Traffic light compliance**: Cars stop at red lights

## Project Structure

```
trafficgoheur/
├── main.go              # Entry point and simulation setup
├── go.mod              # Go module definition
└── simulation/         # Core simulation package
    ├── cars.go         # Car entity and movement logic
    ├── decisions.go    # Movement decision algorithms
    ├── road.go         # Road grid and cell types
    ├── sim.go          # Main simulation engine
    ├── spawner.go      # Car spawning system
    └── traffic_lights.go # Traffic light behavior
```

## Core Components

### Road System

- 2D grid of cells representing the road
- Cell types: Empty (`.`), Car (`C`), Red Light (`R`), Green Light (`G`)

### Cars

- Move in four directions: North, South, East, West
- Currently configured to spawn moving East
- Follow collision avoidance rules

### Traffic Lights

- Automatic switching between red and green states
- Configurable switch rates (time between state changes)
- Cars must stop at red lights

### Decision Engine

- Prevents cars from colliding with each other
- Enforces traffic light compliance
- Looks ahead to avoid potential collisions

## How to Run

1. **Prerequisites**: Go 1.24.5 or later

2. **Clone and navigate to the project**:

   ```bash
   cd /Users/ashi/Desktop/trafficgoheur
   ```

3. **Run the simulation**:
   ```bash
   go run main.go
   ```

## Configuration

The simulation can be customized in `main.go`:

- **Road dimensions**: `NewRoad(width, height)`
- **Traffic light positions**: Set X, Y coordinates for each light
- **Switch rates**: Control how fast lights change (in ticks)
- **Spawn rate**: Probability of spawning a car each tick (0.0-1.0)
- **Tick delay**: Time between simulation steps

## Example Configuration

```go
road := simulation.NewRoad(30, 4)  // 30 cells wide, 4 lanes
spawner := simulation.NewSpawnner(road, 0.3)  // 30% spawn chance per tick
TickDelay: 400 * time.Millisecond  // 400ms between updates
```

## Simulation Rules

1. Cars spawn randomly on the left side of the road
2. Cars move eastward (left to right)
3. Cars stop at red traffic lights
4. Cars avoid collisions by checking 1-2 cells ahead
5. Cars disappear when they reach the right edge
6. Traffic lights switch automatically based on their switch rate

## Future Enhancements

Potential improvements could include:

- Multiple directions of traffic
- Intersections and turning
- Different car types with varying speeds
- More sophisticated traffic light timing
- Performance metrics and statistics
- GUI-based visualization
