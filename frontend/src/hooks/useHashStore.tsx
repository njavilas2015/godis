const useHashStore = () => {

    const url = ""

    const set = async (signal?: AbortSignal): Promise<string> => {

        const response: Response = await fetch(url, {
            method: "POST",
            signal
        })

        return await response.json()
    }

    const get = async (signal?: AbortSignal): Promise<string> => {

        const response: Response = await fetch(url, {
            method: "GET",
            signal
        })

        return await response.json()
    }

    const dropAll = async (signal?: AbortSignal): Promise<string> => {

        const response: Response = await fetch(`${url}/all`, {
            method: "DELETE",
            signal
        })

        return await response.json()
    }

    const drop = async (signal?: AbortSignal): Promise<string> => {

        const response: Response = await fetch(url, {
            method: "DELETE",
            signal
        })

        return await response.json()
    }

    const all = async (signal?: AbortSignal): Promise<string[]> => {

        const response: Response = await fetch(url, {
            method: "GET",
            signal
        })

        return await response.json()
    }

    return { set, get, dropAll, drop, all }
}

export default useHashStore