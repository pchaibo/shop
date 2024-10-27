<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="seusername" placeholder="用户名" style="width: 200px;" class="filter-item" @keyup.enter.native="fetchData" />
      <el-button class="filter-item" type="primary" icon="el-icon-search" @click="fetchData()">搜索</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="outerVisible = true">
        增加
      </el-button>
    </div>
    <el-table v-loading="listLoading" :data="list"  element-loading-text="拼命加载中" border fit highlight-current-row>
      <el-table-column align="center" label="ID" width="95" sortable>
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="域名" width="220">
        <template slot-scope="scope">
          {{ scope.row.sitename }}
        </template>
      </el-table-column>
      <el-table-column label="操作处理" width="" align="center">
        <template slot-scope="scope">
           <div class="textedit"> 
          <a :href="scope.row.sitename+'/index.php?pingsitemap'"  target="_blank"><i class="el-icon-s-help"  title="提交地图"/></a>
          <a :href="scope.row.sitename+'/index.php?a=del'" target="_blank"><i class="el-icon-minus" alt="封堵根目录"  title="封堵根目录"/></a>
          <a :href="scope.row.sitename+'/index.php?a=del'" target="_blank"><i class="el-icon-remove-outline" alt="封堵所有目录" title="封堵所有目录" /></a>
          
          <a :href="scope.row.sitename+'/index.php?a=del'" target="_blank"> <i class="el-icon-close"  title="杀掉进程"/></a>
          <a :href="scope.row.sitename+'/index.php?a=del'" target="_blank"><i class="el-icon-lock"  title="开启进程"/></a>
          </div>
        
        
          
        </template>
      </el-table-column>
 
     
      <el-table-column label="进程" width="100" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.status == 2">停用 </span>
          <span v-if="scope.row.status == 1">正常 </span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="分组" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.groupid }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="后台" width="220">
        <template slot-scope="{row,$index}">
          <svg-icon icon-class="message"  @click="showurl(row,$index)"/>
        </template>
      </el-table-column>

      <el-table-column align="center" label="添加时间" width="220">
        <template slot-scope="scope">
          <span>{{ scope.row.createtime }}</span>
        </template>
      </el-table-column>
    
      <!-- <el-table-column label="操作" width="160" align="center">
        <template slot-scope="{row,$index}">
           <el-button type="primary" size="mini" @click="handleUpdate(row)">
            更改
          </el-button> 
          <el-button v-if="row.status != 'deleted'" size="mini" type="danger" @click="handleDelete(row, $index)">
            删除
          </el-button>
        </template>
      </el-table-column>
       -->
    </el-table>
    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="fetchData" />

    <el-dialog title="增加入口" :visible.sync="outerVisible" width="width: 500px;">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="接口url:">
          <el-input v-model="form.siteurl" placeholder="url" style="width: 500px;" :disabled="edituser" />
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
import { SiteList, Sitedel, siteadd } from '@/api/user'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
export default {
  name: 'User',
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
        label: '超级会员'
      },
      {
        value: 0,
        label: '普通会员'
      }
      ],
      form: {
        // id: 0,
        // username: '',
        // password: '',
        // status: 1,
        // mobile: '',
        // groupid: 0
      },
      total: 0,
      listQuery: {
        page: 1,
        limit: 10,
        username: ''
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
      this.listQuery['username'] = this.seusername
      SiteList(this.listQuery).then(response => {
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
        siteurl: ''
      }
      this.edituser = false
      this.outerVisible = false
      if (s === 1) { setTimeout(() => { location.reload() }, 3000) }
    },
    setadd() {
      console.log(this.form)
      siteadd(this.form).then((res) => {
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
    //
    showurl(row,index){
      console.log(row);
    },
    //
    /*
    更新
    */
    handleUpdate(row) {
      this.form = {
        id: row.id,
        username: row.username,
        status: row.status,
        mobile: row.mobile,
        groupid: row.Group_id

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
      Sitedel(row['id']).then((res) => {
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
    },
    sortmethod(x){
      console.log(x )
    }

  }
}
</script>
<style>
.textedit i{width: 40px; color:blue;font-size: 20px;}
</style>
