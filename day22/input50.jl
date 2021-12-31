cuboids = Dict()
for x = -7:46, y = -33:20, z = -18:35
    cuboids[(x, y, z)] = true
end
for x = -30:15, y = -49:0, z = -45:4
    cuboids[(x, y, z)] = true
end
for x = -8:43, y = -37:11, z = -35:16
    cuboids[(x, y, z)] = true
end
for x = -47:5, y = -29:25, z = -14:40
    cuboids[(x, y, z)] = true
end
for x = -42:9, y = -22:25, z = -19:34
    cuboids[(x, y, z)] = true
end
for x = -2:43, y = 0:48, z = -21:31
    cuboids[(x, y, z)] = true
end
for x = -31:16, y = -37:8, z = -20:33
    cuboids[(x, y, z)] = true
end
for x = -49:-2, y = -6:42, z = -38:8
    cuboids[(x, y, z)] = true
end
for x = -28:20, y = -23:25, z = 0:49
    cuboids[(x, y, z)] = true
end
for x = -38:9, y = -24:27, z = -46:-1
    cuboids[(x, y, z)] = true
end
for x = -26:-17, y = -40:-29, z = 19:30
    cuboids[(x, y, z)] = false
end
for x = -44:10, y = -5:45, z = -1:43
    cuboids[(x, y, z)] = true
end
for x = -7:6, y = -43:-31, z = -40:-21
    cuboids[(x, y, z)] = false
end
for x = -31:19, y = -38:8, z = -13:34
    cuboids[(x, y, z)] = true
end
for x = -43:-33, y = -8:6, z = -35:-21
    cuboids[(x, y, z)] = false
end
for x = -36:18, y = -40:8, z = -5:43
    cuboids[(x, y, z)] = true
end
for x = 3:18, y = -12:7, z = 2:16
    cuboids[(x, y, z)] = false
end
for x = -31:21, y = -47:0, z = -13:37
    cuboids[(x, y, z)] = true
end
for x = 3:20, y = -25:-15, z = 36:47
    cuboids[(x, y, z)] = false
end
for x = -42:2, y = -37:12, z = -34:14
    cuboids[(x, y, z)] = true
end

sum(values(cuboids))
