<template>
  <div class="list">
    <input class="search" type="search" placeholder="Search..." v-model="query"/>
    <div :class="userNormal" tabindex="-1" v-for="u, i in myUserList" :key="i" @click="select">
      {{ u }}
      <div v-show="showHint(u)" class="hint">
        <img src="../assets/hint.png">
      </div>
    </div>
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
      userNormal: "user-normal",
      userFocus: "user-focus",
      myHintList: [],
      myUserList: [],
    }
  },

  mounted() {
    this.addFocus();
  },

  methods: {
    select(sender) {
      this.removeFocus();
      this.addFocus(sender.target);
      this.removeHint(sender.target.innerText);
      this.$emit("select", sender.target.innerText);
    },

    search(query) {
      this.$emit("search", query);
    },

    addFocus(element) {
      if (element != undefined) {
        element.className = this.userFocus;
      }
    },

    removeFocus() {
      var element = document.getElementsByClassName(this.userFocus)[0];
      if (element != undefined) {
        element.className = this.userNormal;
      }
    },

    updateElements() {
      var self = this;
      this.$nextTick(() => {
        this.removeFocus();
        var elements = document.getElementsByClassName(this.userNormal);
        for (var i = 0; i < elements.length; i++) {
          if (elements[i].innerText == self.toName) {
            this.addFocus(elements[i]);
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
.user-normal {
  font-size: 30px;
	box-shadow: 0 0 1px #000;
  background-color:aliceblue;
  overflow-x: hidden;
}

.user-normal:hover {
  background-color:beige;
}

.user-focus {
  font-size: 30px;
	box-shadow: 0 0 1px #000;
  background-color:bisque;
  overflow-x: hidden;
}

.list {
  box-shadow: 0 0 1px #000;
  background-color: white;
  overflow-x: hidden;
	overflow-y: auto;
}

.search {
  width: 100%;
  font-size: 20px;
  border: 2px solid black;
  background-color:azure;
}

.hint {
  width: 30px;
  height: 30px;
  float: right;
  overflow: hidden;
}

.hint img {
  width: 100%;
  height: 100%;
	vertical-align: middle;
}
</style>
