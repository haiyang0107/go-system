<template>
  <div>
    <header class="head-nav">
      <el-row>
        <el-col :span="3" class="logo-container">
          <div class="header-logo">GSM平台</div>
        </el-col>
        <el-col :span="15">
          <el-menu
            :default-active="activeIndex2"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#313946"
            text-color="#fff"
            active-text-color="#ffd04b"
          >
            <el-menu-item index="1">处理中心</el-menu-item>
            <el-submenu index="2">
              <template slot="title">我的工作台</template>
              <el-menu-item index="2-1">选项1</el-menu-item>
              <el-menu-item index="2-2">选项2</el-menu-item>
              <el-menu-item index="2-3">选项3</el-menu-item>
              <el-submenu index="2-4">
                <template slot="title">选项4</template>
                <el-menu-item index="2-4-1">选项1</el-menu-item>
                <el-menu-item index="2-4-2">选项2</el-menu-item>
                <el-menu-item index="2-4-3">选项3</el-menu-item>
              </el-submenu>
            </el-submenu>
            <el-menu-item index="3" disabled>消息中心</el-menu-item>
            <el-menu-item index="4">
              <a href="https://www.ele.me" target="_blank">订单管理</a>
            </el-menu-item>
          </el-menu>
        </el-col>
        <el-col :span="6" class="userinfo">
          <span class="username">
            <el-dropdown trigger="click" @command="setDialogInfo">
              <span class="el-dropdown-link">
                admin
                <i class="el-icon-caret-bottom"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="info">修改信息</el-dropdown-item>
                <el-dropdown-item command="pass">修改密码</el-dropdown-item>
                <el-dropdown-item command="logout">退出</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </span>
        </el-col>
      </el-row>
    </header>
     <el-dialog :title="dialog.title"
                :visible.sync="dialog.show_pass">
        <el-form style="margin:20px;width:80%;"
                    label-width="100px"
                    :model="dialog.user_info"
                    :rules="dialog.user_info_rules"
                    ref='user_info'>
            <el-form-item class='edit-form'
                            label="用户名称"
                            prop='userName'>
                <el-input v-model="dialog.user_info.userName" disabled></el-input>
            </el-form-item>
            <el-form-item class='edit-form'
                            label="当前密码"
                            prop='oldPassword'>
                <el-input
                        type='password'
                        placeholder='当前密码'
                        auto-complete='off'
                        v-model="dialog.user_info.oldPassword"></el-input>
            </el-form-item>
            <el-form-item class='edit-form'
                            label="新密码"
                            prop='password'>
                <el-input
                        type='password'
                        placeholder='新密码'
                        auto-complete='off'
                        v-model="dialog.user_info.password"></el-input>
            </el-form-item>
            <el-form-item class='edit-form'
                            label="确认密码"
                            prop='passwordConfirm'>
                <el-input
                        type='password'
                        placeholder='确认密码'
                        auto-complete='off'
                        v-model="dialog.user_info.passwordConfirm"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialog.show_pass = false">取 消</el-button>
            <el-button type="primary" @click="updatePassword('user_info')">确 定</el-button>
        </span>
    </el-dialog>
    <el-dialog :title="dialog.title"
                :visible.sync="dialog.show_set">
        <el-form style="margin:20px;width:80%;"
                    label-width="100px"
                    v-model='dialog.set_info'
                    ref='set_info'>
            <el-form-item label="用户名称">
                <el-input type="text" disabled v-model="dialog.set_info.userName"></el-input>
            </el-form-item>
            <el-form-item label="用户角色">
                <el-input type="text" disabled v-model="dialog.set_info.roleName"></el-input>
            </el-form-item>
            <el-form-item label="常用邮箱">
                <el-input type="text" v-model="dialog.set_info.email"></el-input>
            </el-form-item>
            <el-form-item label="联系电话">
                <el-input type="text" v-model="dialog.set_info.phoneNumber"></el-input>
            </el-form-item>
            <el-form-item label="默认首页">
                <el-input type="text" v-model="dialog.set_info.customHomePage"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="dialog.show_set = false">取 消</el-button>
            <el-button type="primary" @click="updateUserInfo">确 定</el-button>
        </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'head-nav',
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1',
      dialog: {
        show_set: false,
        show_pass: false,
        title: '修改信息',
        user_info: {
          userId: '',
          userName: '',
          oldPassword: '',
          password: '',
          passwordConfirm: ''
        },
        set_info: {
          userId: '',
          userName: '',
          roleName: '',
          customHomePage: '',
          email: '',
          phoneNumber: ''
        },
        user_info_rules: {
          oldPassword: [{
            required: true,
            message: '旧密码不能为空！',
            trigger: 'blur'
          }],
          password: [{
            required: true,
            message: '新密码不能为空！',
            trigger: 'blur'
          }, {
            trigger: 'blur',
            validator: (rule, value, callback) => {
              if (!(/^[a-zA-Z0-9_-]{5,16}$/.test(value))) {
                callback(new Error('密码至少5位,由大小写字母和数字,-,_组成'))
              } else {
                this.$refs.user_info.validateField('passwordConfirm')
                callback()
              }
            }
          }],
          passwordConfirm: [{
            required: true,
            message: '确认密码不能为空！',
            trigger: 'blur'
          }, {
            trigger: 'blur',
            validator: (rule, value, callback) => {
              if (value === '') {
                callback(new Error('请再次输入密码'))
              } else if (value !== this.dialog.user_info.password) {
                callback(new Error('两次输入密码不一致!'))
              } else {
                callback()
              }
            }
          }]
        }
      }
    }
  },
  created () {},
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    },
    /* 弹出框-修改密码或者系统设置 */
    setDialogInfo (cmditem) {
      if (!cmditem) {
        this.$message('菜单选项缺少command属性')
        return
      }
      switch (cmditem) {
        case 'pass':
          this.dialog.show_pass = true
          this.dialog.title = '修改密码'
          this.$refs['user_info'] && this.$refs['user_info'].resetFields()
          let sessionScope = JSON.parse(sessionStorage.getItem('sessionScope'))
          this.dialog.user_info.userName = sessionScope.userName
          this.dialog.user_info.userId = sessionScope.userId
          break
        case 'info':
          this.dialog.show_set = true
          this.dialog.title = '修改用户信息'
          this.dialog.set_info.userName = 'admin'
          this.dialog.set_info.roleName = '超级管理员'
          //   this.getUserInfo()
          break
        case 'logout':
          this.logout()
          break
      }
    },
    updateUserInfo () {
      this.$message('正在开发中')
    },
    /* 修改密码 */
    updatePassword (userinfo) {
      this.$refs[userinfo].validate((valid) => {
        if (valid) {
          this.$message('该模块正在开发中')
        }
      })
    },
    logout () {
      this.$confirm('你确定退出登录么?', '确认退出', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        window.sessionStorage.removeItem('sessionScope')
        window.sessionStorage.removeItem('isLoginFlag')
        this.$router.push('/login')
      }).catch(action => {
      })
    }
  },
  watch: {
    $route (to, from) {
      console.log(to)
    }
  }
}
</script>

<style>
.head-nav {
  width: 100%;
  height: 60px;
  background: #313946;
  position: fixed;
  top: 0px;
  left: 0px;
  z-index: 1000;
  color: #fff;
}
.header-logo {
  max-width: 152px;
  height: 60px;
  line-height: 60px;
  font-size: 24px;
  text-align: center;
}
.userinfo {
  text-align: right;
}
.username {
  height: 60px;
  line-height: 60px;
  cursor: pointer;
  margin-right: 20px;
  user-select: none;
}
.el-dropdown {
  color: #fff;
}
</style>
