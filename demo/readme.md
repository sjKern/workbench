https://direnv.net/
direnv is in use with

.envrc
```
goversion=1.22.6
eval "$(gvm $goversion)"
export GO11MODULE=on
```
direnv allow

https://github.com/andrewkroh/gvm

gvm golang versions-manager
```
gmv list
gvm use
gvm available
```

ex:
gvm use 1.22.6

```
export GOROOT="/Users/skern/.gvm/versions/go1.22.6.darwin.arm64"
export ```PATH="/Users/skern/.gvm/versions/go1.22.6.darwin.arm64/bin:$PATH"