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
