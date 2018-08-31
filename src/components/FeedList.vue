<template>
  <v-list three-line>
    <div 
      v-for="(item, index) in this.$store.state.items" 
      :key="index"
    >
      <!-- <v-subheader
        v-if="item.header"
        :key="item.header"
      >
        {{ item.header }}
      </v-subheader>

      <v-divider
        v-else-if="item.divider"
        :inset="item.inset"
        :key="index"
      ></v-divider> -->

      <v-list-tile
        :key="item.Name"
        avatar
        @click="openItem(item.Url)"
      >
        <v-list-tile-avatar>
          <!-- <img :src="item.avatar"> -->
          <!-- <img :src="'/static/img/clear.gif'"> -->
          {{item.Source}}
        </v-list-tile-avatar>

        <v-list-tile-content>
          <v-list-tile-title v-html="item.Name"></v-list-tile-title>
          <v-list-tile-sub-title v-html="item.Description"></v-list-tile-sub-title>
        </v-list-tile-content>

        <v-list-tile-action>
          <v-list-tile-action-text>{{ item.CreatedAt }}</v-list-tile-action-text>
        </v-list-tile-action>

      </v-list-tile>
    </div>
  </v-list>
</template>

<script>
export default {
  name: "NavNotesItem",
  methods: {
    openItem(url) {
      const win = window.open(url, "_blank");
      win.focus();
    }
  }
};
</script>