# spw
Simple Password Generator

## help
```
Usage of spw:
  -c	do not wipe password in clipboard after 60s
  -n int
    	length of password (default 32)
  -p	print the password, in addition to copying to clipboard)
  -w	without special characters
```

## examples

```bash
$ spw -p
-8n[z_DkzO;x4byPI4i]>sO!i1hjH}~_

$ spw -p -w
9BX96PrVSSXJAZYVAdChHy6Ygt76U368

$ spw -p -n 50
DJM.g}GN5zj.0##Ppwug'&SCeoE8Bmr/~}y7zJ9S1K!s%Gt/6E

$ spw
NOTE: Generated password sent to clipboard.
NOTE: Clipboard content to be removed in 60s.
NOTE: Clipboard content removed.
```
