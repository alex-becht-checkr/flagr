<template>
  <div class="tools-debugging-body">
    <div class="tools-debugging-h4">
      <span>Single Evaluation</span>
    </div>  
    <el-row :gutter="10" style="margin-bottom: 8px;" class="tools-debugging-subheader">
      <el-col :span="6">
        <span>Request</span>
      </el-col>
      <el-col :span="12" :offset="6">
        <span>Response</span>
      </el-col>
    </el-row>
    <el-row :gutter="10">
      <el-col :span="12">
        <vue-json-editor
            v-model="evalContext"
            :showBtns="false"
            ref="evalContextEditor"
            class="json-editor"
        ></vue-json-editor>
      </el-col>
      <el-col :span="12">
        <vue-json-editor
            v-model="evalResult"
            :showBtns="false"
            ref="evalResultEditor"
            class="json-editor"
        ></vue-json-editor>
      </el-col>
    </el-row>
    <el-row :gutter="10">
      <el-col :span="6" :offset="6">
        <div align="right" class="tools-debugging-space">
          <el-button
              size="mini"
              @click="emitEvaluate"
              type="primary"
              plain
          >Evaluate
          </el-button>
        </div>
      </el-col>
    </el-row>

    <div class="tools-debugging-h4">
      <span>Batch Evaluation</span>
    </div>  
    <el-row :gutter="10" style="margin-bottom: 8px;" class="tools-debugging-subheader">
      <el-col :span="6">
        <span>Request</span>
      </el-col>
      <el-col :span="12" :offset="6">
        <span>Response</span>
      </el-col>
    </el-row>
    <el-row :gutter="10">
      <el-col :span="12">
        <vue-json-editor
            v-model="batchEvalContext"
            :showBtns="false"
            ref="batchEvalContextEditor"
            class="json-editor"
        ></vue-json-editor>
      </el-col>
      <el-col :span="12">
        <vue-json-editor
            v-model="batchEvalResult"
            :showBtns="false"
            ref="batchEvalResultEditor"
            class="json-editor"
        ></vue-json-editor>
      </el-col>
    </el-row>
    <el-row :gutter="10">
      <el-col :span="6" :offset="6">
        <div align="right" class="tools-debugging-space">
          <el-button
              size="mini"
              @click="emitEvaluateBatch"
              type="primary"
              plain
          >Batch Evaluate
          </el-button>
        </div>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import vueJsonEditor from "vue-json-editor";

export default {
  name: "advanced-debug-console",
  props: ["flag", "evalContext", "evalResult", "batchEvalContext", "batchEvalResult"],
  methods: {
    emitEvaluate() {
      this.$emit('post-evaluation');
    },
    emitEvaluateBatch() {
      this.$emit('post-evaluation-batch');
    },
    toggleDebugModes() {
      this.showSimpleDebug = !this.showSimpleDebug;
    }
  },
  components: {
    vueJsonEditor
  },
  mounted() {
    this.$refs.evalContextEditor.editor.setMode("code");
    this.$refs.evalResultEditor.editor.setMode("code");
    this.$refs.batchEvalContextEditor.editor.setMode("code");
    this.$refs.batchEvalResultEditor.editor.setMode("code");
  },
  data() {
    return {
      activeName: 'evaluation'
    };
  }
};
</script>
