# Shikata ga nai on html

Bringing Shikata ga nai to the front html

## Try

```shell
git clone github.com/akkuman/sgn-html
cd sgn-html
GO111MODULE=on go get -v github.com/projectdiscovery/simplehttpserver/cmd/simplehttpserver
simplehttpserver -listen 127.0.0.1:8000
```

If you want to try it right now, please see live demo: [http://akkuman.github.io/sgn-html/](http://akkuman.github.io/sgn-html/)

## Implementation

Use of the original project for https://github.com/EgeBalci/sgn, but the project relies on cgo, a c library keystone, rely on cgo cannot compile wasm, for bridge method

keystone is compiled into wasm via Emscripten

Instead of calling cgo from sgn, keystone wasm is called from syscall/js

Now sgn becomes a pure go project and exposes the interface to js and compiles to wasm

more details: [github.com/akkuman/sgn](https://github.com/akkuman/sgn)

## Reference

- [github.com/EgeBalci/sgn](https://github.com/EgeBalci/sgn)
- [github.com/AlexAltea/keystone.js](https://github.com/AlexAltea/keystone.js)
