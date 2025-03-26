# Guidelines For Functions

## Function Naming

- give functions a good name
  - behavior can be understood at a glance
  - parameter naming counts too

```golang
func ProcessArray(a[]int)float{}
func ComputeRMS(samples[]float)float{}
```

- RMS = Root Meat Square
- `samples` is a slice of samples of a time-varying signal

## Functional Cohesion

- function should perform only one "operation"
- an "operation" depends on the context
  - ex: geometry application
- good functions:
  - `PointDist()`, `DrawCircle()`, `TriangleArea()`
- merging behaviors makes code complicated
  - `DrawCircle()` + `TriangleArea()`
    - this merge doesn't even make sense!

## Few Parameters

- debugging requires tracing function input data
- more difficult with a large numver of parameters
- function may have bad functional cohesion
  - `DrawCircle()` and `TriangleArea()` require different arguments
    - if merged, then the new function will have more parameters

## Reducing Parameter Number

- may need to group related arguments into structures
- `TriangleArea()`, bad solution
  - 3 points needed to define triangle
  - each point has 3 floats (in 3D)
  - total, 9 arguments
- `TriangleArea()`, good solution
  - `type Point struct {x,y,z float}`
  - total, 3 arguments
