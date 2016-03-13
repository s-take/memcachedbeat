# memcachedbeat

`memcached-tool localhost:11211 stats` の実行結果をelasticsearchに送るbeat

## todo

* etc/memcachedbeat.template.json を作る
* dashboardを作る
* testを書く

## run

```
go run main.go -c etc/beat.yml
```

### sample

```
{
    "@timestamp" : "2016-03-13T10:20:40.629Z",
    "beat" : {
      "hostname" : "localhost.localdomain",
      "name" : "localhost.localdomain"
    },
    "bytes" : 0,
    "bytes_read" : 91,
    "bytes_written" : 12334,
    "cas_badval" : 0,
    "cas_hits" : 0,
    "cas_misses" : 0,
    "cmd_flush" : 0,
    "cmd_get" : 0,
    "cmd_set" : 0,
    "cmd_touch" : 0,
    "conn_yields" : 0,
    "connection_structures" : 11,
    "curr_connections" : 10,
    "curr_items" : 0,
    "decr_hits" : 0,
    "decr_misses" : 0,
    "delete_hits" : 0,
    "delete_misses" : 0,
    "evicted_unfetched" : 0,
    "evictions" : 0,
    "expired_unfetched" : 0,
    "get_hits" : 0,
    "get_misses" : 0,
    "hash_bytes" : 524288,
    "hash_is_expanding" : 0,
    "hash_power_level" : 16,
    "incr_hits" : 0,
    "incr_misses" : 0,
    "libevent" : 0,
    "limit_maxbytes" : 6.7108864E7,
    "listen_disabled_num" : 0,
    "pid" : 4905,
    "pointer_size" : 64,
    "reclaimed" : 0,
    "reserved_fds" : 20,
    "rusage_system" : 0.058971,
    "rusage_user" : 0.035382,
    "threads" : 4,
    "time" : 1.457863019E9,
    "total_connections" : 23,
    "total_items" : 0,
    "touch_hits" : 0,
    "touch_misses" : 0,
    "type" : "memcachedbeat",
    "uptime" : 1781,
    "version" : 0
  }
}
```
