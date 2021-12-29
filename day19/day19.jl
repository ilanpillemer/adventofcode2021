include("rots.jl")
println("day 19")
include("example.jl") # load exScans
include("input.jl") # load  inputScans

#scans = exScans
scans = inputScans
origin = scans[1]

struct Link
    from::Any
    to::Any
    rot::Any
    trans::Any
end



#struct Tree
#    parent::tree
#    val::DecoratedLink
#    children::Vector{Tree}
#end


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
            c = Dict([((r * z), r * b[z]) for z in keys(b)])
            result = intersect(keys(a), keys(c))
            c = filter(p -> first(p) in result, c)
            d = filter(p -> first(p) in result, a)

            #println(length(result))
            if length(result) == 11
                key = first(result)
                transl = d[key] - c[key]
                return (r, true, transl)
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



function getNext(j, scans)
    #find link from one scanner to next scanner
    links = []
    scan = scans[j]
    n = length(scans)
    for i = 1:length(scans)
        (rot, ok, transl) = test(scan, scans[i])
        if ok
            #link = Link(j, i, rot, transl)

            link = Link(i, j, rot, transl)
            #println("loading.. $(link.from) -> $(link.to)")
            push!(links, link)
            #println(link)
            continue
        end
    end
    return links
end

function getTranslation(link)
    c = Dict([((link.rot * z), z) for z in keys(link.right)])
    d = collect(intersect(keys(link.left), keys(c)))
    common = first(d)
    inverse = c[common]
    translation = link.right[inverse] - ((link.rot) * link.left[common])
    translation
end

function getSequence(scans)
    connects = []
    for i = 1:length(scans)
        links = getNext(i, scans)
        for link in links
            push!(connects, link)
        end
    end
    connects
end




function getTree(nodes)
    d = Dict()
    for node in nodes
        xs = get(d, node.from, [])
        d[node.from] = xs
        d[node.from] = push!(d[node.from], node.to)
    end
    d
end

tree = getTree(nodes)
println(tree[2])

function getPath(tree, acc)
    if length(acc) > 0 && acc[length(acc)] == 1
        return acc
    end
    for xs in values(tree[acc[length(acc)]])
        if xs in acc
            continue
        end
        push!(acc, xs)
        if getPath(tree, acc) != -1
            return acc
        end
        deleteat!(acc, length(acc))
    end
    -1
end

tree = getTree(nodes)
function getLink(from, to, nodes)
    for node in nodes
        if node.from == from && node.to == to
            return node
        end
    end
end

function getPoints(from, nodes)
    path = getPath(tree, Int64[from])
    points = scans[from]
    return getPoints(nodes, path, points)
end
function getPoints(nodes, acc, points)
    if length(acc) == 1
        return points
    end
    from = acc[1]
    to = acc[2]
    deleteat!(acc, 1)
    link = getLink(from, to, nodes)

    points = mapslices(x -> (link.rot) * x + (link.trans), points, dims = 2)
    return getPoints(nodes, acc, points)
end




#x1 = getPoints(2, nodes)
#x2 = mapslices((x) -> [-1 0 0; 0 1 0; 0 0 -1] \ x, x1, dims = [2])
#mapslices((x) -> [68, -1246, -43] + x, x2, dims = [2])


nodes = getSequence(scans)

for node in nodes
    println("$(node.from) -> $(node.to) :: $(node.rot) :: $(node.trans)")
end
println("...")
ps = Set()
for i = 1:length(scans)
    x = getPoints(i, nodes)
    for x1 in eachrow(x)
        push!(ps, x1)
    end
end
println(length(ps))


function manhattan(a, b)

end




function getMaxDist(from, nodes)
    path = getPath(tree, Int64[from])
    #start = scans[from].trans
    return getMaxDist(nodes, path, [0; 0; 0])
end
function getMaxDist(nodes, acc, sofar)
    if length(acc) == 1
        return sofar
    end
    from = acc[1]
    to = acc[2]
    deleteat!(acc, 1)
    link = getLink(from, to, nodes)
    sofar = (link.rot) * sofar + (link.trans)

    return getMaxDist(nodes, acc, sofar)
end

scanLocations = []
for i = 1:length(scans)
    x = getMaxDist(i, nodes)
    push!(scanLocations, x)
end
using Distances
z = 0
for a in scanLocations
    for b in scanLocations
        global z = max(z, cityblock(a, b))
    end
end
println(z)
