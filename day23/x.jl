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

b = Board(exampleBoard())
display(b)
