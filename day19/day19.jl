include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]

function test()
    a = Set([])
    b = Set([])

    for p1 in eachrow(origin)
        for p2 in eachrow(origin)
            push!(a, p1 - p2)
        end
    end

    for p1 in eachrow(exScans[2])
        for p2 in eachrow(exScans[2])
            push!(b, p1 - p2)
        end
    end
    for (i,r) in enumerate(rall)
        c = [r * z for z in b]
        intersect(a, c)
        println(length(intersect(a, c)))
    end
end


function checkRotation(origin, point)
    
    found = 0
    for row = 1:div(length(origin), 3)
        if point == origin[row, :]
            found = 1
            println("scanned", point)
        end
    end
    found
end


for scan = 2:length(exScans)
    for (i, rot) in enumerate(rall)
        #println("rot", i)
        found = 0
        for p in eachrow(exScans[scan])
            #println("...",rot*p,typeof(rot*p))
            x1 = rot * p
            x1 = x1 + [68, -1246, -43]
            found = found + checkRotation(origin, x1)
        end
        if found != 0
            println("found $found matches at rotation $i for scanner $scan")
        end
    end
end

test()
