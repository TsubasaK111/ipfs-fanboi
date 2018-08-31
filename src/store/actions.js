// const apiUrl = "https://cc5-assessment-halfway-tsubasa.herokuapp.com";

// if (!process.env.PORT) throw Error ('please define a $PORT.');
// const apiUrl = process.env.HEROKU_APP_NAME ? `https://${process.env.HEROKU_APP_NAME}.herokuapp.com` : `http://localhost:${process.env.PORT}`;
// const apiUrl = `http://localhost:8080`;
const apiUrl = `https://ipfs-fanboi.herokuapp.com`

export const actions = {
  getAllFeeds(context) {
    console.log('getting all feeds...')

    fetch(`${apiUrl}/api/feeds`)
      .then(response => {
        console.log(response)
        return response
      })
      .then(response => response.json())
      .then(items => {
        items.forEach(newItem => context.commit('addItem', newItem));
      })
      .catch(error => { throw Error(error) });
  },
  getAllNotes(context) {
    fetch(`${apiUrl}/api/notes`)
      .then(response => {
        console.log(response)
        return response
      })
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
    fetch(`${apiUrl}/api/notes/${deleteNoteId}`, {
      method: 'DELETE'
    })
      .then(response => response.json())
      .then(notes => {
        context.commit('deleteNote', deleteNoteId);
        notes.forEach(note => { console.log(note) });
      })
      .catch(error => { throw Error(error) });
  },
}