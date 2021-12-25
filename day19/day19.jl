include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]

function test(left, right)
    a = Set([])
    b = Set([])
    c = Set([])

    for p1 in eachrow(left)
        for p2 in eachrow(left)
            push!(a, p1 - p2)
            #push!(a, p1 - left[1, :])
        end
    end

    for p1 in eachrow(right)
        for p2 in eachrow(right)
            push!(b, p1 - p2)
            #push!(b, p1 - right[1, :])
        end
    end
    for (i, r) in enumerate(rall)
        c = [r * z for z in b]
        result = intersect(a, c)
        #println(length(result))
        if length(result) == 133
            #println(length(result))
            return (r,true)
        end
    end
    rall[1],false
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


function test2()
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
end

function getSequence() 


end
 

println(test(origin, exScans[2]))
println(test(exScans[2], exScans[1]))
println(test(exScans[2], exScans[3]))
println(test(exScans[2], exScans[4]))
