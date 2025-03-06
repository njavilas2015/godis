import { DBSchema, IDBPDatabase, openDB } from "idb";

const dbName: string = 'godis';

const dbVersion: number = 1;

interface IProfile {

}

export type { IDBPDatabase }

export interface BoxerDB {
    profile: {
        key: string
        value: IProfile
    }

    /*products: {
        value: {
            name: string;
            price: number;
            productCode: string;
        };
        key: string;
        indexes: { 'by-price': number };
    };*/
}

export type Schema = { [K in keyof BoxerDB]: BoxerDB[K]; } & DBSchema;

export async function openDatabase(): Promise<IDBPDatabase<Schema>> {

    const db: IDBPDatabase<Schema> = await openDB<Schema>(dbName, dbVersion, {

        upgrade(db) {

            const tables: (keyof BoxerDB)[] = [
               
            ];

            for (const x of tables) {

                if (!db.objectStoreNames.contains(x)) {

                    db.createObjectStore(x, { keyPath: 'id' });
                }
            }
        },
    });

    return db;
};