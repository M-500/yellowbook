<template>
  <div>
    <el-card class="card_body">
      <div class="search_top">
        <!--        顶部搜索-->
        <el-form ref="form" :model="form" inline>
          <el-form-item>
            <el-input class="top_search" maxlength="10" prefix-icon="el-icon-search" placeholder="请输入问题(关键字)搜索" @input="searchQuestion" v-model="form.searchKey" size="small" clearable>
            </el-input>
          </el-form-item>
          <el-form-item label="问答类别">
            <el-select v-model="form.searchType" clearable size="small" @change="searchQuestion" placeholder="请选择问答类别">
              <el-option v-for="(item,index) in categories" :key="index" :label="item.name" :value="item.id"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="上传时间">
            <el-date-picker type="date" placeholder="选择日期" @change="searchQuestion" size="small" v-model="form.createTime" style="width: 100%;"></el-date-picker>
          </el-form-item>
          <el-form-item style="margin-left: 40px">
            <el-button size="small" type="primary" @click="searchQuestion">查询</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="body_top">
        <el-button size="small" type="primary" @click="clickHandler('1')"><i class="el-icon-plus"></i> 新增问答</el-button>
        <el-button size="small" type="success" @click="clickHandler('0')">
          <i class="el-icon-upload el-icon--right"></i>批量上传问答
        </el-button>
        <el-button size="small" type="text"><i class="el-icon-bottom"></i> <a :href="ossDownloadUrl">下载上传模板</a></el-button>
      </div>
      <div class="content_body">
        <div class="tag_mark">
          <el-descriptions class="margin-top" :column="10" size="small">
            <el-descriptions-item v-for="(item,index) in cateGroup" :key="index" :label="item.name">{{ item.count }}</el-descriptions-item>
          </el-descriptions>
        </div>
        <!--        表格展示-->
        <div class="table_body">
          <template>
            <el-table stripe size="mini" :header-cell-style="{ background: '#ecf5ff', color: '#606266' }" :data="tableData" border style="width: 100%">
              <el-table-column min-width="50" prop="id" label="ID" width="50">
              </el-table-column>
              <el-table-column prop="categoryId" stripe="true" label="问题类别" width="100">
                <template slot-scope="scope">
                  <el-tag size="mini" effect="dark" type="success">{{cateDict[scope.row.categoryId]}}</el-tag>

                </template>
              </el-table-column>
              <el-table-column prop="title" label="问题" width="280">
              </el-table-column>
              <el-table-column prop="content" label="解答">
              </el-table-column>
              <el-table-column prop="create_time" width="150" label="上传时间">
              </el-table-column>
              <el-table-column fixed="right" width="280" label="操作" align="center">
                <template slot-scope="scope">
                  <el-button type="text" icon="el-icon-edit" size="mini" @click="openEdit(scope.row)">编辑内容</el-button>
                  <el-button type="text" icon="el-icon-delete" size="mini" @click="openDelete(scope.row)">删除</el-button>
                  <el-button type="text" v-if="scope.row.starTime" icon="el-icon-magic-stick" size="mini" @click="cancelStar(scope.row)">取消星标</el-button>
                  <el-button type="text" v-else icon="el-icon-magic-stick" size="mini" @click="openStar(scope.row)">设为星标(最近问答)</el-button>
                </template>
              </el-table-column>
            </el-table>
          </template>
        </div>
      </div>
      <!--      底部分页器-->
      <div class="body_paging">
        <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="curPage" :page-sizes="[8, 15, 20, 30]" :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total">
        </el-pagination>
      </div>
    </el-card>
    <!--    批量上传问答抽屉-->
    <el-drawer :with-header="false" :visible.sync="openDrawer" :close-on-click-modal="false">
      <div class="drawer_top">
        <div class="top_title">批量问答 上传</div>
        <div class="top_mark">
          <i class="el-icon-warning"></i>
          <span class="mark_detail">注: 请严格按照模板上传Excel,否则将解析失败</span>
        </div>
      </div>
      <el-form class="form_item" :model="uploadForm" label-width="100px">
        <el-form-item label="上传EXCEL">
          <el-upload style="width: 100%;" drag :multiple="false" :on-success="successUpload" :headers="headers" :action="uploadUrl">
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            <div class="el-upload__tip" slot="tip">只能上传xlsx/xlsm/xls文件，文件不要超过200行</div>
          </el-upload>
        </el-form-item>
        <el-form-item v-show="canParse">
          <el-button type="primary" @click="parseExcel">解 析</el-button>
          <el-button @click="closeDrawer">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!--    单个上传问答表单-->
    <el-dialog title="新增问答" :visible.sync="openAddDrawer" width="50%">
      <el-form class="form_item" :rules="rules" :model="addQuestionForm" ref="addQuestionForm" label-width="100px">
        <el-form-item label="问答类别:" prop="questionType">
          <el-select style="width:100%" v-model="addQuestionForm.questionType" placeholder="请选择问答类别">
            <el-option v-for="(item,index) in categories" :key="index" :label="item.name" :value="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="问题:" prop="questionTitle">
          <el-input type="textarea" placeholder="请输入问题" :autosize="{ minRows: 2, maxRows: 6}" v-model="addQuestionForm.questionTitle"></el-input>
        </el-form-item>
        <el-form-item label="解答:" prop="questionAnswer">
          <el-input type="textarea" placeholder="请输入解答" :autosize="{ minRows: 6, maxRows: 12}" v-model="addQuestionForm.questionAnswer"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="addArticle">新 增</el-button>
          <el-button size="small" @click="openAddDrawer=false" style="margin-left: 20px">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-dialog title="编辑内容" :visible.sync="openUpdateDrawer" width="50%">
      <el-form class="form_item" :model="updateQuestionForm" label-width="100px">
        <el-form-item label="类别">
          <el-select style="width:100%" v-model="updateQuestionForm.questionType" placeholder="请选择问答类别" disabled>
            <el-option v-for="(item,index) in categories" :key="index" :label="item.name" :value="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="问题">
          <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 6}" v-model="updateQuestionForm.questionTitle"></el-input>
        </el-form-item>
        <el-form-item label="解答">
          <el-input type="textarea" :autosize="{ minRows: 6, maxRows: 12}" v-model="updateQuestionForm.questionAnswer"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="updateArticleBtn">更 新</el-button>
          <el-button size="small" @click="openUpdateDrawer=false" style="margin-left: 20px">取 消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

  </div>
</template>

<script>
import { host } from '@/config'
import { checkExcelAPI } from "@/api/sj-kit/financeExcel";
import {
  addArticleAPI,
  updateArticleAPI,
  getArticlesCountAPI,
  parseExcel,
  getArticlesAPI,
  cancelStar,
  getCateListAPI,
  deleteArticle,
  starArticle
} from "@/api/question/question";



export default {
  name: "questionManager",
  data() {
    return {
      ossDownloadUrl: "https://51huoke.oss-cn-shenzhen.aliyuncs.com/sj-file-templates/%E9%97%AE%E7%AD%94%E6%A8%A1%E6%9D%BF.xlsx",
      curPage: 1,
      pageSize: 8,
      titleTemplate: ["分类名称", "解答", "问题"],
      total: 0,
      openDrawer: false,
      openAddDrawer: false,
      openUpdateDrawer: false,
      canParse: false, // 是否可以解析
      uploadUrl: host + "/cms/excel/upload",
      remoteFilePath: "resource/excel/20220811/1660205671770581000.xlsm",
      uploadForm: {
        filePath: "",
        fileName: "",
      },
      cateGroup: [],
      form: {
        searchKey: "",
        searchType: "",
        createTime: ''
      },
      categories: [],
      categorieNames: [],
      cateDict: {},
      addQuestionForm: {
        questionTitle: "",
        questionAnswer: "",
        questionType: "",
      },
      rules: {
        questionType: [{ required: true, message: '请选择问答类别', trigger: 'change' },],
        questionTitle: [{ required: true, message: '请输入问题', trigger: 'blur' },],
        questionAnswer: [{ required: true, message: '请输入回答', trigger: 'blur' },],
      },
      updateQuestionForm: {
        id: null,
        questionTitle: "",
        questionAnswer: "",
        questionType: "",
      },
      headers: {
        Authorization: window.localStorage.getItem('cust-token') // 从cookie里获取token，并赋值  Authorization ，而不是token
      },
      tableData: [
        {
          date: '2016-05-03',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1516 弄'
        }
      ]
    }
  },
  created() {
    // this.getAllCategories()
    // this.searchQuestion()
    // this.getGroupCate()
  },
  methods: {
    /** 获取文章分类统计信息 */
    getGroupCate() {
      getArticlesCountAPI().then(res => {
        this.cateGroup = res
      }).catch(e => {
        this.$message.error(e.message)
      })
    },
    /** 取消星标 */
    cancelStar(data) {
      this.$confirm('是否将文章取消星标?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let params = {
          id: data.id
        }
        cancelStar(params).then(res => {
          this.$message({
            type: 'success',
            message: "操作成功！"
          });
          this.searchQuestion()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消操作'
        });
      });
    },
    /** 将文章设置星标 */
    openStar(data) {
      this.$confirm('是否将文章设置为最近问答?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let params = {
          id: data.id
        }
        starArticle(params).then(res => {
          this.$message({
            type: 'success',
            message: "操作成功！"
          });
          this.searchQuestion()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消星标问答'
        });
      });

    },
    /** 编辑内容 */
    openEdit(item) {
      this.updateQuestionForm.id = item.id
      this.updateQuestionForm.questionType = this.cateDict[item.categoryId]
      this.updateQuestionForm.questionTitle = item.title
      this.updateQuestionForm.questionAnswer = item.content
      this.openUpdateDrawer = true
    },
    /** 更新文章*/
    updateArticleBtn() {
      updateArticleAPI(this.updateQuestionForm).then(res => {
        this.$message.success("问答更新成功！")
        this.searchQuestion()
        this.openUpdateDrawer = false
      }).catch(e => {
        this.$message.error(e.message)
      })
    },
    /** 删除文章 */
    openDelete(data) {
      this.$confirm('此操作将永久删除问答, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        let params = {
          id: data.id
        }
        deleteArticle(params).then(res => {
          this.$message({
            type: 'success',
            message: "操作成功！"
          });
          this.searchQuestion()

          this.getGroupCate()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除问答'
        });
      });
    },
    /**  获取所有分类列表 */
    getAllCategories() {
      getCateListAPI().then(res => {
        this.categories = res.list
        for (let item of res.list) {
          this.cateDict[item.id] = item.name
          this.categorieNames.push(item.name)
        }
      })
      this.searchQuestion()
    },
    // 上传文件的钩子函数
    successUpload(res, file) {
      if (res.success) {
        this.uploadForm.filePath = res.data.file_path
        this.uploadForm.fileName = res.data.file_name
        this.$message.success(res.message)
        this.canParse = true
        // this.checkExcelTemplate() // 校验文件模板是否合法
      } else {
        return this.$message.success(res.msg)
      }
    },
    parseExcel() {
      let params = {
        file_path: this.uploadForm.filePath,
        file_name: this.uploadForm.fileName,
      }
      parseExcel(params).then(res => {
        this.$message.success("导入成功:" + res.SuccessNum + "条，失败：" + res.Failed + "条")
      }).catch(e => {
        this.$message.error(e.msg)
      })
    },
    closeDrawer() {
      this.openDrawer = false
    },

    clickHandler(i) {
      switch (i) {
        case '0':
          this.openDrawer = true;
          break
        case '1':
          this.openAddDrawer = true;
          break
        default:
          console.log('嘿嘿')
      }

    },
    /** 搜索问答列表 */
    searchQuestion() {
      let params = {
        searchKey: this.form.searchKey,
        searchType: this.form.searchType,
        createTime: this.form.createTime,
        pageNum: this.curPage,
        pageSize: this.pageSize
      }
      getArticlesAPI(params).then(res => {
        this.tableData = res.list
        this.total = res.total
      })
    },
    /** 点击添加问答 */
    addArticle() {
      this.$refs['addQuestionForm'].validate((valid) => {
        if (valid) {
          addArticleAPI(this.addQuestionForm).then(res => {
            this.openAddDrawer = false
            this.searchQuestion()
            this.$message.success("创建问答成功！")
            this.addQuestionForm.questionAnswer = ""
            this.addQuestionForm.questionTitle = ""
            this.addQuestionForm.questionType = ""
          }).catch(e => {
            this.$message.error(e.message)
          })
        } else {
          console.log('error submit!!');
          return false;
        }
      });

    },
    /** 文件上传成功触发事件 */
    handleUploadSuccess(res, file) {
      this.$message(res.message);
      this.remoteFilePath = res.data.file_path
      console.log(this.remoteFilePath)
      // 校验excel
      let params = {
        file_path: this.remoteFilePath
      }
      checkExcelAPI(params).then(res => {
        console.log("校验excel", res)
      }).catch(e => {

      })
    },
    /** 页大小变化激活事件 */
    handleSizeChange(e) {
      this.pageSize = e
      this.searchQuestion()
    },
    /** 页码变化激活事件 */
    handleCurrentChange(e) {
      this.curPage = e
      this.searchQuestion()
    },

    checkExcelTemplate() {
      let params = {
        file_path: this.remoteFilePath
      }
      checkExcelAPI(params).then(res => {
        this.$message({
          message: 'excel校验通过',
          type: 'success'
        });
      }).catch(e => {
        console.log(e.message)
        // 1. 关闭抽屉
        this.openDrawer = false
        // 2. 弹出需要手动关闭的通知
        this.$notify.error({
          title: 'Excel校验错误',
          message: e.message,
          duration: 0
        });
      })
    }
  }
}
</script>

<style scoped lang="scss">
.el-card.is-always-shadow,
.el-card.is-hover-shadow:focus,
.el-card.is-hover-shadow:hover {
  box-shadow: none;
}

.content_body {
  margin-top: 10px;
  height: 100%;
}

.tag_mark {
  margin-top: 10px;
}

::v-deep .el-drawer__body {
  background: #f5f5f5 !important;
}

.drawer_top {
  margin: 10px 30px 20px 20px;
}
.body_paging {
  display: flex;
  justify-content: end;
  padding-top: 10px;
}
.top_title {
  font-size: 19px;
  font-weight: bold;
  padding-bottom: 10px;
  margin-bottom: 5px;
  border-bottom: 1px solid #ffffff;
}

.form_item {
  margin-left: 20px;
  margin-right: 30px;
}

.top_mark {
  height: 40px;
  display: flex;
  background-color: #fff5ed;
  align-items: center;
  border-radius: 2px;
  margin-bottom: 12px;
  padding-left: 10px;
}

.el-icon-warning {
  color: #f67207;
}

.mark_detail {
  line-height: 22px;
  color: #f67207;
  margin-left: 8px;
  font-size: 14px;
}

::v-deep .el-upload.el-upload--text {
  height: 100px !important;

  .el-upload-dragger {
    width: 100%;
    //height: 100px !important;
  }

  width: 100%;
}
</style>
