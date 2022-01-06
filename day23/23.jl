println("... day 23...")

# 	example
#	01234567890234
#	1#############
#	2#...........#
#	3###B#C#B#D###
#	4  #A#D#C#A#
#	5  #########

amber = [
    3 4
    4 4
]

bronze = [
    3 6
    4 6
]

copper = [
    3 8
    4 8
]

desert = [
    3 10
    4 10
]

home = Dict("A" => amber, "B" => bronze, "C" => copper, "D" => desert)

corridor = [
    2 2
    2 3
    2 4
    2 5
    2 6
    2 7
    2 8
    2 9
    2 10
    2 11
    2 12
    2 13
]
allplaces = [corridor; amber; bronze; copper; desert]

mutable struct Amphipod
    kind::String
    pos::Vector{Int64}
end

# example initial state
pod1 = Amphipod("B", amber[1, :])
pod2 = Amphipod("A", amber[2, :])
pod3 = Amphipod("C", bronze[1, :])
pod4 = Amphipod("D", bronze[2, :])
pod5 = Amphipod("B", copper[1, :])
pod6 = Amphipod("C", copper[2, :])
pod7 = Amphipod("D", desert[1, :])
pod8 = Amphipod("A", desert[2, :])

pods = [pod1; pod2; pod3; pod4; pod5; pod6; pod7; pod8]
amhome(pod) = pod.pos in eachrow(home[pod.kind])
function allhome(pods)
    for pod in pods
        if !(pod.pos in eachrow(home[pod.kind]))
            return false
        end
    end
    true
end

function moves(pod::Amphipod)
    for a in eachrow(allplaces)
        v = vec([a[1] a[2]])
        if islegal(pod, v, pods)
            println("$a is true for $pod")
        end
    end
end

function islegal(pod::Amphipod, move::Vector{Int64}, pods::Vector{Amphipod})
    # cant move to itself
    if pod.pos == move
        return false
    end

    # cant move where someone else is
    for p in pods
        if p != pod && p.pos == move
            return false
        end
    end

    # if a pod is trapped in a side channel
    if pod.pos[1] == 4
        for p in pods
            if p.pos[2] == pod.pos[2] && p.pos[1] == 3
                return false
            end
        end
    end

    # cant move into someone elses home
    for (k, v) in home
        if move in v && k != pos.kind
            return false
        end
        if move in v && k == pos.kind
            # if it is your home, you only move in
            # if no other kinds of pods are there
            for p in pods
                if p.pos in v && p.kind != pod.kind
                    return false
                end
            end
        end
    end

    # if already in corridor cannot move to corridor
    if pod.pos in corridor && move in corridor
        return false
    end

    # if you are moving into the corridor
    if move in eachrow(corridor)
        # you cannot move to an entrance
        if move[2] in [4, 6, 8, 10]
            return false
        end
        # if there are other pods in the corridor

        for p in pods
            if p.pos in eachrow(corridor)
                #if you are moving left
                if move[1] < pod.pos[1]
                    #there must not be pod in the way
                    if p.pos[1] <= pod.pos[1]
                        return false
                    end
                end
                # if you are moving right
                if move[1] > pod.pos[1]
                    #there must not be pod in the way
                    if p.pos[1] >= pod.pos[1]
                        return false
                    end
                end
            end

        end
    end

    return true
end

moves(pod1)
