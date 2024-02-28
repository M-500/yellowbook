
-- 你的验证码在 Redis 上的 Key
-- phone_code:login:185xxxx
local key = KEYS[1]
-- 验证次数，一个验证码 校验三次失败，无法再用

local cntKey = key..":cnt"

-- 你的验证码
local val = ARGV[1]
-- 过期时间 在 Redis 中，TTL 命令用于获取指定键的剩余生存时间（Time To Live）。它返回一个以秒为单位的整数值，表示指定键的剩余生存时间，即该键还有多少秒过期。
local ttl = tonumber(redis.call("ttl",key))
if ttl == -1 then
    -- key 存在 但是没有过期时间，可能是技术人为处理的
    return -2
elseif ttl== -2 or ttl < 540 then
    -- key 不存在，那么就设置这个key 并设置过期时间
    redis.call("set", key,val)
    redis.call("expire",key,600)
    redis.call("set",cntKey,3)
    redis.call("expire",cntKey,600)
    return 0
else
    -- 发送太频繁
    return -1
end