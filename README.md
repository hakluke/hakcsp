# hakcsp
Return domains in CSP headers in http response. This might be used to uncover more domains that are affiliated with a target.

# Installation

```
go get github.com/hakluke/hakcsp
```

# Usage

Pipe URLs into the tool. You can use `-t` to set the number of goroutines.

```
echo urls | hakcsp -t 8
```
