# Forest fire simulator

This repository contains my solution for the forest struck by lighting problem. It involves conducting multiple
simulations of a forest struck by lightning in random
location. The goal is to find the maximum forest coverage percentage where the number of burned trees remains below a
specified threshold. If this threshold is exceeded, the suggested forest coverage percentage is considered suboptimal
and rejected. This optimization problem is crucial for balancing forest conservation and fire risk mitigation.

The forest is represented using a nested two-dimensional array. I have chosen **Go** (Golang) for this problem, because
it
is a highly performant language, which is crucial for this type of problem, which requires handling with a lot of heavy
calculations. Additionally, I wanted to learn and explore this language.

### How to run it locally?

1. **Clone the repository**
2. **Run simulation with your desired parameters**

```shell
go run . -x 10 -y 10 -maxBurn 20
```

You can also look up more information by using help.

```shell
go run . --help
```
