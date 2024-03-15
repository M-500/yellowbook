<template>
  <div>
    <el-row>
<!--      左侧用户信息显示区域 -->
      <el-row :span="6">
        <div class="body_left">
          <div class="user_card">
            <div class="user_avatar">
              <el-avatar shape="square" :size="100" :fit="fit" :src="url"></el-avatar>
            </div>
          </div>
        </div>
      </el-row>

      <el-col :span="18">

      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "userCenter",
  data (){
    return{
      fit:"琳琳可狗可乐",
      url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
    }
  }
}
</script>

<style scoped>
.user_card{
  padding: 30px 20px;
  /*box-shadow: -2px 0 20px -16px;*/
  /*width: 80%;*/
  text-align: center;
  height: 100%;
}
.el-avatar {
  border-radius: 50%;
}
.user_avatar{
  width: 178px;
  height: 178px;
  display: block;
}
</style>
