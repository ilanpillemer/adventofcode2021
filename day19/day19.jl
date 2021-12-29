include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]

struct Link
    from::Any
    to::Any
    rot::Any
    left::Any
    right::Any
end

struct DecoratedLink
    link::Link
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
            c = Dict([((r * z), b[z]) for z in keys(b)])
            result = intersect(keys(a), keys(c))
            c = filter(p -> first(p) in result, c)
            d = filter(p -> first(p) in result, a)
            #println(length(result))
            if length(result) == 11
                #println(result)
                return (r, true, d, b, c)
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
            translation = getTranslation(link)
            push!(connects, DecoratedLink(link, translation))
        end
    end
    connects
end


nodes = getSequence(exScans)

for node in nodes
    println("$(node.link.from) -> $(node.link.to) :: $(node.link.rot) :: $(node.trans)")
end

function getTree(nodes)
    d = Dict()
    for node in nodes
        xs = get(d, node.link.from, [])
        d[node.link.from] = xs
        d[node.link.from] = push!(d[node.link.from], node.link.to)
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
        if node.link.from == from && node.link.to == to
            return node
        end
    end
end

function getPoints(from, nodes)
    path = getPath(tree, Int64[from])
    points = exScans[from]
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
    println("from $(link.link.from)")
    println("to $(link.link.to)")
    println("rot $(link.link.rot)")
    println("trans $(link.trans)")

    return getPoints(nodes, acc, points)
end


x1 = getPoints(2, nodes)
x2 = mapslices((x) -> [-1 0 0; 0 1 0; 0 0 -1] \ x, x1, dims = [2])
mapslices((x) -> [68, -1246, -43] + x, x2, dims = [2])
