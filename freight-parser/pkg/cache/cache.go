package cache

type ICache interface {
	get(key string) interface{}
	set(ket, value string, ttl int64)
}
