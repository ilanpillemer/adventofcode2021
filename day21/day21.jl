println("day21")

function init()
    global start = 0
end

function roll()
    global start = start + 3
    [start - 2, start - 1, start]
end

function move(p)
    p = p + sum(roll())
    p = p % 10
    if p == 0
        p = 10
    end
    p
end

#Player 1 starting position: 4
#Player 2 starting position: 3
function play(p1, p2)
    init()
    rolls = 0
    score1 = 0
    score2 = 0
    while true
        p1 = move(p1)
        rolls = rolls + 3
        score1 = score1 + p1
        if score1 >= 1000
            break
        end
        p2 = move(p2)
        rolls = rolls + 3
        score2 = score2 + p2
        if score2 >= 1000
            break
        end
    end
    println("number of rolls: $rolls")
    println("player 1 score: $score1")
    println("player 2 score: $score2")
    println("losing score is: $(min(score1,score2))")
    println("$((min(score1,score2)) * rolls)")
end

memo = Dict()
function play2(p1, p2, s1, s2)
    if s1 >= 21
        return [1, 0]
    end
    if s2 >= 21
        return [0, 1]
    end
    if haskey(memo, (p1, p2, s1, s2))
        return memo[(p1, p2, s1, s2)]
    end
    total = [0, 0]

    for i = 1:3, j = 1:3, k = 1:3
        next_p1 = (p1 + i + j + k) % 10
        if next_p1 == 0
            next_p1 = 10
        end
        local next_s1 = s1 + next_p1
        local score = play2(p2, next_p1, s2, next_s1)
        total = [total[1] + score[2], total[2] + score[1]]
    end
    global memo[(p1, p2, s1, s2)] = total
    return total
end
#Player 1 starting position: 4
#Player 2 starting position: 3
result = play2(4, 3, 0, 0)
max(result[1], result[2])
