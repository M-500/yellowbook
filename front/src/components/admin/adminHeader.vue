<template>
  <div class="layout_container">
    <el-header class="header">
      <el-dropdown @command="handleCommand">
        <div class="avatar-warp">
          <!--          <img class="avatar" :src="cover_image_link" alt="" />-->
          <img class="avatar" src="@/static/nav/user__manager.png" alt="" />
          <span>{{ username }}</span>
          <i class="el-icon-arrow-down el-icon--right"></i>
        </div>
        <el-dropdown-menu>
          <!-- <el-dropdown-item command="userCenter">个人中心</el-dropdown-item> -->
          <el-dropdown-item command="exitSys">用户退出</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </el-header>
  </div>
</template>

<script>
import curUser from "@/utils/cur-user";

export default {
  name: "adminHeader",
  data() {
    return {
      username: curUser.getUserName(),
    };
  },
  methods: {
    logout() {
      // 删除session中的数据 清除所有
      window.sessionStorage.clear();
      window.localStorage.clear();
      // 如果当前界面刷新 那么就会清空vuex的数据
      this.$router.push("/login");
      window.location.reload();
    },
    handleCommand(command) {
      if (command === "exitSys") {
        this.logout();
      }
      // if (command === "userCenter") {
      //   this.$router.push({
      //     name: "userCenter",
      //   });
      // }
    },
    outSide() {
      console.log("点个锤子");
    },
  },
};
</script>

<style scoped>
.layout-container {
  width: 100%;
  /* position: fixed; */
}
.header {
  width: 100%;
  height: 164px;
  /*background: #fff;*/
  border-bottom: 1px solid #f5f5f5;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
.brisk-header-tool .indent-increase {
  width: 18px;
  height: 18px;
  cursor: pointer;
  transition-duration: 0.3s;
}
.brisk-header-tool .indent-decrease {
  width: 18px;
  height: 18px;
  cursor: pointer;
  transform: rotate(180deg);
  transition-duration: 0.3s;
}
.header .avatar-warp {
  display: flex;
  align-items: center;
  font-weight: 500;
}
.header .avatar-warp .avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 10px;
}
</style>
