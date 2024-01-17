<template>
  <div class="chat">
    <div class="chat-list">
      <chat-list :toName="toName" :userList="userList" :hintList="hintList" @select="selectUser" @search="searchCall"></chat-list>
    </div>
    <div class="chat-content">
      <h1 style="background-color:#efe9e7; text-align: center;">{{ toName }}</h1>
      <chat-content :userName="userName" :messageList="messageList" @history="history"></chat-content>
      <input class="chat-text" type="text" placeholder="Input your message" v-model="text" @keydown.enter="messageCall"/>
    </div>
  </div>
</template>

<script>
import ChatList from '../components/ChatList.vue'
import ChatContent from '../components/ChatContent.vue'

export default {
  name: 'ChatView',
  components: {
    ChatList,
    ChatContent,
  },

  data() {
    return {
      text: '',
      toName: "Lobby",
      hintList: [],
      userList: ["Lobby"],
      messageList: [],
      isSearching : false,
      pollTimeout : null,
    }
  },

  props: {
    userName: String,
    messageRecall: null,
    historyRecall: null,
    searchRecall: null,
    updateRecall: null,
  },

  mounted() {
    if (this.userName == "") {
      this.$router.push("/");
    } else {
      this.searchCall("");
      this.historyCall();
      this.polling();
    }
  },

  methods: {
    messageCall() {
      if (!this.text.trim().length) {
        return;
      }

      var data = JSON.stringify({
        header: "MessageCall",
        to: this.toName,
        message: {
          name: this.userName,
          text: this.text,
          time: new Date().getTime()
        }
      });

      this.$emit("post", data)
      this.text = '';  
    },

    updateCall() {
      var data = JSON.stringify({
        header: "UpdateCall",
        from: this.userName,
        to: this.toName,
      });
      this.$emit("get", data);
    },

    historyCall(count) {
      var data = JSON.stringify({
        header: "HistoryCall",
        from: this.userName,
        to: this.toName,
        count: count,
      });
      this.$emit("get", data);
    },

    searchCall(name) {
      var data = JSON.stringify({
        header: "SearchCall",
        from: this.userName,
        name: name,
      });
      this.$emit("get", data);
      this.isSearching = name != ""
    },

    selectUser(user) {
      this.toName = user;
      this.historyCall(0);
    },

    history() {
      if (this.messageList != null) {
        var count = this.messageList.length;
        this.historyCall(count);
      }
    },

    polling() {
      if (this.pollTimeout != null) {
        clearTimeout(this.pollTimeout);
        this.pollTimeout = null;
      }

      this.pollTimeout = setTimeout(() => {
        this.polling();
      }, 300000);

      this.updateCall();
    }
  },

  watch: {
    messageRecall: {
      handler(v) {
        this.messageList = this.messageList == null ? [] : this.messageList;
        this.messageList.push(v.message)
      }
    },

    historyRecall: {
      handler(v) {
        if (v.to == this.toName) {
          this.messageList = v.messageList;
        }
      }
    },

    searchRecall: {
      handler(v) {
        var self = this;
        if (v.userList != null) {
          this.userList = v.userList.filter(user => user != self.userName);
        } else {
          this.userList = v.userList;
        }
      }
    },

    updateRecall: {
      handler(v) {
        if (v.to == this.toName) {
          this.hintList = v.hintList;
          if (!this.isSearching) {
            this.userList = v.userList;
          }
          this.messageList = v.messageList;
        }

        this.polling();
      }
    }
  }
}
</script>

<style scoped>
.chat {
  display: flex;
}

.chat-list {
  float: left;
  height: 100vh;
  width: 20%;
  overflow-y: scroll;
}

.chat-list::-webkit-scrollbar {
  display: none;
}

.chat-content {
  display: inherit;
  flex-direction: column;
  height: 100vh;
  flex-grow: 1;
}

.chat-text {
  height: 10vh;
  flex-grow: 1;
  font-size: 30px;
}
</style>