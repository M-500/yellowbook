<template>
  <div class="main__content">
    <div class="loginBox">
      <div class="top">
        <img class="logo__img" src="@/static/sign-in/login/logo__2x.png" alt="" />
        <h1>数鲸数字后台管理系统</h1>
      </div>
      <el-form :model="form" ref="loginForm" :rules="rules" class="loginForm" size="middle">
        <el-form-item prop="username">
          <el-input v-model="form.username" maxlength="11" show-word-limit prefix-icon="el-icon-mobile-phone" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" placeholder="请输入密码" type="password" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>
        <el-form-item prop="captcha_code">
          <div class="captcha">
            <el-input v-model="form.captcha_code" placeholder="图片验证码" prefix-icon="el-icon-document"></el-input>
            <img @click="getCaptchaImg" class="captcha__img" :src="captchaUri" alt="" />
          </div>
        </el-form-item>
      </el-form>
      <div class="boxFooter">
        <el-button :disabled="!(form.username && form.password && form.captcha_code)" type="primary" @click="pwdLogin" class="loginBtn" round>登 录</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import { imageCaptchaAPI, pwdLoginAPI } from "@/api/sign-in/login";
import curUser from "@/utils/cur-user";
export default {
  name: "userLogin",
  data() {
    return {
      logoImg: "",
      captchaUri: "",
      form: {
        username: "admin",
        password: "wulinlin",
        captcha_code: "",
        captcha_id: "",
        login_type: "cms",
      },
      rules: {
        username: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur",
          },
          {
            min: 3,
            max: 15,
            message: "长度在 3 到 15 个字符",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        captcha_code: [
          {
            required: true,
            message: "请输入验证码",
            trigger: "blur",
          },
          {
            min: 1,
            max: 6,
            message: "验证码长度不合法",
            trigger: "blur",
          },
        ],
      },
    };
  },
  created() {
    this.getCaptchaImg();
  },
  methods: {
    getCaptchaImg() {
      imageCaptchaAPI()
        .then((res) => {
          this.form.captcha_id = res.captcha_id;
          this.captchaUri = res.pic_path;
        })
        .catch((e) => {
          console.log(e);
        });
    },
    pwdLogin() {
      this.$refs["loginForm"].validate(async (valid) => {
        if (!valid) return this.$message.error("非法输入数据，请重新输入");
        pwdLoginAPI(this.form)
          .then((res) => {
            this.$message({
              message: "登录成功！",
              type: "success",
            });
            curUser.setToken(res.token);
            curUser.setUserName(res.user_name);
            this.$router.push("/ldxf-cms/question-manager");
          })
          .catch((e) => {
            this.$message({
              message: e.message,
              type: "error",
            });
          });
      });
    },
  },
};
</script>

<style scoped>
.main__content {
  height: 100%;
  background: #01020b url('../../static/sign-in/login/backgroud.png') no-repeat
    center/cover;
  background-size: cover;
}
.loginBox {
  width: 450px;
  height: 500px;
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -60%);
}
.top {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}
.logo__img {
  width: 116px;
  height: 116px;
}

.loginForm {
  padding-top: 20px;
  padding-left: 30px;
  padding-right: 30px;
}

.captcha {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.captcha__img {
  margin-left: 20px;
  width: 150px;
  height: 40px;
  border-bottom: 1px solid #dbdbdb;
}
.boxFooter {
  padding-left: 30px;
  padding-right: 30px;
}
.loginBtn {
  width: 100%;
  font-weight: bold;
  font-size: 20px;
}
</style>
