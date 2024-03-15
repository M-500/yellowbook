import api from "@/utils/ajax";


/**
 *  获取后台所有规则描述列表 => 用户table展示
 * */
export const getAllRuleListAPI = (params = {}) => {
    return api.$get("/tax-rule-all", params)
}
