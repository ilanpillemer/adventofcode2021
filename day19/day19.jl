include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]

function test(left, right)
    found = false
    i = 1
    while i < length(left) / 3
        a = Set([])
        b = Set([])
        c = Set([])

        for p1 in eachrow(left)
            for p2 in eachrow(left)
                #push!(a, p1 - p2)
                push!(a, p1 - left[i, :])
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
            if length(result) == 12
                #println(length(result))
                return (r, true)
            end
        end
        i = i + 1
    end
    (rall[1], false)
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

function getNext(scan, scans)
    #find link from one scanner to next scanner
    n = length(scans)
    for i = 1:length(scans)
        (rot, ok) = test(scan, scans[i])
        if ok
            #println("scanner $i is next scanner")
            return rot, i
        end
    end
end

function getSequence(scans)
    for (i, scan) in enumerate(scans)
        (rot, j) = getNext(scan, scans)
        println("Scanner $i links to Scanner $j with rotation $rot")
    end
end


#println(test(origin, exScans[2]))
#println(test(exScans[2], exScans[1]))
#println(test(exScans[2], exScans[3]))
#println(test(exScans[2], exScans[4]))
#println(test(exScans[2], exScans[5]))

#getNext(origin, exScans)
getSequence(exScans)
