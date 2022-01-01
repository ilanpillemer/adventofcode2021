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

#Part 2 make data easier to read in

```
Edit $ a /\n/
Edit , s/=/=\(/g
Edit , s/,/),/g
Edit , s/\n/)\n/g
Edit , s/\n/)\n/g
Edit , s/on/on(/g
Edit , s/off/off(/g
Edit ,s /.=//g
Edit , s/\.\./,/g
```

#Part 2 with structureal regex

This failed, as too much data.

```
Edit , s/\.\./:/g
Edit , s/x=/A\[/g
Edit , s/y=//g
Edit , s/z=//g
Edit , s/\n/\]\n/g
Edit ,x/^on.*]/ a / .=1/
Edit ,x/^off.*]/ a / .=0/
Edit , s/on//g
Edit , s/off//g
```