## Run ##

Using [godo](https://github.com/anacrolix/godo):

```bash
export GO15VENDOREXPERIMENT=1
godo github.com/Gaboose/examples/Go/nested-vendors
```

Output:

```
From root:
I see AN APPLE - it's YELLOW and it tastes (can't import flavor)

From 'iAlso':
I see AN APPLE - it's RED and it tastes SWEET
```

## Package Structure ##

`GetFruit()` in `fruit` returns *AN APPLE* and is only in the root vendor directory.

`GetColor()` in `color` returns *YELLOW* in the root vendor directory and
*RED* in the `i-also-use-those-pkgs` vendor directory.

`GetFlavor()` in `flavor` returns *SWEET* and is only in the `i-also-use-those-pkgs`
vendor directory. 
