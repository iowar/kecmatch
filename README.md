# Kecmatch

[Use](https://github.com/iowar/kecmatch-gpu) gpu version if you have a nvidia card.

Finds matching solidity function signatures, i.e. method ID.
# Build
~~~sh
$ cd build
$ go build -o kecmatch
~~~ 
### Run
~~~sh
$ ./kecmatch -fn_name "Withdraw" -fn_args "(address[],uint256)" -fn_sig "000000ff" -interval 10
~~~

# Docker
~~~sh
$ docker build . -t kecmatch
~~~
### Run with docker
~~~sh
$ docker run --rm kecmatch -fn_name "SwapFunc" -fn_args "(address,uint256,address)" -fn_sig "00000007"
~~~

# Usage
~~~sh
$ docker run --rm kecmatch -h
~~~

```
  -fn_args string
        # solidity function arguments (default "(address[],uint256)")
  -fn_name string
        # solidity function name (default "DefaultFunc")
  -fn_sig string
        # solidity function method signature (default "00badfee")
  -goroutines int
        # goroutines  (default 8)
  -interval int
        # logger interval (second) (default 3)
```

# Output
```
2022/10/24 22:58:55  SIGNATURE FOUND:

 data: Withdraw50000395536(address[],uint256)
 signature: 000000ff
 iteration: 3181945
 goroutines: 8
 exec time: 894.849036ms
```
# Note
A probability of one in every 4 billion. Finding the wanted signature may take longer than expected.

License
----
[MIT](https://github.com/iowar/kecmatch/blob/main/LICENSE)
