export default class AudioRecorder {
        audioBlobs: Blob[]

        mediaRecorder: MediaRecorder | null

        streamBeingCaptured: MediaStream | null

        public async start(): Promise<MediaStream | void> {
            if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
                throw new Error('mediaDevices API or getUserMedia method is not supported in this browser.');
            }
    
            return navigator.mediaDevices.getUserMedia({ audio: true })
                .then((stream: MediaStream) => {
                    //save the reference of the stream to be able to stop it when necessary
                    this.streamBeingCaptured = stream;

                    this.mediaRecorder = new MediaRecorder(stream);

                    this.audioBlobs = [];

                    this.mediaRecorder.addEventListener('dataavailable', (event: BlobEvent) => {
                        this.audioBlobs.push(event.data);
                    });

                    this.mediaRecorder.start();
                });
        }

        public stop(): Promise<Blob> {
            return new Promise(resolve => {
                if (!this.mediaRecorder) {
                    return;
                }

                const mimeType = this.mediaRecorder.mimeType;    

                this.mediaRecorder.addEventListener('stop', () => {
                    const audioBlob = new Blob(this.audioBlobs, { type: mimeType });
                    resolve(audioBlob);
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
                .forEach((track: MediaStreamTrack) => track.stop()); //stop each one
        }

        public resetRecordingProperties() {
            this.mediaRecorder = null;
            this.streamBeingCaptured = null;
        }
}
