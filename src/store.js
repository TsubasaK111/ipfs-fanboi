import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

// const apiUrl = "https://cc5-assessment-halfway-tsubasa.herokuapp.com";
const apiUrl = "http://localhost:4000";
export default new Vuex.Store({
  state: {
    notes: [
      {
        id: 0,
        title: "sample",
        text: "sexesade",
        createdAt: Date.now(),
        updatedAt: Date.now(),
      },
    ]
  },
  mutations: {
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
  },
  actions: {
    getAllNotes(context) {
      fetch(`${apiUrl}/api/notes`)
        .then(response => response.json())
        .then(notes => {
          notes.forEach(newNote => context.commit('addNote', newNote));
        })
        .catch(error => { throw Error(error) });
    },
    addNote(context, newNote) {
      fetch(`${apiUrl}/api/notes/`, {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newNote)
      })
        .then(response => response.json())
        .then(newNote => {
          context.commit('addNote', newNote);
        })
        .catch(error => { throw Error(error) });
    },
    deleteNote(context, deleteNoteId) {
      fetch(`${apiUrl}/api/notes/${deleteNoteId}`,{
        method: 'DELETE'
      })
        .then(response => response.json())
        .then(notes => {
          context.commit('deleteNote', deleteNoteId);
          notes.forEach(note => console.log(note));
        })
        .catch(error => { throw Error(error) });
    },
  },
});
