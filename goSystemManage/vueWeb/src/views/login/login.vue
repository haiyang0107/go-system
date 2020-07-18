<template>
    <div class="bg-wrap">
      <div class="login-form" @keydown.enter="login">
        <h3>GSM账号登录</h3>
        <div class="login-wrap">
            <div class="input-group">
              <input class="input-text" type="text" v-model="loginId" placeholder="登录用户名" />
            </div>
            <div class="input-group">
              <input class="input-text" type="password" v-model="password" placeholder="登录密码" />
            </div>
            <el-button type="primary" class="login-btn" @click="login">登 录</el-button>
        </div>
        <div style="width:100%;display:flex;justify-content:center" v-if="isShowNotice">
          <span style="font-size:13px;color:red">{{noticeContent}}</span>
        </div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      loginId: 'admin',
      password: '123456',
      isShowNotice: false,
      noticeContent: ''
    }
  },
  methods: {
    login () {
      let _this = this
      if (!_this.loginId || !_this.password) {
        _this.isShowNotice = true
        _this.noticeContent = '请输入账号和密码'
        return
      }
      if (_this.loginId !== 'admin' || _this.password !== '123456') {
        _this.isShowNotice = true
        _this.noticeContent = '请输入正确的账号和密码'
        return
      }
      // 设置登录操作
      sessionStorage.setItem('isLoginFlag', true)
      sessionStorage.setItem('sessionScope', JSON.stringify({
        'userName': _this.loginId,
        'userId': _this.loginId,
        'password': _this.password
      }))
      // 路由跳转
      _this.$router.push({name: 'home'})
    }
  },
  mounted () {
    this.$message('用户名：admin 密码：123456')
  }
}
</script>

<style>
.bg-wrap {
    width: 100%;
    height: 100%;
    background: url('../../assets/img/bg.png') no-repeat center;
    /* background: url('~@/assets/img/bg.png') no-repeat center; */
    background-size: cover;
    position: relative;
}
.login-form {
  width: 346px;
  height: 350px;
  position: absolute;
  top: 50%;
  right: 120px;
  margin-top: -175px;
  z-index: 4;
  background-color: rgba(255, 255, 255, 1);
}
.login-form h3 {
  padding-left: 20px;
  padding-top: 33px;
  line-height: 20px;
  color: #000;
  font-weight: 400;
  font-size: 18px;
}
.login-wrap {
  width: 323px;
  padding-top: 50px;
  padding-left: 23px;
  margin-bottom: 20px;
}
.input-text {
    color: #333;
    width: 100%;
    height: 40px;
    display: block;
    border: 1px solid #999;
    border-radius: 2px;
    padding: 0 16px;
    margin-bottom: 20px;
    box-sizing: border-box;
}

.login-btn {
  width: 100%;
  height: 40px;
  font-family: PingFangSC-Medium;
  font-size: 16px;
}
</style>
