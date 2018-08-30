<template>
  <!-- <v-list-tile> -->
  <v-list-tile @click="$emit('select-note')">
    <v-list-tile-action>
      <v-icon>note_add</v-icon>
    </v-list-tile-action>
    <v-list-tile-content>
      <v-list-tile-title>{{title}}</v-list-tile-title>
    </v-list-tile-content>
  </v-list-tile>
</template>

<script>
export default {
  name: 'NavNotesItem',
  props: {
    title: String,
    selectNote: Function,
    id: Number,
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
