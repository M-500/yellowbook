import api from "@/utils/ajax";

export const getArticlesAPI = (params = {}) => {
    return api.$get("/cms/question", params)
}
/** 获取文章分组对应的分类统计 */
export const getArticlesCountAPI = (params = {}) => {
    return api.$get("/cms/question-count", params)
}

export const addArticleAPI = (params = {}) => {
    return api.$post("/cms/question", params)
}

export const updateArticleAPI = (params = {}) => {
    return api.$post("/cms/question-change", params)
}


export const getCateListAPI = (params = {}) => {
    return api.$get("/cms/categories", params)
}

export const deleteArticle = (params = {}) => {
    return api.$post("/cms/del-question", params)
}

export const starArticle = (params = {}) => {
    return api.$post("/cms/question-star", params)
}

export const cancelStar = (params = {}) => {
    return api.$post("/cms/question-star-cancel", params)
}

export const parseExcel = (params = {}) => {
    return api.$post("/cms/excel-check", params)
}