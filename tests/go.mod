module github.com/szpnygo/rcounter-redis/test

go 1.15

replace github.com/szpnygo/rcounter-redis v0.0.0 => ../

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/szpnygo/rcounter-redis v0.0.0
)
