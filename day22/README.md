#Part 1 with structural regex
```
Edit  , s/\.\./:/g
Edit ,x/^off.*\n/ a/cuboids[(x,y,z)]=false\nend\n/
Edit ,x/^on.*\n/ a/cuboids[(x,y,z)]=true\nend\n/
Edit , s/on/for/g
Edit , s/off/for/g
Edit , /^/ i /cuboids=Dict()\n/
Edit $ a /\nsum(values(cuboids))/
```