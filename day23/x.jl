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


board = repeat([Pod(1)], 11)
push!(board, Pod(B))
push!(board, Pod(C))
push!(board, Pod(B))
push!(board, Pod(D))
push!(board, Pod(A))
push!(board, Pod(D))
push!(board, Pod(C))
push!(board, Pod(A))
b = Board(board)

display(b)
