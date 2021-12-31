# tidy up first line in sample

```
Edit /\n/ d
```

# After adding 'img[' and ']' after the end of the image lines

```
Edit /img/+,/]/- x/.*\n/i/"/
Edit /img/+,/]/- x/.*$/a/",/
Edit , x s/\./0/g
Edit , x s/#/1/g
```