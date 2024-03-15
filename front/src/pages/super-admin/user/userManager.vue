<template>
  <div class="main_container">
    <el-card>
      <!-- 顶部搜索用户 -->
      <div class="search_top">
        <div class="left">
          <el-input
              class="top_search"
              maxlength="15"
              prefix-icon="el-icon-search"
              placeholder="输入用户名搜索"
              @input="remoteUsers"
              v-model="searchUserForm.searchKey"
              size="mini"
              clearable></el-input>
          <div class="item">
            <span>用户类型</span>
            <el-select v-model="searchUserForm.userType" clearable placeholder="请选择用户类型" size="mini" @change="remoteUsers">
              <el-option label="全部" value="0"></el-option>
              <el-option label="小程序" value="1"></el-option>
              <el-option label="管理后台" value="2"></el-option>
            </el-select>
          </div>
          <div class="item">
            <span>用户状态</span>
            <el-select v-model="searchUserForm.userStatus" clearable placeholder="请选择用户状态" size="mini" @change="remoteUsers">
              <el-option label="使用中" value="false"></el-option>
              <el-option label="已禁用" value="true"></el-option>
            </el-select>
          </div>
          <div class="item">
            <span>创建日期</span>
            <el-date-picker type="date" placeholder="选择日期"
                            @change="remoteUsers"
                            v-model="searchUserForm.createDate"
                            format="yyyy 年 MM 月 dd 日"
                            value-format="yyyy-MM-dd"
                            size="mini">
            </el-date-picker>
          </div>
        </div>
        <el-button class="right" size="mini" type="primary" @click="remoteUsers"><i class="el-icon-search"></i> 查询
        </el-button>
      </div>
<!--      新增用户按钮-->
      <div class="body_top">
        <el-button size="mini" type="primary" @click="addUser"><i class="el-icon-plus"></i> 新增用户
        </el-button>
      </div>
      <!--      中间部分用户表格展示-->
      <div class="body_table">
        <el-table
            stripe
            size="mini"
            :header-cell-style="{ background: '#ecf5ff', color: '#606266' }"
            :data="userTableList"
            height="550"
            style="width: 100%"
        >
          <el-table-column min-width="50" prop="id" label="ID" width="50">
          </el-table-column>
<!--          <el-table-column-->
<!--              prop="cover_img_link"-->
<!--              min-width="55"-->
<!--              stripe="true"-->
<!--              align="center"-->
<!--              label="头像"-->
<!--          >-->
<!--            <template slot-scope="scope">-->
<!--              <el-avatar-->
<!--                  size="medium"-->
<!--                  :src="scope.row.cover_img_link"-->
<!--              ></el-avatar>-->
<!--            </template>-->
<!--          </el-table-column>-->
          <el-table-column
              prop="user_name"
              min-width="100"
              align="center"
              label="用户名"
              width="100"
          >
          </el-table-column>
          <el-table-column
              prop="nickname"
              min-width="130"
              align="center"
              label="状态"
              width="用户名"
          >
            <template slot-scope="scope">
              <el-tag size="mini" effect="dark" type="success" v-if="scope.row.active">使用中</el-tag>
              <el-tag size="mini" effect="dark" type="info" v-else>已禁用</el-tag >
            </template>
          </el-table-column>
          <el-table-column
              min-width="180"
              prop="phone"
              label="手机号"
              align="center"
              width="150"
          >
          </el-table-column>
          <el-table-column
              min-width="180"
              prop="password"
              label="密码"
              align="center"
              width="180"
          >
            <template slot-scope="scope">
              <div  style="margin-left:10px;font-size:12px" v-if="!scope.row.show_pwd">
                ******** <span style="color:#4d70ff;margin-left:10px;font-size:12px" @click="showPwd(scope.$index)">查看</span>
              </div>
              <div  style="margin-left:10px;font-size:12px" v-else>
                {{ scope.row.password }} <span style="color:#4d70ff;margin-left:10px;font-size:12px" @click="hidePwd(scope.$index)">隐藏</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column min-width="200" prop="is_admin" label="权限范围" align="center">
            <template slot-scope="scope">
<!--              <el-tag size="mini" effect="dark" type="danger" v-if="scope.row.is_admin && scope.row.is_wechat">小程序+后台</el-tag>-->
<!--              <el-tag size="mini" effect="dark" type="info"  v-else-if="scope.row.is_wechat" style="margin-left: 2px">小程序</el-tag>-->
<!--              <el-tag size="mini" effect="dark"  v-else>后台</el-tag >-->
              <el-checkbox size="mini" label="后台" name="type" v-model="scope.row.is_admin" @change="changePermission(scope.row,'cms')"></el-checkbox>
              <el-checkbox size="mini" label="小程序" name="type" v-model="scope.row.is_wechat" @change="changePermission(scope.row,'wechat')"></el-checkbox>
            </template>
          </el-table-column>
          <el-table-column
              min-width="150"
              prop="create_time"
              align="center"
              label="创建时间"
          >
<!--            <template slot-scope="scope">-->
<!--              {{Date.from(scope.row.create_time).toDateString()}}-->
<!--            </template>-->
          </el-table-column>
          <el-table-column fixed="right" width="250" label="操作" align="center">
            <template slot-scope="scope">
              <el-button
                  type="text"
                  icon="el-icon-delete"
                  size="mini"
                  v-if="scope.row.active"
                  @click="openDelete(scope.row)">禁用账号</el-button>
              <el-button
                  type="text"
                  icon="el-icon-check"
                  size="mini"
                  v-else
                  @click="openActive(scope.row)">解封账号</el-button>
              <el-button
                  type="text"
                  icon="el-icon-magic-stick"
                  size="mini"
                  @click="changePwd(scope.row)">修改密码</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!--      底部分页器-->
      <div class="body_paging">
        <el-pagination
            background
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="curPage"
            :page-sizes="[10, 15, 20, 30]"
            :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
        >
        </el-pagination>
      </div>
    </el-card>
    <!--    新增用户弹出框 -->
    <el-dialog
        title="创建用户"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :visible.sync="addUserDialog"
        width="42%">
      <el-form
          ref="addUserForm"
          :model="addUserForm"
          :rules="addUserRules"
          label-width="80px"
          size="mini">
        <el-form-item label="用户名：" prop="userName" clearable>
          <el-input v-model="addUserForm.userName"></el-input>
        </el-form-item>
        <el-form-item label="手机号：" prop="phone">
          <el-input v-model="addUserForm.phone" maxlength="11"></el-input>
        </el-form-item>
        <el-form-item label="密码：" prop="password" clearable>
          <el-input v-model="addUserForm.password"></el-input>
        </el-form-item>
        <el-form-item label="权限：" prop="nickname" clearable>
          <el-radio-group v-model="addUserForm.isAdmin">
            <el-radio :label="1">小程序</el-radio>
            <el-radio :label="2">后台管理</el-radio>
            <el-radio :label="3">全部</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div class="dialog_footer">
        <el-button
            class="footer_btn"
            @click="addUserDialog = false"
            size="mini">取 消</el-button>
        <el-button
            class="footer_btn"
            type="primary"
            @click="createUserBtn('addUserForm')"
            size="mini">添 加</el-button>
      </div>
    </el-dialog>
    <el-dialog
        title="修改密码"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :visible.sync="changePwdDialog"
        width="42%">
      <el-form
          ref="addUserForm"
          :model="changePwdForm"
          :rules="changePwdRules"
          label-width="100px"
          size="mini">
        <el-form-item label="新密码：" prop="password" clearable>
          <el-input v-model="changePwdForm.password" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认密码：" prop="password2">
          <el-input v-model="changePwdForm.password2" show-password></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog_footer">
        <el-button
            class="footer_btn"
            @click="changePwdDialog = false"
            size="mini">取 消</el-button>
        <el-button
            type="primary"
            @click="changeUserPwd"
            size="mini">确 认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  delUserListAPI,
  getUserListAPI,
  resetUserPasswordAPI,
  activeUserListAPI,
  changeUserPermissionAPI,
  createUserListAPI
} from "@/api/super-admin/userManager";

export default {
  name: "userManager",
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.changePwdForm.password) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      curPage: 1,
      total: 0,
      pageSize: 10,
      addUserDialog: false,
      changePwdDialog: false,
      // 修改密码的表单
      changePwdForm:{
        password:null,
        password2:null,
        id: null,
      },
      // 修改密码规则
      changePwdRules:{
        password:[ { validator: validatePass, trigger: 'blur' }],
        password2:[ { validator: validatePass2, trigger: 'blur' }],
      },
      addUserRules: {
        userName: [
          {required: true, message: "请输入用户名", trigger: "blur"},
          {
            min: 1,
            max: 15,
            message: "长度在 1 到 15 个字符",
            trigger: "blur",
          },
        ],
        password: [
          {required: true, message: "请输入密码", trigger: "blur"},
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        phone: [
          {required: true, message: "请输入手机号", trigger: "blur"},
          {min: 11, max: 11, message: "长度11位手机号", trigger: "blur"},
        ],
        isAdmin: [
          {required: true, message: "请选择用户权限", trigger: "blur"},
        ],
      },
      addUserForm: {
        userName: "",
        password: "",
        phone: "",
        isAdmin: 1,
      },
      userTableList: [],
      searchUserName: "",
      searchUserForm: {
        userType: "0",
        userStatus:"",
        createDate:"",
        searchKey:"",
        pageNum: this.curPage,
        pageSize: this.pageSize
      }
    };
  },
  created() {
    this.remoteUsers();
    this.calcTableHeight();
  },
  methods: {
    /** 重置用户密码*/
    changeUserPwd(){
      resetUserPasswordAPI(this.changePwdForm).then(res =>{
        this.$message.success("密码修改成功")
        this.changePwdDialog = false
        this.remoteUsers()
      }).catch(e =>{
        this.$message.error(e.message)
      })

    },

    showPwd (index){
      this.userTableList[index].show_pwd = true
    },
    hidePwd(index){
      this.userTableList[index].show_pwd = false
    },
    /** 获取用户名列表 */
    remoteUsers() {
      let params = {
        searchKey: this.searchUserForm.searchKey,
        userType: this.searchUserForm.userType,
        userStatus: this.searchUserForm.userStatus,
        createDate: this.searchUserForm.createDate,
        pageNum: this.curPage,
        pageSize: this.pageSize
      }
      getUserListAPI(params).then((res) => {
        this.userTableList = res.list;
        this.total = res.total;
      });
    },
    changePermission(data,userType){
      var parmas = {
        id: data.id,
        permission_tag:userType,
        value: data.is_admin
      }
      if (userType === 'cms') {
        parmas.value = data.is_admin
      }
      if (userType === 'wechat') {
        parmas.value = data.is_wechat
      }
      changeUserPermissionAPI(parmas).then(res =>{
        this.$message.success("操作成功")
      }).catch(e =>{
        this.$message.error(e.message)
      })
    },
    // /** 通过用户名搜索用户列表 */
    // searchUser() {
    //   getUserListAPI(this.searchUserForm).then((res) => {
    //     this.userTableList = res.list;
    //     this.total = res.total;
    //   });
    // },
    calcTableHeight (){
      console.log(window.innerHeight)
    },

    handleSizeChange(e) {
      this.pageSize = e
      this.remoteUsers()
    },
    handleCurrentChange(e) {
      this.curPage = e
      this.remoteUsers()
    },
    /** 点击新增用户按钮 */
    addUser() {
      this.addUserDialog = true;
    },
    /** 点击创建用户*/
    createUserBtn(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid){
          let params = {
            phone: this.addUserForm.phone,
            userName: this.addUserForm.userName,
            password: this.addUserForm.password,
            isAdmin: this.addUserForm.isAdmin,
          }
          createUserListAPI(params).then(res =>{
            this.remoteUsers()
            this.$message.success("成功创建用户 "+res.user_name)
            this.addUserDialog = false;
          }).catch(e=>{
            this.$message.error(e.message)
          })
          // 发起请求创建用户
        }else{
          return false
        }
      })
    },
    /** 修改用户密码 */
    changePwd (data) {
      this.changePwdDialog = true
      this.changePwdForm.id = data.id
    },
    /** 账号解封*/

    openActive(data){
      let params={
        uid: data.id
      }
      activeUserListAPI(params).then(res => {
        this.$message({
          type: 'success',
          message: "操作成功！"
        });
        this.remoteUsers()
      }).catch(e=>{
        this.$message.error(e.message)
      })
    },
    /** 删除用户信息 */
    openDelete(data){
      this.$confirm('此操作将永久禁用用户, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let params={
          uid: data.id
        }
        delUserListAPI(params).then(res => {
          this.$message({
            type: 'success',
            message: "操作成功！"
          });
          this.remoteUsers()
        }).catch(e=>{
          this.$message.error(e.message)
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消禁用用户'
        });
      });
    }
  },
};
</script>

<style scoped lang="scss">
//.main_container {
//  overflow-y: auto;
//}

.body_table {
  margin-top: 13px;
}

.search_top{
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  .left{
    display: flex;
    justify-content: start;
    .item{
      span{
        margin-left: 5px;
        margin-right: 8px;
        font-size: 13px;
        font-weight: 400;
      }
    }
  }

  .search_btn{
    margin-left: 10px;
  }
}
.body_top {
  display: flex;
  justify-content: space-between;
}

.top_search {
  width: 250px;
}

.body_paging {
  display: flex;
  justify-content: end;
  padding-top: 10px;
}

.el-button.el-button--text {
  color: #4d70ff !important;
  margin-right: 10px;
  font-weight: bold;
  border: #4d70ff;
}

.el-button.el-button--text:hover {
  color: #27c27e !important;
}


.dialog_footer {
  display: flex;
  justify-content: end;
}

.footer_btn {
  margin-left: 25px;
  padding-left: 25px;
  padding-right: 25px;
}

/* dialog 头部样式修改 */
::v-deep .el-dialog__header {
  margin-left: 20px;
  margin-right: 20px;
  border-bottom: 1px solid #f5f5f5;
}
/** 修改多选框文字大小*/
::v-deep .el-checkbox__label{
  font-size: 11px;
}


::v-deep .el-dialog__title {
  font-weight: bold;
  color: #333333;
}

::v-deep .el-dialog__body {
  margin-top: -10px;
}
</style>
