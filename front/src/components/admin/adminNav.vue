<template>
  <el-aside class="nav" width="initial">
    <!--    侧边栏菜单展示 通过动态形式-->
    <el-menu
      default-active="2"
      router
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#191a23"
      text-color="#fff"
      active-text-color="#4D70FF"
    >
      <div class="logo-aside" @click="clickTitle">
        <img src="@/assets/logo.png" alt="" class="logo" />
        <p class="title">路打先锋</p>
      </div>

      <el-menu-item index="/ldxf-cms/question-manager">
        <template slot="title">
          <i class="el-icon-notebook-1"></i>
          <span>问答上传</span>
        </template>
      </el-menu-item>
      <el-menu-item index="/ldxf-cms/user-manager">
        <template slot="title">
          <i class="el-icon-user"></i>
          <span>用户管理</span>
        </template>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: "adminNav",
  data() {
    return {
      isCollapse: false,
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    /** 点击头部跳转首页 */
    clickTitle() {
      console.log("点个头");
    },
  },
};
</script>

<style scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  min-height: 100%;
}
.el-menu-vertical-demo {
  min-height: 100%;
}
.nav {
  height: 100%;
}

.logo-aside {
  width: 160px;
  height: 35px;
  padding: 5px 20px;
  margin-top: 10px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.logo {
  width: 30px;
  height: 30px;
  border-radius: 30%;
  margin-right: 10px;
}
.title {
  font-size: 16px;
  font-family: OPPOSans-B, OPPOSans;
  font-weight: 600;
  color: #fff;
}
</style>
