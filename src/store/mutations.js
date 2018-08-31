export const mutations = {
  addItem(state, newItem) {
    state.items.push(newItem);
  },
  addNote(state, newNote) {
    state.notes.push(newNote);
  },
  deleteNote(state, deleteNoteId) {
    state = state.notes.filter(note => note.id !== deleteNoteId);
  },
  updateNote(state, updateNoteId, updatedNote) {
    newNotes = state.notes.map(note => {
      if (note.id === updateNoteId) {
        return { ...note, ...newNote };
      } else {
        return note;
      }
    });

    state.notes = newNotes;
  }
}