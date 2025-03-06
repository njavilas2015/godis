import { create } from 'zustand';

export interface Author {
    login: string;
    id: number;
    node_id: string;
    avatar_url: string;
    gravatar_id: string;
    url: string;
    html_url: string;
    followers_url: string;
    following_url: string;
    gists_url: string;
    starred_url: string;
    subscriptions_url: string;
    organizations_url: string;
    repos_url: string;
    events_url: string;
    received_events_url: string;
    type: string;
    user_view_type: string;
    site_admin: boolean;
}

export interface Uploader extends Author { }

export interface Asset {
    url: string;
    id: number;
    node_id: string;
    name: string;
    label: string | null;
    uploader: Uploader;
    content_type: string;
    state: string;
    size: number;
    download_count: number;
    created_at: string;
    updated_at: string;
    browser_download_url: string;
}

export interface Release {
    url: string;
    assets_url: string;
    upload_url: string;
    html_url: string;
    id: number;
    author: Author;
    node_id: string;
    tag_name: string;
    target_commitish: string;
    name: string;
    draft: boolean;
    prerelease: boolean;
    created_at: string;
    published_at: string;
    assets: Asset[];
    tarball_url: string;
    zipball_url: string;
    body: string;
}

export interface Store {
    data: Release
    fetch: (url: string, signal?: AbortSignal) => Promise<void>
}

async function fetchData(url: string, signal?: AbortSignal): Promise<Release> {

    const response: Response = await fetch(url, {
        method: "GET",
        signal
    })

    return await response.json()
}

const useStore = create<Store>((set) => ({
    data: {
        url: '',
        assets_url: '',
        upload_url: '',
        html_url: '',
        id: 0,
        author: {
            login: '',
            id: 0,
            node_id: '',
            avatar_url: '',
            gravatar_id: '',
            url: '',
            html_url: '',
            followers_url: '',
            following_url: '',
            gists_url: '',
            starred_url: '',
            subscriptions_url: '',
            organizations_url: '',
            repos_url: '',
            events_url: '',
            received_events_url: '',
            type: '',
            user_view_type: '',
            site_admin: false
        },
        node_id: '',
        tag_name: '',
        target_commitish: '',
        name: '',
        draft: false,
        prerelease: false,
        created_at: '',
        published_at: '',
        assets: [],
        tarball_url: '',
        zipball_url: '',
        body: ''
    },
    fetch: async (url: string, signal?: AbortSignal) => {

        const data: Release = await fetchData(url, signal);

        if (data) {
            set({ data });
        }
    },
}));

export default useStore;
