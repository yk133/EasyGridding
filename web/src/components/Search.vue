<template>
<!--  &lt;!&ndash; 引入样式 &ndash;&gt;-->
<!--  <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">-->
<!--  &lt;!&ndash; 引入组件库 &ndash;&gt;-->
<!--  <script src="https://unpkg.com/element-ui/lib/index.js"></script>-->

  <div id ="app">
    <H1>记录查找
    </H1>
    <P>
    </P>
    <div>

      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content bg-purple">用户名：</div></el-col>
        <el-col :span="12"><div class="grid-content bg-purple">
          <el-input v-model="inputUserId" placeholder="请输入用户id"></el-input>
        </div></el-col>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content bg-purple">开始时间：</div></el-col>
        <el-col :span="12"><div class="grid-content bg-purple">
          <div class="block">
            <span class="demonstration"></span>
            <el-date-picker
                v-model="startTime"
                type="date"
                placeholder="选择日期">
            </el-date-picker>
          </div>
        </div></el-col>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content bg-purple">结束时间：</div></el-col>
        <el-col :span="12"><div class="grid-content bg-purple">
          <div class="block">
            <span class="demonstration"></span>
            <el-date-picker
                v-model="endTime"
                type="date"
                placeholder="选择日期">
            </el-date-picker>
          </div>
        </div></el-col>
      </el-row>

      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content "></div></el-col>
        <el-col :span="12"><div class="grid-content bg-purple">
          <el-button>查找</el-button>
        </div></el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "Search",
  data() {
    return {
      inputUserId: '',
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now();
        }
      },
      startTime: '',
      endTime: '',
    };
  }
}
</script>

<style scoped>

.el-row {
  margin-bottom: 20px;
&:last-child {
   margin-bottom: 0;
 }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #d3dce6;
}
.bg-purple {
  background: #ffccac;
}
.bg-purple-light {
  background: #ffffff;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}

</style>