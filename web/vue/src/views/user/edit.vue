<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <strong>新增用户</strong>
      </div>
    </template>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="width: 500px;">
        <el-form-item>
          <el-input v-model="form.id" type="hidden"></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <template v-if="!form.id">
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" type="password"></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="confirm_password">
            <el-input v-model="form.confirm_password" type="password"></el-input>
          </el-form-item>
        </template>
        <el-form-item label="是否是管理员" prop="is_admin">
          <el-switch
              v-model="form.is_admin"
              inline-prompt
              active-text="是"
              :active-value="1"
              inactive-text="否"
              :inactive-value="0"
          />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="form.roles" multiple clearable style="width:100%">
            <el-option v-for="(role,i) in roles" :key="i" :value="role.slug" :label="role.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submit()">保存</el-button>
          <el-button @click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
  </el-card>
</template>
<script>
import userService from '../../api/user'
import roleService from '../../api/role'

export default {
  name: 'user-edit',
  data: function () {
    return {
      roles:[],
      form: {
        id: '',
        name: '',
        email: '',
        is_admin: 0,
        password: '',
        confirm_password: '',
        status: 1,
        roles:[],
      },
      formRules: {
        name: [
          {required: true, message: '请输入用户名', trigger: 'blur'}
        ],
        email: [
          {type: 'email', required: true, message: '请输入有效邮箱地址', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'}
        ],
        confirm_password: [
          {required: true, message: '请再次输入密码', trigger: 'blur'}
        ]
      }
    }
  },
  created () {
    const id = this.$route.params.id
    if (!id) {
      return
    }
    let that = this;
    roleService.all({}, function (roles) {
      console.log('roles',roles)
      that.roles = roles
    })
    userService.detail(id, (data) => {
      if (!data) {
        this.$message.error('数据不存在')
        return
      }
      this.form.id = data.id
      this.form.name = data.name
      this.form.email = data.email
      this.form.is_admin = data.is_admin
      this.form.status = data.status
      this.form.roles = data.roles.split(',').filter(v => {
        return !!v
      })
    })
  },
  methods: {
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        this.save()
      })
    },
    save () {
      let data = Object.assign({}, this.form)
      data.roles = data.roles.join(',')
      userService.update(data, () => {
        this.$router.push('/user/index')
      })
    },
    cancel () {
      this.$router.push('/user/index')
    }
  }
}
</script>
