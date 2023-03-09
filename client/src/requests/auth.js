import axios from "./axiosConfig"

export const register = async (data) => {
    const url = "/auth/register"
    try {
        const result = await axios.post(url, data)
        return result.data
    } catch (err) {
        console.error(err)
        throw err
    }
}

export const login = async (data) => {
    const url = "/auth/login"
    try {
        const result = await axios.post(url, data)
        return result.data
    } catch (err) {
        console.error(err)
        throw err
    }
}

export const logout = async () => {
    const url = "/auth/logout"
    try {
        await axios.get(url)
    } catch (err) {
        console.error(err)
        throw err
    }
}