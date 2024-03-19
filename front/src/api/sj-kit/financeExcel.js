import api from "@/utils/ajax";

export const checkExcelAPI = (params = {}) => {
    return api.$post("/excel/check", params)
}
