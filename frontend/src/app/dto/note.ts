export class Note {
    noteId: string;
    userId: string;
    title: string;
    description: string;

    constructor(noteId: string, userId: string, title: string, description: string) {
        this.noteId = noteId;
        this.userId = userId;
        this.title = title;
        this.description = description;
    }
}
