<template>
  <div class="errPage-container">
    <el-button size="mini" type="primary"  @click="genercode()" @load="genercode" >生成TRC20二维码</el-button>
    <!-- <el-button size="mini" type="primary"  @click="copyToClipboard()" >复制地址</el-button> -->
    <h5>{{ kk }}</h5>
    <div id="qrcode" ref="qrcode"  />
  </div>
</template>
<script>
import QRCode from 'qrcodejs2'
import { getaddress } from '@/api/user'
import VueClipboard from 'vue-clipboard2'
export default {
  name: 'Demo',
  data() {
    return {
      kk: '',
    }
  },
  mounted() {
    this.genercode()
  },
  methods: {
    back() {
      if (this.$route.query.noGoBack) {
        this.$router.push({ path: '/dashboard' })
      } else {
        this.$router.go(-1)
      }
    },
    genercode() {
      let _this=this
      getaddress().then((res) => {
        this.kk = res.date.address
        //_this.kk = res.date.address
        //console.log(res.date.address)
        document.getElementById('qrcode').innerHTML = ''
        //this.html=''
        new QRCode('qrcode', {
          width: 200,
          height: 200,
          text: _this.kk,
        })
      })
      
    },
    copyToClipboard() {
      VueClipboard.copy('要复制的文本内容').then(() => {
        // 复制成功后的处理
        console.log('文本已复制到剪贴板');
      }).catch(() => {
        // 复制失败后的处理
        console.error('复制失败');
      });
    }
  }
}
</script>

<style lang="scss" scoped>
  .errPage-container {
    width: 800px;
    max-width: 100%;
    margin: 100px auto;
    text-align: center;
    #qrcode{
      margin: 0 auto;
      text-align: center;
      width: 200px;
      .img{
        margin: 0 auto;
      }
    }
  }
</style>
