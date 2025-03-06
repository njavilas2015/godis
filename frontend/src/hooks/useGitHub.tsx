import { useEffect } from "react"
import useGitHubStore from '@store/useGitHubStore'

const useGitHub = () => {

    const github = useGitHubStore()

    useEffect(() => {

        const url: string = "https://api.github.com/repos/njavilas2015/godis/releases/latest"

        const controller: AbortController = new AbortController()

        github.fetch(url, controller.signal)

        return () => {

            controller.abort()
        }

    }, [])
}

export default useGitHub