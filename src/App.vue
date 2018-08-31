<template>
  <div id="app">
    <v-app>
      <v-navigation-drawer
        fixed
        v-model="drawer"
        app
      >
        <v-list dense>
          <NavNotesItem
            v-on:select-note="currentNote=note.id"
            v-for="note in this.$store.state.notes"
            v-bind:title="note.title"
            v-bind:key="note.id"
          />
        </v-list>
      </v-navigation-drawer>
      <v-toolbar color="indigo" dark fixed app>
        <!-- <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon> -->
        <v-toolbar-title>IPFS Fanboi </v-toolbar-title>
      </v-toolbar>
      <v-content>
        <h1>Welcome to IPFS Fanboi!</h1>
        <div>Below is an aggregated feedlist of <br> all the activity that is taking place regarding IPFS:</div>
        <FeedList />
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
