<template>
  <el-container class="container">
    <adminNav></adminNav>
    <el-container >
      <el-header class="headerBtn">
        <adminHeader></adminHeader>
      </el-header>
      <router-view class="main__body"></router-view>
    </el-container>
  </el-container>
</template>

<script>
import adminHeader from '@/components/admin/adminHeader'
import adminNav from '@/components/admin/adminNav'
export default {
  name: 'indexPage',
  components: {
    adminHeader,
    adminNav
  },
  data (){
    return {

    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.main__body{
  background-color: #f8f8f8;
  height: calc(100% - 60px);
  overflow-y: auto;
}
</style>
