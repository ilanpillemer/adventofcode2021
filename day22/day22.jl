println("day22")

X = []
Y = []
Z = []
steps = []
grid = Dict()
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

include("example2.jl") # load data
sort!(X)
sort!(Y)
sort!(Z)
steps

for s in steps
    x = findfirst(x -> x == s.x, X)
    x1 = findfirst(x -> x == s.x1, X)
    y = findfirst(x -> x == s.y, Y)
    y1 = findfirst(x -> x == s.y1, Y)
    z = findfirst(x -> x == s.z, Z)
    z1 = findfirst(x -> x == s.z1, Z)
    for i = x:x1, j = y:y1, k = z:z1
        grid[(i, j, k)] = (s.state == "on")
    end
end

N = length(X)

total = 0
for x = 0:N-1, y = 0:N-1, z = 0:N-1

    if haskey(grid, (x, y, z)) && grid[(x, y, z)]
        cuboid_volume = ((X[x+1] - X[x]) * (Y[y+1] - Y[y]) * (Z[z+1] - Z[z]))
        global total = total + cuboid_volume
    end
end
total
