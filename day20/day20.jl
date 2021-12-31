# provides rawImgEnh and img
#include("sample.jl")
# part 1 is 5057
# part 2 is 18502
include("input.jl")
testpos = parse(Int64, "000100010", base = 2)
getNext(x) = rawImgEnh[x]
total(d) = sum([x - '0' for x in values(d)])

function D()
    d = Dict()
    for (y, v) in enumerate(img)
        for (x, v2) in enumerate(v)
            d[(x, y)] = v2
        end
    end
    d
end

function lookup(dict, x, y, round)
    border = "0"
    #odd is 0
    #even is 1
    if round % 2 == 0
        border == "1"
    end


    if !haskey(dict, (x - 1, y - 1))
        dict[(x - 1, y - 1)] = border
    end
    if !haskey(dict, (x - 1, y))
        dict[(x - 1, y)] = border
    end
    if !haskey(dict, (x - 1, y + 1))
        dict[(x - 1, y + 1)] = border
    end

    if !haskey(dict, (x, y - 1))
        dict[(x, y - 1)] = border
    end
    if !haskey(dict, (x, y))
        dict[(x, y)] = border
    end
    if !haskey(dict, (x, y + 1))
        dict[(x, y + 1)] = border
    end

    if !haskey(dict, (x + 1, y - 1))
        dict[(x + 1, y - 1)] = border
    end
    if !haskey(dict, (x + 1, y))
        dict[(x + 1, y)] = border
    end
    if !haskey(dict, (x + 1, y + 1))
        dict[(x + 1, y + 1)] = border
    end

    p1 = dict[(x - 1, y - 1)]
    p2 = dict[(x, y - 1)]
    p3 = dict[(x + 1, y - 1)]

    p4 = dict[(x - 1, y)]
    p5 = dict[(x, y)]
    p6 = dict[(x + 1, y)]

    p7 = dict[(x - 1, y + 1)]
    p8 = dict[(x, y + 1)]
    p9 = dict[(x + 1, y + 1)]

    lu = p1 * p2 * p3 * p4 * p5 * p6 * p7 * p8 * p9
    pos = parse(Int64, lu, base = 2)
    pos = pos + 1 # julia is one based
    getNext(pos)
end

function turn(d, round)
    border = '1'
    #odd is 1
    #even is 0
    if round % 2 == 0
        border = '0'
    end
    z = Dict()
    for x = -300:1:300, y = -300:1:300
        z[(x, y)] = lookup(d, x, y, round)
    end
    for (x, y) in keys(z)
        if x == -300 || x == 300
            z[(x, y)] = border
        end
    end
    z
end

function display(thing)
    for y = -300:1:-290
        for x = -300:1-290
            if haskey(thing, (x, y))
                print(thing[(x, y)])
            else
                print('0')
            end
        end
        println()
    end
end

global next = D()
for i = 1:50
    global next = turn(next, i)
end
println(total(next))
