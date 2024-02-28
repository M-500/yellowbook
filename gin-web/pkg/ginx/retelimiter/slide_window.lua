-- 1, 2, 3, 4, 5, 6, 7 这是你的元素
-- ZREMRANGEBYSCORE key1 0 6
-- 7 执行完之后

-- 限流对象
local key = KEYS[1]
-- 窗口大小
local window = tonumber(ARGV[1])
-- 阈值
local threshold = tonumber( ARGV[2])
local now = tonumber(ARGV[3])
-- 窗口的起始时间
local min = now - window

-- 该命令会删除有序集合中分数在指定范围内的所有成员。如果有序集合不存在，该命令将返回 0；如果指定范围内有成员被成功删除，该命令将返回被删除的成员的数量。
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