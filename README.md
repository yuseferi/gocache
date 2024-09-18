# gocache (basic go in-memory caching)
[![codecov](https://codecov.io/github/yuseferi/gocache/graph/badge.svg?token=98CX2MN5XF)](https://codecov.io/github/yuseferi/gocache)
[![CodeQL](https://github.com/yuseferi/gocache/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/yuseferi/gocache/actions/workflows/github-code-scanning/codeql)
[![Check & Build](https://github.com/yuseferi/gocache/actions/workflows/ci.yml/badge.svg)](https://github.com/yuseferi/gocache/actions/workflows/ci.yml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/yuseferi/gocache)

gocache is a data race-free cache implementation in Go, providing efficient caching capabilities for your applications.

### Installation

```shell
  go get -u github.com/yuseferi/gocache
```

### Usage:


```Go
cache := gocache.NewCache(time.Minute * 2) // with 2 minutes interval cleaning expired items
cache.Set("key", "value", time.Minute) // set cache 
value, found := cache.Get("key") // retrive cache data 
cache.Delete("key") // delete specific key manually
cache.Clear() // clear all cache items ( purge)
size := cache.Size() // get cache size
```


### Contributing
We strongly believe in open-source ❤️😊. Please feel free to contribute by raising issues and submitting pull requests to make gocache even better!


Released under the [GNU GENERAL PUBLIC LICENSE](LICENSE).




