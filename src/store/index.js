import Vue from "vue";
import Vuex from "vuex";

import {mutations} from "./mutations";
import {actions} from "./actions";

Vue.use(Vuex);

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
    ],
    feeds: [
      {
        id: 0,
        title: "all",
        text: "all the things",
      },
      {
        id: 1,
        title: "github",
        text: "github repos",
      },
      {
        id: 2,
        title: "reddit",
        text: "all subreddits",
      },
      {
        id: 3,
        title: "twitter",
        text: "any and all tweets",
      },
    ],
    items: []
  },
  mutations: mutations,
  actions: actions,
});
