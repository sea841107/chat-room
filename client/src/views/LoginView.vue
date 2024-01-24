<template>
  <div class="login-container">
    <b-container class="mt-5">
      <b-row class="justify-content-center">
        <b-col md="6">
          <b-card title="聊天室">
            <b-form @submit.prevent="loginCall">
              <b-form-group label="帳號" label-for="username">
                <b-form-input v-model="name" required></b-form-input>
              </b-form-group>

              <b-form-group label="密碼" label-for="password">
                <b-form-input type="password" v-model="password" required></b-form-input>
              </b-form-group>

              <b-button type="submit" variant="primary" class="mt-3">登入</b-button>
            </b-form>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
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

  mounted() {
    const user = localStorage.getItem('user');
    if (user) {
      this.$router.push("/chat");
    }
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
          localStorage.setItem('user', v.name);
          this.$router.push('chat');
        } else if (v.status == "Fail") {
          alert('密碼錯誤!');
        } else if (v.status == "New") {
          alert('註冊帳號成功!');
          localStorage.setItem('user', v.name);
          this.$router.push('chat');
        } else {
          alert('伺服器錯誤!');
        }
      }
    }
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

</style>