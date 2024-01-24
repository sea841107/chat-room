<template>
  <div class="list">
    <b-input
        v-model="query"
        type="search"
        placeholder="搜尋用戶"
        class="search"
      ></b-input>
    <b-list-group>
      <b-list-group-item class="d-flex justify-content-between align-items-center"
        v-for="(user, index) in myUserList"
        :key="index"
        :class="{ 'active': user === activeUser }"
        @click="select(user)"
        href="#"
      >
        {{ user }}
        <b-badge v-if="showHint(user)" variant="primary" pill>N</b-badge>
      </b-list-group-item>
    </b-list-group>
  </div>
</template>

<script>
export default {
  name: 'ChatList',
  props: {
    toName: String,
    userList: Array,
    hintList: Array,
  },

  data() {
    return {
      query: "",
      myHintList: [],
      myUserList: [],
      activeUser: "",
    }
  },

  methods: {
    select(user) {
      this.activeUser = user;
      this.removeHint(user);
      this.$emit("select", user);
    },

    search(query) {
      this.$emit("search", query);
    },

    updateElements() {
      var self = this;
      this.$nextTick(() => {
        this.activeUser = "";
        var elements = document.getElementsByClassName("list-group-item");
        for (var i = 0; i < elements.length; i++) {
          if (elements[i].innerText == self.toName) {
            this.activeUser = elements[i].innerText;
            break;
          }
        }
      });
    },

    showHint(user) {
      if (this.myHintList != null) {
        return this.myHintList.indexOf(user) != -1;
      }

      return false;
    },

    removeHint(name) {
      if (this.myHintList != null) {
        this.myHintList = this.myHintList.filter(hintName => hintName != name);
      }
    }
  },

  watch: {
    query: {
      handler(v) {
        this.search(v.trim());
      }
    },

    userList: {
      handler(v) {
        if (this.query == "") {
          this.myUserList = ["Lobby"].concat(v);
        } else {
          this.myUserList = v;
        }

        this.updateElements();
      }
    },

    hintList: {
      handler(v) {
        this.myHintList = v;
      }
    },
  }
}
</script>

<style scoped>
.list {
  box-shadow: 0 0 1px #000;
  background-color: white;
  overflow-x: hidden;
	overflow-y: auto;
}

.search {
  width: 100%;
  border: 2px solid black;
  background-color:azure;
}
</style>
