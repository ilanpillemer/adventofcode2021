include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]


function checkRotation(origin, point)
    found = 0
    for row = 1:div(length(origin), 3)
        if point == origin[row, :]
            found = 1
            #println("scanner", scan)
            #println("translated", x1)
            #println("rotated", rot * p)
            println("scanned", point)
        end
    end
    found
end


for scan = 2:length(exScans)
    for (i, rot) in enumerate(rall)
        println("rot", i)
        found = 0
        for p in eachrow(exScans[scan])
            #println("...",rot*p,typeof(rot*p))
            x1 = rot * p
            x1 = x1 + [68, -1246, -43]
            found = found + checkRotation(origin, x1)
        end
        if found != 0
            println("found $found matches at rotation $i")
        end
    end
end
