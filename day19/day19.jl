include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]


#c = [rot * z for z in keys(right)]
#
#julia> d = intersect(keys(left),c)
#Set{Any} with 12 elements:
#  [-941, -235, 443]
#  [-1065, -228, 326]
#  [140, -39, 11]
#  [-851, 259, 1219]
#  [-749, 277, 1282]
#  [19, -113, 1335]
#  [-14, -87, 108]
#  [124, -55, 1310]
#  [-889, 231, 1248]
#  [55, -119, 1302]
#  [0, 0, 0]
#  [-1022, -236, 280]
#
#julia> right[rot*([-1065, -228, 326])]
#3-element view(::Matrix{Int64}, 11, :) with eltype Int64:
# 729
# 430
# 532
#
#julia> left[[-1065, -228, 326]]
#3-element view(::Matrix{Int64}, 8, :) with eltype Int64:
# -661
# -816
# -575
#julia> x=left[[-1065, -228, 326]]
#3-element view(::Matrix{Int64}, 8, :) with eltype Int64:
# -661
# -816
# -575
#
#julia> y=right[rot*([-1065, -228, 326])]
#3-element view(::Matrix{Int64}, 11, :) with eltype Int64:
# 729
# 430
# 532
#
#julia> y=-rot*y
#3-element Vector{Int64}:
#  729
# -430
#  532
#
#julia> x-y
#3-element Vector{Int64}:
# -1390
#  -386
# -1107
#
#julia> x+y
#3-element Vector{Int64}:
#    68
# -1246
#   -43



#Because of this, scanner 1 must be at 68,-1246,-43 (relative to scanner 0

function test(left, right)
    found = false
    i = 1
    while i < length(left) / 3
        a = Dict()
        b = Dict()
        for p1 in eachrow(left)
            for p2 in eachrow(left)
                if p1 != p2
                    a[p1-left[i, :]] = p1
                end
            end
        end
        for p1 in eachrow(right)
            for p2 in eachrow(right)
                if p1 != p2
                    b[p1-p2] = p1
                end
            end
        end
        for (i, r) in enumerate(rall)
            c = [r * z for z in keys(b)]
            result = intersect(keys(a), c)
            #println(length(result))
            if length(result) == 11
                #println(result)
                return (r, true, a, b)
            end
        end
        i = i + 1
    end
    (rall[1], false, Dict(), Dict())
end


function checkRotation(origin, point)
    found = 0
    for row = 1:div(length(origin), 3)
        if point == origin[row, :]
            found = 1
            #println("scanned", point)
        end
    end
    found
end

struct Link
    from::Any
    to::Any
    rot::Any
    left::Any
    right::Any
end

function getNext(j, scans)
    #find link from one scanner to next scanner
    links = []
    scan = scans[j]
    n = length(scans)
    for i = 1:length(scans)
        (rot, ok, left, right) = test(scan, scans[i])
        if ok
            link = Link(j, i, rot, left, right)
            push!(links, link)
            #println(link)
            continue
        end
    end
    return links
end

function getSequence(scans)

    for i = 1:length(scans)
        #(rot, j, left, right) = getNext(i, scans)
        links = getNext(i, scans)
        for link in links
            println("$(link.from) -> $(link.to) with $(link.rot)")
        end
        for link in links
            #          #println(link)
            #            println("$(link.from) -> $(link.to)")
            #            c = [link.rot * z for z in keys(link.right)]
            #            d = collect(intersect(keys(link.left), c))
            #            common = d[1]
            #            x = link.left[common]
            #            y = link.right[link.rot*common]
            #            y = -(link.rot) * y
            #            translation = x + y
            #            #println("Scnr $(link.from-1) -> Scnr $(link.to-1) at rot $link.rot transl $translation")

        end


    end
end


#println(test(origin, exScans[2]))
#println(test(exScans[2], exScans[1]))
#println(test(exScans[2], exScans[3]))
#println(test(exScans[2], exScans[4]))
#println(test(exScans[2], exScans[5]))

#getNext(origin, exScans)
getSequence(exScans)
