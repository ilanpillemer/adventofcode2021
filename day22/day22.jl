println("day22")

X = []
Y = []
Z = []
steps = []

struct Step
    x::Any
    x1::Any
    y::Any
    y1::Any
    z::Any
    z1::Any
    state::Any
end

function on(xs, ys, zs)
    step = Step(
        first(xs),
        last(xs) + 1,
        first(ys),
        last(ys) + 1,
        first(zs),
        last(zs) + 1,
        "on",
    )
    push!(steps, step)

    coords(xs, ys, zs)
end

function off(xs, ys, zs)
    step = Step(
        first(xs),
        last(xs) + 1,
        first(ys),
        last(ys) + 1,
        first(zs),
        last(zs) + 1,
        "off",
    )
    push!(steps, step)
    coords(xs, ys, zs)
end

function coords(xs, ys, zs)
    push!(X, first(xs))
    push!(X, last(xs) + 1)
    push!(Y, first(ys))
    push!(Y, last(ys) + 1)
    push!(Z, first(zs))
    push!(Z, last(zs) + 1)
end

#include("example2.jl") # load data
include("input.jl") # load data # answer should be 1125649856443608

function part2(X, Y, Z)
    sort!(X)
    sort!(Y)
    sort!(Z)
    N = length(X)
    grid = zeros(Bool, N, N, N)

    allsteps = length(steps)
    for (i, s) in enumerate(steps)
        #println("processing step $i of $allsteps")
        x = findfirst(x -> x == s.x, X)
        x1 = findfirst(x -> x == s.x1, X)
        y = findfirst(x -> x == s.y, Y)
        y1 = findfirst(x -> x == s.y1, Y)
        z = findfirst(x -> x == s.z, Z)
        z1 = findfirst(x -> x == s.z1, Z)
        for i = x:x1-1, j = y:y1-1, k = z:z1-1
            grid[i, j, k] = (s.state == "on")
        end
    end
    total = 0
    for x = 1:N-1, y = 1:N-1, z = 1:N-1
        process = grid[x, y, z]
        if process
            cuboid_volume = ((X[x+1] - X[x]) * (Y[y+1] - Y[y]) * (Z[z+1] - Z[z]))
            total = total + cuboid_volume
        end
    end
    total
end

part2(X, Y, Z)
