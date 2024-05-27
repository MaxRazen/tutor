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

type User = {
    id: number
    name: string
    email: string
    avatar: string
    socialId: string
    lastLoggedAt: string
};

class Application {
    private config: AppConfig;
   
    constructor(config: AppConfig) {
        this.config = config;
    }

    public setTitle(title: string) {
        document.title = `${title} | ${this.config.title}`;
    }

    public async createRoom({mode}): Promise<Number> {
        const response = await fetch(this.apiUrl('room'), {
            method: 'POST',
            body: JSON.stringify({mode}),
            headers: new Headers({
                'content-type': 'application/json',
                'accept': 'application/json',
            })
        });
        if (!response.ok) {
            return new Promise(() => {
                alert('Sorry, you cannot open a new room at the moment. Please try later');
                return null;
            });
        }
        const content = await response.json();

        return content.roomId;
    }

    private apiUrl(path: string) {
        return '/api/v1/' + path;
    }
}

const app = new Application(window.AppConfig);

export const TUTOR_AVATAR = 'https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg';
export const TUTOR_NAME = 'Scarlett';

export default app;
