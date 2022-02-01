struct Pod
    v::Int8
end
struct Board
    b::Vector{Pod}
end

A = 2
B = 3
C = 4
D = 5
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
    board = fill(Pod(1), 11 + 4 + 4)
    board[12] = Pod(B)
    board[13] = Pod(C)
    board[14] = Pod(B)
    board[15] = Pod(D)
    board[16] = Pod(A)
    board[17] = Pod(D)
    board[18] = Pod(C)
    board[19] = Pod(A)
    board
end

b = Board(exampleBoard())
display(b)
