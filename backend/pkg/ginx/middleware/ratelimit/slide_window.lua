-- 1, 2, 3, 4, 5, 6, 7 这是你的元素
-- ZREMRANGEBYSCORE key1 0 6
-- 7 执行完之后

-- 限流对象 在Lua脚本中，KEYS是一个全局变量，它通常用于与Redis的Lua脚本交互。在Redis中，KEYS是一个用于存储键（Key）的数组，而KEYS[1]表示获取这个数组的第一个元素，即第一个键。
local key = KEYS[1]
-- 窗口大小   ARGV 是一个全局变量，用于与 Redis 的 Lua 脚本交互。与 KEYS 变量类似，ARGV 是一个数组，用于存储从 Redis 传递给 Lua 脚本的命令参数。
local window = tonumber(ARGV[1])
-- 阈值
local threshold = tonumber( ARGV[2])
local now = tonumber(ARGV[3])
-- 窗口的起始时间
local min = now - window

redis.call('ZREMRANGEBYSCORE', key, '-inf', min)
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf')
-- local cnt = redis.call('ZCOUNT', key, min, '+inf')
if cnt >= threshold then
    -- 执行限流
    return "true"
else
    -- 把 score 和 member 都设置成 now
    redis.call('ZADD', key, now, now)
    redis.call('PEXPIRE', key, window)
    return "false"
end