module main

go 1.16

require (
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	local.packages/web v0.0.0
	rsc.io/quote v1.5.2
)

replace local.packages/web => ./web
