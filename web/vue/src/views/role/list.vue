<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <strong>角色列表</strong>
      </div>
    </template>
    <el-row type="flex" justify="end">
      <el-col :span="2">
        <el-button type="primary" @click="editRole({})">新增</el-button>
      </el-col>
      <el-col :span="2">
        <el-button type="info" @click="refresh">刷新</el-button>
      </el-col>
    </el-row>
    <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="roleTotal"
        :page-size="searchParams.page_size"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
    </el-pagination>
    <el-table
        :data="roles"
        border
        v-loading="loading"
        style="width: 100%">
      <el-table-column
          prop="id"
          label="角色Id">
      </el-table-column>
      <el-table-column
          prop="name"
          label="名称">
      </el-table-column>
      <el-table-column prop="permissions" label="权限">
        <template #default="scope">
          <el-tag class="ml-2" type="warning" v-for="permission in scope.row.permissions.split(',')"
                :key="permission">{{ !!map[permission] ? map[permission] : permission }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="300">
        <template #default="scope">
          <el-button type="primary" @click="editRole(scope.row)">编辑授权</el-button>
          <el-button type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <el-dialog v-model="dialogVisible" title="编辑">
    <el-form :model="role" ref="roleForm" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="role.name" />
      </el-form-item>
      <el-form-item label="权限">
        <el-transfer
            v-model="role.permissions"
            filterable
            :props="{key: 'key',label: 'name'}"
            :titles="['可选权限','已选权限']"
            filter-placeholder="请输入权限名"
            :data="allPermissions"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submit">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<style scoped>
.ml-2 {margin-left: 5px;margin-bottom: 5px}
</style>
<script>
import roleService from '../../api/role'

export default {
  name: 'user-list',
  data() {
    return {
      loading: false,
      dialogVisible: false,
      allPermissions: [],
      map:{},
      roles: [],
      role: {
        name: '',
        permissions: []
      },
      roleTotal: 0,
      searchParams: {
        page_size: 20,
        page: 1
      }
    }
  },
  created() {
    this.search()
  },
  methods: {
    changePage(page) {
      this.searchParams.page = page
      this.search()
    },
    changePageSize(pageSize) {
      this.searchParams.page_size = pageSize
      this.search()
    },
    search(callback = null) {
      let _this = this;
      _this.loading = true;
      roleService.list(this.searchParams, (data) => {
        _this.loading = false;
        _this.roles = data.data
        _this.allPermissions = data.permissions
        for (let i in _this.allPermissions) {
          _this.map[_this.allPermissions[i].key] = _this.allPermissions[i].name
        }
        _this.roleTotal = data.total
        if (callback) {
          callback()
        }
      })
    },
    editRole(obj) {
      let role = Object.assign({},obj)
      if (role.permissions) {
        role.permissions = role.permissions.split(',')
      }
      this.role = role
      this.dialogVisible = true
    },
    submit(){
      let _this = this
      this.$refs['roleForm'].validate(valid => {
        if (!valid) {
          return false
        }
        let data = Object.assign({},this.role)
        data.permissions = data.permissions.join(',')
        roleService.store(data, function () {
          _this.$message.success('操作成功')
          _this.dialogVisible = false
          _this.search()
        })
      })
    },
    refresh() {
      this.search(() => {
        this.$message.success('刷新成功')
      })
    }
  }
}
</script>
