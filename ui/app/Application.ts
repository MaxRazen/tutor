interface AppConfig {
    baseUrl: string
}

declare global {
    interface Window {
        AppConfig: AppConfig
    }
}

class Application {
    constructor(private config: AppConfig) {
        this.config = config;
    }

    private apiUrl(path: string) {
        return '/api/v1/' + path;
    }
}

const app = new Application(window.AppConfig);

export default app;
