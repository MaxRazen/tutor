type eventHandler = (e: Event) => void;
type messageHandler = (e: MessageEvent<any>) => void;
type closeHandler = (e: CloseEvent) => void;
type genericHandler = eventHandler | messageHandler | closeHandler;

export default class WSConnection {
    private conn: WebSocket;
    private onOpenSubscribers: eventHandler[] = [];
    private onErrorSubscribers: eventHandler[] = [];
    private onCloseSubscribers: closeHandler[] = [];
    private onMessageSubscribers: messageHandler[] = [];

    constructor(private endpoint: string) {
    }

    public connect() {
        if (this.conn && this.conn.readyState != WebSocket.CLOSED) {
            return this;
        }

        this.conn = new WebSocket(this.endpoint);
        this.conn.onopen = (e: Event) => {
            this.onOpenSubscribers.forEach((handler: eventHandler) => this.safeCall(handler, e));
        }
        this.conn.onerror = (e: Event) => {
            this.onErrorSubscribers.forEach((handler: eventHandler) => this.safeCall(handler, e));
        }
        this.conn.onclose = (e: CloseEvent) => {
            this.onCloseSubscribers.forEach((handler: closeHandler) => this.safeCall(handler, e));
        }
        this.conn.onmessage = (e: MessageEvent<any>) => {
            this.onMessageSubscribers.forEach((handler: messageHandler) => this.safeCall(handler, e));
        }

        return this;
    }

    public onConnect(func: eventHandler) {
        this.onOpenSubscribers.push(func);
    }

    public onMessage(func: messageHandler) {
        this.onMessageSubscribers.push(func);
    }

    public onError(func: eventHandler) {
        this.onErrorSubscribers.push(func);
    }

    public onClose(func: closeHandler) {
        this.onCloseSubscribers.push(func);
    }

    public disconnect(): void {
        this.conn.close();
    }

    private safeCall(handler: genericHandler, arg: any) {
        try {
            handler(arg);
        } catch (e) {
            console.error(e);
        }
    }
}
