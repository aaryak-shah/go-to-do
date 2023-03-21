import axios from "./axiosConfig"

export const viewAllToDos = async () => {
    try {
        const url = "/view/"
        const result = await axios.get(url)
        return result.data
    } catch (err) {
        console.error(err)
        throw err
    }
}

export const viewToDoID = async (data) => {
    try {
        const url = `/view/${data}`
        const result = await axios.get(url)
        return result.data
    } catch (err) {
        console.error(err)
        throw err
    }
}