-- 你的验证码在 Redis 上的 Key
-- phone_code:login:185xxxx
local key = KEYS[1]
-- 用户输入的验证码
local expectedCode = ARGV[1]
local key = key..":cnt"

local redisCode = redis.call("get",key)
local cnt = tonumber(redis.call("get",key))

if cnt <= 0 then
    -- 重试次数已经耗尽
    return -1
elseif expectedCode == redisCode then
    -- 校验通过 删除这个key
    redis.call("set",key,-1)
    return 0
else
    -- 验证码不匹配，需要重新输入 需要减少重试次数
    redis.call("decr",key,-1)
    return -2
end
-- 1. 是否存在key

-- 2. key是否过期

-- 3. 重试次数是否大于0

-- 4. 校验值是否相等
