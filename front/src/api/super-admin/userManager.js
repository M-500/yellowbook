import api from "@/utils/ajax";

export const getUserListAPI = (params = {}) => {
    return api.$get("/cms/users", params)
}

/** 更新用户信息 */
export const updateUserListAPI = (params = {}) => {
    return api.$get("/cms/users", params)
}


/** 重置用户密码 */
export const resetUserPasswordAPI = (params = {}) => {
    return api.$post("/cms/pwd-reset", params)
}

/** 修改用户的权限范围 */
export const changeUserPermissionAPI = (params = {}) => {
    return api.$post("/cms/user-change-permission", params)
}


/** 删除用户的接口 */
export const delUserListAPI = (params = {}) => {
    return api.$post("/cms/user-stop", params)
}

/** 激活用户的接口 */
export const activeUserListAPI = (params = {}) => {
    return api.$post("/cms/user-active", params)
}

/** 创建用户接口 */
export const createUserListAPI = (params = {}) => {
    return api.$post("/cms/user", params)
}
