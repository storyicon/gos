## What's new in `v1.1`

When using GOS for CI in a production environment, we encountered some new requirements, so we made some changes!
If you have any new ideas, let us know!

- [1. Custom upstream proxy address](#011-custom-upstream-proxy-address)
- [2. Debug Mode](#012-debug-mode)
- [3. Correct ExitCode](#013-correct-exitcode)

### 1. Custom upstream proxy address

Now, you can specify which upstream proxy `GOS` uses through the environment variable `GOS_UPSTREAM_ADDRESS`.

At present, the default upstream address of `GOS` is `https://athens.azurefd.net`, but for some Chinese users, it seems that the experience is still not good enough.

The most popular golang proxies are the following:

name | address | Speed in China | describtion    
- | - | - | -
goproxy.io | https://goproxy.io | `0.9` | Fast in China
gocenter.io | https://gocenter.io | `0.85` | You may encounter many 404 
goproxy.cn | https://goproxy.cn | `0.8` | Fast in China
athens | https://athens.azurefd.net | `0.5` | Microsoft Azure engineers provided it
proxy.golang.org | https://proxy.golang.org | `0` | From Google

So now you can make `GOS` work better for you by setting the environment variable `GOS_UPSTREAM_ADDRESS`, Like the following:

```
GOS_UPSTREAM_ADDRESS=https://goproxy.io
```

These goproxys bring us a lot of convenience, but you should still pay attention to the risks of using third-party proxies.

### 2. Debug Mode

Now you can let `GOS` output more information at runtime by setting `GOS_DEBUG` to any non-null value.
Like the following:
```
GOS_DEBUG=1
```

It will output:
1. The upstream proxy address currently in use
2. The listening port of the local proxy
3. Detailed process and error message when pulling the repository
...

### 3. Correct ExitCode

Now gos returns a `non-zero` code when it encounters a `critical` error, which makes it perform better when performing `CI` usage.


