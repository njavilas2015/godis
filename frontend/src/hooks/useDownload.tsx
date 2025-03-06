const useDownload = () => {

    return (params: { url: string, name: string }) => {

        const link = document.createElement("a");

        link.href = params.url;

        link.download = params.name

        document.body.appendChild(link);

        link.click();

        document.body.removeChild(link);
    }
}

export default useDownload