
import api from '@/utils/ajax'

export const imageCaptchaAPI = (params = {}) => {
    return api.$get("/captcha", params)
}


export const pwdLoginAPI = (params = {}) => {
    return api.$post("/users/pwd-login", params)
}
