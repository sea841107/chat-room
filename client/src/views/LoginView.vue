<template>
  <div class="about">
    <h1>Login</h1>
    帳號：<input type="text" placeholder="Input your name" v-model="name" @keydown.enter="loginCall"/>
    <p></p>
    密碼：<input type="password" placeholder="Input your password" v-model="password" @keydown.enter="loginCall"/>
    <p></p>
    <button @click='loginCall'>登入</button>
  </div>
</template>

<script>
import Sha from 'jssha';

export default {
  name:'LoginPage',
  data() {
    return {
      name: '',
      password: '',
    }
  },

  props: {
    loginRecall: null,
  },

  methods: {
    loginCall() {
      var name = this.name.trim();
      if (name.length < 2) {
        alert('帳號長度不可小於2個字');
        return;
      } else if (name.length > 20) {
        alert('帳號長度不可大於20個字');
        return;
      }

      var password = this.password.trim();
      if (password.length < 4) {
        alert('密碼長度不可小於4個字');
        return;
      } else if (password.length > 20) {
        alert('密碼長度不可小於20個字');
        return;
      }

      var data = JSON.stringify({
        header: "LoginCall",
        name: name,
        password: this.sha256(name, password)
      });

      this.$emit("post", data);
    },

    sha256(name, password) {
      var sha = new Sha("SHA-256", "BYTES");
      sha.update(name + password);
      return sha.getHash("BYTES");
    }
  },

  watch: {
    loginRecall: {
      handler(v) {
        if (v.status == "Success") {
          alert('登入成功!');
          this.$router.push('chat');
        } else if (v.status == "Fail") {
          alert('密碼錯誤!');
        } else if (v.status == "New") {
          alert('註冊帳號成功!');
          this.$router.push('chat');
        } else {
          alert('伺服器錯誤!');
        }
      }
    }
  }
}
</script>