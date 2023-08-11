<template>
  <div>
    <el-dialog title="Delete constraint" :visible.sync="dialogDeleteConstraintVisible">
      <span>Are you sure you want to delete this constraint?</span>
      <span slot="footer" class="dialog-footer">
              <el-button @click="cancelConstraintDeletion">Cancel</el-button>
              <el-button type="primary" @click="confirmConstraintDeletion">Confirm</el-button>
            </span>
    </el-dialog>

    <h5>Constraints (match ALL of them)
      <el-tooltip placement="bottom" effect="light">
        <div slot="content">Used to validate if a particular segment<br>
          is eligible to be matched with a variant.</div>
        <span class="el-icon-info" style="vertical-align: middle;"></span>
      </el-tooltip>
    </h5>

    <div class="constraints">
      <div class="constraints-inner" v-if="segment.constraints.length">
        <div v-for="constraint in segment.constraints" :key="constraint.id">
          <el-row :gutter="3" class="segment-constraint">
            <el-col :span="20">
              <el-input size="small" placeholder="Property" v-model="constraint.property">
                <template slot="prepend">Property</template>
              </el-input>
            </el-col>
            <el-col :span="4">
              <el-select class="width--full" size="small" v-model="constraint.operator"
                         placeholder="operator">
                <el-option v-for="item in operatorOptions" :key="item.value" :label="item.label"
                           :value="item.value"></el-option>
              </el-select>
            </el-col>
            <el-col :span="20">
              <el-input v-if="listConstraint(constraint)" autosize size="small" label="value"
                        type="textarea" v-model="constraint.value">
              </el-input>
              <el-input v-else size="small" v-model="constraint.value">
                <template slot="prepend">Value&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</template>
              </el-input>
            </el-col>
            <el-col :span="2">
            </el-col>
            <el-col :span="2">
              <el-button type="danger"
                         plain
                         class="width--full"
                         @click="deleteConstraint(segment, constraint)"
                         size="small">
                <i class="el-icon-delete"></i>
              </el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="card--empty" v-else>
        <span>No constraints (ALL will pass)</span>
      </div>
      <div>
        <el-row :gutter="3">
          <el-col :span="5">
            <el-input size="small" placeholder="Property"
                      v-model="segment.newConstraint.property"></el-input>
          </el-col>
          <el-col :span="4">
            <el-select size="small" v-model="segment.newConstraint.operator" placeholder="operator">
              <el-option v-for="item in operatorOptions" :key="item.value" :label="item.label"
                         :value="item.value"></el-option>
            </el-select>
          </el-col>
          <el-col :span="11">
            <el-input v-if="listConstraint(segment.newConstraint)" type="textarea" autosize
                      v-model="segment.newConstraint.value"></el-input>
            <el-input v-else size="small" v-model="segment.newConstraint.value"></el-input>
          </el-col>
          <el-col :span="4">
            <el-button class="width--full" size="small" type="primary" plain
                       :disabled="!segment.newConstraint.property || !segment.newConstraint.value"
                       @click.prevent="() => createConstraint(segment)">
              Add constraint
            </el-button>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import helpers from "@/helpers/helpers";
import clone from "lodash.clone";
import { operators } from "@/operators.json";
import ConstraintService from "@/services/constraintService";

const { handleErr } = helpers;

export default {
  name: "constraints",
  props: {
    segment: Object,
    flagId: Number
  },
  data() {
    return {
      operatorOptions: operators,
      dialogDeleteConstraintVisible: false,
      constraintToDelete: {},
    };
  },
  methods: {
    listConstraint(constraint) {
      return ConstraintService.listConstraints.includes(constraint.operator);
    },
    createConstraint() {
      const formattedConstraint = ConstraintService.formatConstraint(this.segment.newConstraint);
      ConstraintService.postConstraintPromise(this.flagId,this.segment.id, formattedConstraint)
          .then(response => {
            let constraint = response.data;
            this.segment.constraints.push(constraint);
            this.segment.newConstraint = clone(ConstraintService.DEFAULT_CONSTRAINT);
            this.$message.success("new constraint created");
          }, handleErr.bind(this));
    },
    deleteConstraint(segment, constraint) {
      this.dialogDeleteConstraintVisible = true;
      this.constraintToDelete = constraint;
    },
    cancelConstraintDeletion() {
      this.dialogDeleteConstraintVisible = false;
      this.constraintToDelete = {};
    },
    confirmConstraintDeletion() {
      ConstraintService.deleteConstraintPromise(this.flagId, this.segment.id, this.constraintToDelete.id
      ).then(() => {
        const index = this.segment.constraints.findIndex(
            c => c.id === this.constraintToDelete.id
        );
        this.segment.constraints.splice(index, 1);
        this.dialogDeleteConstraintVisible = false;
        this.constraintToDelete = {};
        this.$message.success("constraint deleted");
      }, handleErr.bind(this));
    },
  }
};
</script>

<!-- TODO: Port styles -->