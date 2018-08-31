<template>
  <v-list two-line>
    <div v-for="(item, index) in this.$store.state.items">
      <v-subheader
        v-if="item.Name"
        :key="item.Name"
      >
        {{ item.Name }}
      </v-subheader>

      <v-divider
        v-else-if="item.divider"
        :inset="item.inset"
        :key="index"
      ></v-divider>

      <v-list-tile
        v-else
        :key="item.title"
        avatar
        @click=""
      >
        <v-list-tile-avatar>
          <img :src="item.avatar">
        </v-list-tile-avatar>

        <v-list-tile-content>
          <v-list-tile-title v-html="item.title"></v-list-tile-title>
          <v-list-tile-sub-title v-html="item.subtitle"></v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
    </div>
  </v-list>
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