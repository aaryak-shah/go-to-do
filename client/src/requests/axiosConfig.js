import axios from "axios"
const axiosConfig = {
    baseURL: "http://localhost:9000/api/v1/",
}
const instance = axios.create(axiosConfig)
instance.interceptors.request.use(
    function (config) {
        // Do something before request is sent
        config.withCredentials = true
        return config
    },
    function (error) {
        // Do something with request error
        return Promise.reject(error)
    }
)
export default instance