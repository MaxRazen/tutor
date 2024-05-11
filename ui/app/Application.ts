interface AppConfig {
    title: string
    baseUrl: string
}

declare global {
    interface Window {
        AppConfig: AppConfig
        PageData: any
    }
}

class Application {
    constructor(private config: AppConfig) {
        this.config = config;
    }

    public setTitle(title: string) {
        document.title = `${title} | ${this.config.title}`;
    }

    private apiUrl(path: string) {
        return '/api/v1/' + path;
    }
}

const app = new Application(window.AppConfig);

export default app;
