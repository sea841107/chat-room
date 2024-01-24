<template>
  <div id="app">
    <router-view @get="httpGet" @post="httpPost"
    :loginRecall="loginRecall"
    :messageRecall="messageRecall"
    :historyRecall="historyRecall"
    :searchRecall="searchRecall"
    :updateRecall="updateRecall"
    />
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      loginRecall: null,
      messageRecall: null,
      historyRecall: null,
      searchRecall: null,
      updateRecall: null,
    }
  },

  methods: {
    onMessage(r) {
      switch (r.data.header) {
        case "LoginRecall":
          this.loginRecall = r.data;
          break;
        case "MessageRecall":
          this.messageRecall = r.data;
          break;
        case "HistoryRecall":
          this.historyRecall = r.data;
          break;
        case "SearchRecall":
          this.searchRecall = r.data;
          break;
        case "UpdateRecall":
          this.updateRecall = r.data;
          break;
        default:
          break;
      }
    },

    httpGet(...value) {
      var self = this;
      axios.get(`https://chat-room-server-dd.web.app/chat/data/${value}`)
        .then(function (response) {
          self.onMessage(response);
        })
        .catch(function (error) {
          console.log(error);
      });
    },

    httpPost(data) {
      var self = this;
      axios.post('https://chat-room-server-dd.web.app/chat/data', data)
        .then(function (response) {
          self.onMessage(response);
        })
        .catch(function (error) {
          console.log(error);
      });
    }
  }
}
</script>
