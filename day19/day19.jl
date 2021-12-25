include("rots.jl")
println("day 19")
include("example.jl")

origin = exScans[1]

for scan in 2:length(exScans)
for (i,rot) in enumerate(rall)
         println("rot",i)
         found = 0
         for p in eachrow(exScans[scan])
           #println("...",rot*p,typeof(rot*p))
           x1 = rot * p
           x1 = x1 + [68, -1246, -43]
           
           for row in 1:div(length(origin),3)
           
           if x1 == origin[row,:]
              found = found + 1
              println("scanner",scan)
              println("translated",x1)
              println("rotated",rot*p)
              println("scanned",p)
           end
        
           end
           
         end
         if found != 0; println(found); end
end
end
