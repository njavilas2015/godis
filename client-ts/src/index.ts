export class HashStorage {
    url: string

    constructor(url: string) {
        this.url = url;
    }

    async hset(key: string, fieldValues: Record<string, string>): Promise<string> {

        const args: string[] = [key];

        for (const [field, value] of Object.entries(fieldValues)) {
            args.push(field, value);
        }

        const response = await fetch(`${this.url}/hset`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ args }),
        });

        if (!response.ok) {
            throw new Error(`Failed to execute HSET: ${response.statusText}`);
        }

        return await response.text();
    }

    async hget(key: string, field: string): Promise<string> {

        const args = [key, field];

        const response = await fetch(`${this.url}/hget`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ args }),
        });

        if (!response.ok) {
            throw new Error(`Failed to execute HGET: ${response.statusText}`);
        }

        return await response.text();
    }
}