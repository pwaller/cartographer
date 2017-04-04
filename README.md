# Cartographer points you at the maps

In particular, cartographer lists all source locations in a package which are loops over maps. If you're not careful, these may cause your program to be non-deterministic.

# Example of use

```
go get -u github.com/pwaller/cartographer
```

```
$ cartographer github.com/pwaller/cartographer/example
/home/pwaller/.local/src/github.com/pwaller/cartographer/example/main.go:7:2
/home/pwaller/.local/src/github.com/pwaller/cartographer/example/main.go:10:2
/home/pwaller/.local/src/github.com/pwaller/cartographer/example/main.go:14:2
/home/pwaller/.local/src/github.com/pwaller/cartographer/example/main.go:17:2
/home/pwaller/.local/src/github.com/pwaller/cartographer/example/main.go:20:2
```

## License

3-clause BSD-style license.

## Trophies

(Please submit more if you have any! :)

* [c-for-go#27](https://github.com/xlab/c-for-go/pull/27)

  A >7,000 line code base was emitting things in an arbitrary order, which
  made diffing the output difficult.

## Thanks

To [@JohannesEbke](https://github.com/JohannesEbke) for hacking on the earliest
prototype together.
