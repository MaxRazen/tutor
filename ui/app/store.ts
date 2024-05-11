import {User} from './types.d';

const VERSION = '0.0.1';
const AUTH_TOKEN = '_authtoken';
const USER_DATA = '_userdata';

class Store {
    private storage: Storage;

    constructor() {
        this.storage = localStorage;
        this.checkVersion();
    }

    public setAccessToken(accessToken: string): void {
        this.storage.setItem(AUTH_TOKEN, accessToken);
    }

    public getAccessToken(): string | null {
        return this.storage.getItem(AUTH_TOKEN);
    }

    public setUser(user: User): void {
        this.storage.setItem(USER_DATA, JSON.stringify(user));
    }

    public getUser(): User | null {
        const userData: string | null = this.storage.getItem(USER_DATA);
        if (userData) {
            return JSON.parse(userData);
        }
        return null;
    }

    private checkVersion(): void {
        const key = '_version';
        const v = this.storage.getItem(key);
        if (v === VERSION) {
            return;
        }
        this.storage.clear();
        this.storage.setItem(key, VERSION);
    }
}

const store = new Store();

export default store;
