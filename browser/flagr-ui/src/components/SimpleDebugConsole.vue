<template>
  <div>
    <div class="tools-debugging-body">
      <el-row :gutter="20" style="margin-bottom: 8px;" class="tools-debugging-subheader">
        <el-col :span="12" >
          <span><b>Request</b></span>
        </el-col>
        <el-col :span="12">
          <span><b>Response</b></span>
        </el-col>
      </el-row>

      <el-row class="flag-content" type="flex" align="top" :gutter="20">
        <el-col :span="12">
          <div>
            <el-input size="small" placeholder="Enter Entity ID" v-model="evalContext.entityID">
              <template slot="prepend">Entity ID</template>
            </el-input>
          </div>
          <div>
            <el-input size="small" placeholder="Enter Entity Type" v-model="evalContext.entityType">
              <template slot="prepend">Entity Type</template>
            </el-input>
          </div>
          <div>
            <div class="tools-debugging-top-description">
              <span>Entity Context</span>
            </div>
            <vue-json-editor
                    v-model="evalContext.entityContext"
                    :showBtns="false"
                    ref="evalContextSimpleEditor"
                    class="json-editor"
                ></vue-json-editor>
          </div>
          <div style="margin-top:10px; margin-bottom:10px;">
            <el-row :gutter="10">
              <el-col :span="12">
                <el-checkbox v-model="evalContext.enableDebug"><span style="font-size:13px;">Include segment debug logs</span></el-checkbox>
              </el-col>
              <el-col :span="12">
                <div align="right">
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
          </div>        
        </el-col>
        <el-col :span="12">
          <div>
            <el-tag type="primary" :disable-transitions="true">
              Variant ID:
              <b>{{ evalResult.variantID }}</b>
            </el-tag>
            <el-tag type="primary" :disable-transitions="true">
              Variant Key:
              <b>{{ evalResult.variantKey }}</b>
            </el-tag>  
            <div v-if="evalResult.variantAttachment">
              <div class="tools-debugging-top-description">
                <span>Variant attachment:</span>
              </div>
              <pre id="variant-attachment">{{ evalResult.variantAttachment ? evalResult.variantAttachment : ' ' }}</pre>
            </div>
            <div v-if="evalContext.enableDebug" class="tools-debugging-top-description">
              <span>Segment debug logs:</span>
              <template>
                <el-table
                    :data="evalResult?.evalDebugLog?.segmentDebugLogs"
                    class="segment-debug-logs-table"
                >
                  <el-table-column
                      prop="segmentID"
                      label="Segment ID"
                      width="100">
                  </el-table-column>
                  <el-table-column
                      prop="msg"
                      label="Evaluation Message"
                  >
                  </el-table-column>
                </el-table>
              </template>

            </div>          
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import vueJsonEditor from "vue-json-editor";

export default {
  name: "simple-debug-console",
  props: ["flag", "evalContext", "evalResult"],
  methods: {
    emitEvaluate() {
      this.$emit('post-evaluation');
    }
  },
  components: {
    vueJsonEditor
  },
  mounted() {
    this.$refs.evalContextSimpleEditor.editor.setMode("code");
  }
};
</script>

<style lang="less">
.el-form--label-top .el-form-item__label {
  padding: 0;
}

.el-checkbox__input.is-checked + .el-checkbox__label {
  color: #606266;
}

#divider {
  margin: 2em 0;
  background: #f2f6fc;
  height: 10px;
}

#variant-attachment {
  border: 1px solid #E4E7ED;
  border-radius: 4px;
  padding: 5px 10px;
  margin: 0;
}

.evaluate-btn {
  width: 200px;
  margin: 1em 0;
}

.segment-debug-logs-label {
  margin-top: 22px;
}

.segment-debug-logs-table {
  width: 100%;
  font-size: 0.8rem;

  th.el-table__cell > .cell {
    padding-left: 0;
  }
}

.el-input-group__prepend {
  width: 6em;
}

</style>
