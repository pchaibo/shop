<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="seusername" placeholder="用户名" style="width: 200px;" class="filter-item" />
      <el-button class="filter-item" type="primary" icon="el-icon-search">搜索</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="outerVisible = true">
        增加
      </el-button>
    </div>
    <el-table v-loading="listLoading" :data="list" element-loading-text="拼命加载中" border fit highlight-current-row>
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="会员" width="100">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.status == 2">停用 </span>
          <span v-if="scope.row.status == 1">正常 </span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="注册时间" width="220">
        <template slot-scope="scope">
          <span>{{ scope.row.createtime }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.Updatetime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作"  align="center">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            更改
          </el-button>
          <el-button v-if="row.status != 'deleted'" size="mini" type="danger" @click="handleDelete(row, $index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
     <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="fetchData" />

    <el-dialog title="增加组" :visible.sync="outerVisible" width="30%">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="组名:">
          <el-input v-model="form.name" placeholder="账号" style="width: 200px;" :disabled="edituser" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="form.remark" placeholder="备注" style="width: 200px;" />
        </el-form-item>
      
        <el-form-item label="状态">
          <el-select v-model="form.status">
            <el-option v-for=" start in options" :key="start.value" :label="start.label" :value="start.value" />
          </el-select>
        </el-form-item>
        

      
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="setfrom">取 消</el-button>
        <el-button type="primary" @click="setadd">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { fetchList, userdel, useradd } from '@/api/sitegroup'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
export default {
  name: 'Sitegroup',
  components: { Pagination },
  data() {
    return {
      list: null,
      seusername: '',
      listLoading: true,
      downloadLoading: false,
      outerVisible: false,
      edituser: false,
      options: [{
        value: 1,
        label: '正常'
      },
      {
        value: 2,
        label: '停用'
      }
      ],
      groupid: [{
        value: 1,
        label: '超级管理员'
      },
      {
        value: 2,
        label: '普通管理员'
      }
      ],
      form: {
        id: 0,
        name: '',
        status: 1
      },
      total: 0,
      listQuery: {
        page: 1,
        limit: 10
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.listLoading = true
      // const { data } = await fetchList()
      // this.list = data.items
      // this.total = 20
      // this.listLoading = false
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    },
    // formatJson(filterVal, jsonData) {
    //   return jsonData.map(v => filterVal.map(j => v[j]))
    // },
    /*
    默认设置
    */
    setfrom(s = 0) {
      this.form = {
        id: 0,
        name: '',
        status: 1
      }
      this.edituser = false
      this.outerVisible = false
      if (s === 1) { setTimeout(() => { location.reload() }, 3000) }
    },
    setadd() {
      console.log(this.form)
      useradd(this.form).then((res) => {
        console.log(res)
        if (res) {
          if (res.code === 2000) {
            this.$message({
              type: 'success',
              message: res.message
            })
            this.setfrom(1)
          } else {
            this.$message({
              type: 'error',
              message: res.message
            })
          }
        }
      })
      // outerVisible: true
    },
    /*
    更新
    */
    handleUpdate(row) {
      this.form = {
        id: row.id,
        name: row.name,
        status: row.status,
        remark: row.remark
      }
      
      this.outerVisible = true
      this.edituser = true
    },
    handleDelete(row, index) {
      this.$confirm('此操作将永久删除该管理员, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deluser(row, index)
      }).catch(e => e)
    },
    deluser(row, index) {
      userdel(row['id']).then((res) => {
        if (res) {
          if (res.code === 2000) {
            this.list.splice(index, 1)
            this.$message({
              type: 'success',
              message: '成功删除:' + row['username']
            })
          } else {
            this.$message({
              type: 'error',
              message: '删除失败'
            })
          }
        }
      })
    }

  }
}
</script>
