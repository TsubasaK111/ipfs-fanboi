<template>
  <div id="app">
    <v-app>
      <v-navigation-drawer
        fixed
        v-model="drawer"
        app
      >
        <v-list dense>
          <v-list-tile @click="newNote">
            <v-list-tile-action>
              <v-icon>fiber_new</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>Create New Note</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
          <NavNotesItem
            v-on:select-note="currentNote=note.id"
            v-for="note in this.$store.state.notes"
            v-bind:title="note.title"
            v-bind:key="note.id"
          />
        </v-list>
      </v-navigation-drawer>
      <v-toolbar color="indigo" dark fixed app>
        <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <v-toolbar-title>Code Chrysalis Notes</v-toolbar-title>
      </v-toolbar>
      <v-content>
        <FeedList />
        <v-container fluid fill-height>
          <v-layout
            justify-center
            align-center
          >
            <v-flex text-xs-center>
              <!-- <div>{{this.$store.state.notes.filter(note => note.id === currentNote)}}</div> -->
              <div>{{this.$store.state}}</div>
              <div>:trollface:</div>
            </v-flex>
          </v-layout>
        </v-container>
      </v-content>
      <v-footer color="indigo" app inset>
      </v-footer>    
    </v-app>

  </div>
</template>

<script>
import NavNotesItem from "./components/NavNotesItem.vue";
import FeedList from "./components/FeedList.vue";

export default {
  name: "app",
  data() {
    return {
      drawer: null,
      currentNote: 0
    };
  },
  methods: {
    newNote() {
      console.log("new note requested");
    }
  },
  components: {
    NavNotesItem,
    FeedList,
  },
  created() {
    // this.$store.dispatch("getAllNotes");
    this.$store.dispatch("getAllFeeds");
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
