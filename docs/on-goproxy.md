## On GOPROXY

The starting time of this article is `2019.05.23`. This time node is very important. I will talk about some of the future of `GOPROXY`.         

- [1.1. Dilemma](#11-dilemma)     
- [1.2. New version of golang](#12-new-version-of-golang)     
- [1.3. Situation](#13-situation)     
- [1.4. Future](#14-future)     

### 1.1. Dilemma

The original reason I developed `GOS` was that I had some trouble with `GOMODULE` and `GOPROXY`.

There is a dilemma here. If I don't use `GOPROXY`, there may be a large number of Package pull timeouts (network reasons) or non-existence (repository rename, delete or migrate), like the following:
```bash
unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: 
dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
```

If use `GOPROXY`, I will not be able to pull the private repositories (github, gitlab, etc) properly, like that:
```
go get github.com/your_private_repo: unexpected status (https://athens.azurefd.net/github.com/your_private_repo/@v/list): 500 Internal Server Error
```

Then I solved these problems in some designs in `GOS`. It feels good after using it for a while, then I plan to open it up and expect it to help more developers.         

### 1.2. New version of golang

But soon after, I noticed that `GOPROXY` was supported on the `master` branch of golang as a comma-separated list of URLs.

You can see it here:       
https://github.com/golang/go/blob/master/src/cmd/go/internal/modfetch/proxy.go#L35
https://tip.golang.org/cmd/go/#hdr-Module_proxy_protocol


```go
var HelpGoproxy = &base.Command{
	UsageLine: "goproxy",
	Short:     "module proxy protocol",
	Long: `
The go command by default downloads modules from version control systems
directly, just as 'go get' always has. The GOPROXY environment variable allows
further control over the download source. If GOPROXY is unset, is the empty string,
or is the string "direct", downloads use the default direct connection to version
control systems. Setting GOPROXY to "off" disallows downloading modules from
any source. Otherwise, GOPROXY is expected to be a comma-separated list of
the URLs of module proxies, in which case the go command will fetch modules
from those proxies. For each request, the go command tries each proxy in sequence,
only moving to the next if the current proxy returns a 404 or 410 HTTP response.
The string "direct" may appear in the proxy list, to cause a direct connection to
be attempted at that point in the search.
...

func proxyURLs() ([]string, error) {
	proxyOnce.Do(func() {
		for _, proxyURL := range strings.Split(proxyURL, ",") {
			if proxyURL == "" {
				continue
			}
			if proxyURL == "direct" {
				proxyOnce.list = append(proxyOnce.list, "direct")
				continue
			}

			// Check that newProxyRepo accepts the URL.
			// It won't do anything with the path.
			_, err := newProxyRepo(proxyURL, "golang.org/x/text")
			if err != nil {
				proxyOnce.err = err
				return
			}
			proxyOnce.list = append(proxyOnce.list, proxyURL)
		}
	})

	return proxyOnce.list, proxyOnce.err
}
```

The cause of these changes is an [issue](https://github.com/golang/go/issues/26334). 


### 1.3. Situation

I am very happy that golang has made a change for this. Now it seems that we can solve the above [dilemma](#dilemma) by setting `GOPROXY=https://GOPROXY-Address1,https://GOPROXY-Address2,direct`. But in fact, the golang version I used when developing gos was `1.12.4`, it still does not support this feature:

```
invalid $GOPROXY setting: cannot have comma
```

What's more, the latest release of golang on https://golang.google.cn/dl/ is `1.12.5`, and this feature is not supported either.          

The document on the master branch of golang/go has such a [sentence](https://github.com/golang/go/blob/master/src/cmd/go/internal/modfetch/proxy.go#L42):

```
For each request, the go command tries each proxy in sequence,
only moving to the next if the current proxy returns a 404 or 410 HTTP response
```

But now most `GORPOXY` services (such as `goproxy.io`, `azure athens`) return `500 Internal Server Error` when they fail to pull private repositories. This means that even though the latest version of the golang distribution supports the `GOPROXY list` feature, `GOPROXY` still does not allow you to pull the private repository correctly for a while.     

That's the situation.       

### 1.4. Future

I would urge `athens` to make a change, even at that time, all GOS could do on `GOPROXY` was to help developers set environment variables.      

But `GOS` can do more than just `GOPROXY`, isn't it?
