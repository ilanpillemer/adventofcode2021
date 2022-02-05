# inspired by https://youtu.be/dMSqVAOW9Tw
struct Pod
    v::Int8
end
struct Board
    b::Vector{Pod}
end

Empty = Pod(1)
A = Pod(2)
B = Pod(3)
C = Pod(4)
D = Pod(5)
Names = [:. :A :B :C :D]
EmptyBoard = """
#############
#-----------#
###-#-#-#-###
  #-#-#-#-#
  #########
"""
Base.show(io::IO, z::Pod) = print(io, Names[z.v])

function display(b::Board)
    board = EmptyBoard
    for p in b.b
        board = (replace(board, "-" => string(p), count = 1))
    end
    println(board)
end

function exampleBoard()
    board = fill(Empty, 11 + 4 + 4)
    board[12:15] = [B, C, B, D]
    board[16:19] = [A, D, C, A]
    board
end

function input()
    board = fill(Empty, 11 + 4 + 4)
    board[12:15] = [B, B, D, D]
    board[16:19] = [C, A, A, C]
    board
end

function win()
    board = fill(Empty, 11 + 4 + 4)
    board[12:15] = [A, B, C, D]
    board[16:19] = [A, B, C, D]
    board
end


cost = fill(0, D.v)
cost[A.v] = 1
cost[B.v] = 10
cost[C.v] = 100
cost[D.v] = 1000

entry = fill(0, D.v)
entry[A.v] = 3
entry[B.v] = 5
entry[C.v] = 7
entry[D.v] = 9

up = fill(0, 11 + 4 + 4)
up[12] = 3
up[13] = 5
up[14] = 7
up[15] = 9
up[16] = 12
up[17] = 13
up[18] = 14
up[19] = 15

down = fill(0, 11 + 4 + 4)
for (i, v) in enumerate(up)
    if v != 0
        down[v] = i
    end
end


function move(b::Board, c::Int, f)
    for (i, p) in enumerate(b.b)
        if p == Empty
            continue
        end

        if i >= 12
            # in a room
            # should move up to hallway
            d = 0
            e = i
            #can I go up?
            blocked = false
            while up[e] != 0
                d = d + 1
                e = up[e]
                if b.b[e] != Empty
                    blocked = true
                end
            end
            if blocked
                # julia does not seem to have continue to a labelled for
                continue
            end
            #Left
            for j = (e-1):-1:1
                if b.b[j] != Empty
                    break # blocked no passage
                end
                if down[j] != 0
                    continue # cannot stop at an entrance
                end
                # at a legal position in the hallway
                b2 = deepcopy(b)
                b2.b[i] = Empty
                b2.b[j] = p
                #f(b2)
            end
            for j = e:11
                if b.b[j] != Empty
                    break # blocked no passage
                end
                if down[j] != 0
                    continue # cannot stop at an entrance
                end
                # at a legal position in the hallway
                b2 = deepcopy(b)
                b2.b[i] = Empty
                b2.b[j] = p
                f(b2)
            end

        else
            # in hallway
        end
    end
end

b = Board(exampleBoard())
display(b)
#display(Board(input()))

move(b, 0, display)
