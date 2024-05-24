export type RecordingResult = {
    duration: number
    audio: Blob
}

export class AudioRecorder {
    startTime: Date

    audioBlobs: Blob[]

    mediaRecorder: MediaRecorder | null

    streamBeingCaptured: MediaStream | null

    public getDuration(): number {
        return (new Date()).getTime() - this.startTime.getTime()
    }

    public async start(): Promise<MediaStream | void> {
        if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
            throw new Error('mediaDevices API or getUserMedia method is not supported in this browser.');
        }

        return navigator.mediaDevices.getUserMedia({ audio: true })
            .then((stream: MediaStream) => {
                this.streamBeingCaptured = stream;

                this.mediaRecorder = new MediaRecorder(stream);

                this.audioBlobs = [];

                this.mediaRecorder.addEventListener('dataavailable', (event: BlobEvent) => {
                    this.audioBlobs.push(event.data);
                });

                this.mediaRecorder.start();
                this.startTime = new Date();
            });
    }

    public stop(): Promise<RecordingResult> {
        return new Promise(resolve => {
            if (!this.mediaRecorder) {
                return;
            }

            const mimeType = this.mediaRecorder.mimeType;    

            this.mediaRecorder.addEventListener('stop', () => {
                const audio = new Blob(this.audioBlobs, { type: mimeType });
                const duration: number = this.getDuration();

                resolve({
                    duration,
                    audio,
                });
            });

            this.cancel();
        });
    }

    public cancel() {
        if (!this.mediaRecorder) {
            return;
        }

        this.mediaRecorder.stop();

        this.stopStream();

        this.resetRecordingProperties();
    }

    public stopStream() {
        if (!this.streamBeingCaptured) {
            return;
        }

        this.streamBeingCaptured
            .getTracks()
            .forEach((track: MediaStreamTrack) => track.stop());
    }

    public resetRecordingProperties() {
        this.mediaRecorder = null;
        this.streamBeingCaptured = null;
    }
}
