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
