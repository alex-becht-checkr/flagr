<template>
  <div>
    <el-dialog title="Delete segment" :visible.sync="dialogDeleteSegmentVisible">
      <span>Are you sure you want to delete this segment?</span>
      <span slot="footer" class="dialog-footer">
              <el-button @click="cancelSegmentDeletion">Cancel</el-button>
              <el-button type="primary" @click="confirmSegmentDeletion">Confirm</el-button>
            </span>
    </el-dialog>

    <el-dialog title="Create segment" :visible.sync="dialogCreateSegmentOpen">
      <div>
        <p>
          <el-input placeholder="Segment description" v-model="newSegment.description"></el-input>
        </p>
        <p>
          <el-slider v-model="newSegment.rolloutPercent" show-input></el-slider>
        </p>
        <el-button class="width--full"
                   :disabled="!newSegment.description"
                   @click.prevent="createSegment">
          Create segment
        </el-button>
      </div>
    </el-dialog>

    <el-card class="segments-container base-card">
      <div slot="header" class="el-card-header">
        <div class="flex-row">
          <div class="flex-row">
            <h2>Segments
              <el-tooltip placement="bottom" effect="light">
                <div slot="content">How each variant is being evaluated.<br>
                  A variant must match all constraints<br>
                  within a segment or it will be passed<br>
                  to the next segment.</div>
                <span class="el-icon-info" style="vertical-align: middle; font-size: 1rem;"></span>
              </el-tooltip>
            </h2>
          </div>
          <div class="flex-row-right">
            <el-button @click="dialogCreateSegmentOpen = true">New segment</el-button>
          </div>
        </div>
      </div>
      <div class="segments-container-inner" v-if="segments.length">
        <el-card
            shadow="never"
            v-for="segment in segments"
            :key="segment.id"
            class="segment"
        >
          <div class="flex-row id-row">
            <div class="flex-row-left">
              <el-button size="small" @click="segmentUp(segment, segments)">
                          <span class="el-icon-arrow-up">
                          </span>
              </el-button>
              <el-button size="small" @click="segmentDown(segment, segments)" style="margin-right: 15px">
                          <span class="el-icon-arrow-down">
                          </span>
              </el-button>
              <el-tag type="primary" :disable-transitions="true">
                Segment ID:
                <b>{{ segment.id }}</b>
              </el-tag>
            </div>
            <div class="flex-row-right">
              <el-button size="small" @click="cloneSegment(segment)">Clone segment</el-button>
              <el-button @click="deleteSegment(segment)" size="small">
                <span class="el-icon-delete" />
              </el-button>
            </div>
          </div>
          <el-row :gutter="10" class="id-row">
            <el-col :span="15">
              <el-input size="small" placeholder="Description" v-model="segment.description">
                <template slot="prepend">Description</template>
              </el-input>
            </el-col>
            <el-col :span="9">
              <el-input class="segment-rollout-percent" size="small" placeholder="0"
                        v-model="segment.rolloutPercent" :min="0" :max="100">
                <template slot="prepend">Rollout</template>
                <template slot="append">%</template>
              </el-input>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <Constraints :flag-id="flagId"
                           :segment="segment"/>
            </el-col>
            <el-col :span="24">
              <h5>
                <span>Distribution</span>
                <el-tooltip placement="bottom" effect="light">
                  <div slot="content">The possibility (in percentage) a specific<br>
                    variant will be served to a member of that<br>
                    segment if all its constraints are met.</div>
                  <span class="el-icon-info" style="margin-left: 5px; vertical-align: middle;"></span>
                </el-tooltip>
              </h5>
            </el-col>
          </el-row>
          <el-row v-for="variant in variants" :key="variant.id" :gutter="20">
            <el-col :span="24" class="segment-distribution">
              <el-input class="variant-key-input segment-distribution-variant" size="small"
                        placeholder="rollout" type="number" :min="0" :max="100"
                        :value="getDistributionValue(segment, variant)"
                        @input="value => setDistributionValue(value, segment, variant)">
                <template slot="prepend">{{ variant.key }}</template>
                <template slot="append">%</template>
              </el-input>
              <el-progress color="#74E5E0"
                           :percentage="Math.max(0, Math.min(100, getDistributionValue(segment, variant)))"></el-progress>
            </el-col>
          </el-row>
          <el-row v-if="sumDistributions(segment.distributions) !== 100">
            <el-col :span="24" class="segment-distribution-alert">
              <el-alert  class="edit-distribution-alert"
                         :title="'Percentages must add up to 100% (currently at ' +
                                sumDistributions(segment.distributions) +
                                '%)'
                                " type="error" :closable="false" show-icon></el-alert>
            </el-col>
          </el-row>
        </el-card>
      </div>
      <div class="card--error" v-else>No segments created for this feature flag yet</div>
    </el-card>
  </div>
</template>
<script>
import Constraints from "@/components/Flag/Constraints.vue";
import helpers from "@/helpers/helpers";
import SegmentService from "@/services/segmentService";
import clone from "lodash.clone";
import ConstraintService from "@/services/constraintService";

const { handleErr } = helpers;

function processSegment(segment) {
  segment.newConstraint = clone(ConstraintService.DEFAULT_CONSTRAINT);
}

export default {
  name: "segments",
  props: {
    flagId: Number,
    segments: Array,
    variants: Array
  },
  emits: ["fetchFlag"],
  components: {Constraints},
  data() {
    return {
      dialogCreateSegmentOpen: false,
      dialogDeleteSegmentVisible: false,
      segmentToDelete: {},
      newSegment: clone(SegmentService.DEFAULT_SEGMENT),
    };
  },
  methods: {
    deleteSegment(segment) {
      this.dialogDeleteSegmentVisible = true;
      this.segmentToDelete = segment;
    },
    cancelSegmentDeletion() {
      this.dialogDeleteSegmentVisible = false;
      this.segmentToDelete = {};
    },
    confirmSegmentDeletion() {
      SegmentService.deleteSegmentPromise(
          this.flagId,
          this.segmentToDelete.id
      ).catch(handleErr.bind(this)).then(() => {
        const index = this.segments.findIndex(el => el.id === this.segmentToDelete.id);
        this.segments.splice(index, 1);
        this.dialogDeleteSegmentVisible = false;
        this.segmentToDelete = {};
        this.$message.success("segment deleted");
      });
    },
    sumDistributions(distributions) {
      return distributions.reduce((current, { percent }) => current + percent, 0);
    },
    getDistributionValue(segment, variant) {
      const distribution = segment.distributions.find(d => d.variantID === variant.id);
      return distribution ? distribution.percent : 0;
    },
    setDistributionValue(value, segment, variant) {
      SegmentService.setDistributionValue(value, segment, variant);
    },
    cloneSegment(segment) {
      const errorHandler = handleErr.bind(this);
      // create segment first
      SegmentService.postSegmentPromise(
          this.flagId,
          `${segment.description} (Clone)`,
          segment.rolloutPercent
      ).then(response => {
        let newSegment = response.data;
        // clone distribution
        const distributions = segment.distributions.map(({ percent, variantID, variantKey }) => ({ percent, variantID, variantKey }));
        this.putDistributions(newSegment.id, distributions).then(async () => {
          // clone constraints sequentially to preserve order
          for (const constraint of segment.constraints) {
            await ConstraintService.postConstraintPromise(this.flagId, newSegment.id, constraint).catch(handleErr.bind(this));
          }
          this.$message.success('Segment successfully cloned')
          // Re fetch flag to update state
          this.$emit("fetch-flag");
        }, errorHandler);
      }, errorHandler);
    },
    putDistributions(segmentId, distributions) {
      return SegmentService.putDistributionsPromise(this.flagId, segmentId, distributions).catch(handleErr.bind(this));
    },
    segmentUp(segment, segments) {
      const segmentIndex = segments.findIndex(s => s.id === segment.id);

      if (segmentIndex < 1)
        return;

      segments.splice(segmentIndex - 1, 0, segments.splice(segmentIndex, 1)[0]);
    },
    segmentDown(segment, segments) {
      const segmentIndex = segments.findIndex(s => s.id === segment.id);

      if (segmentIndex > segments.length - 1)
        return;

      segments.splice(segmentIndex + 1, 0, segments.splice(segmentIndex, 1)[0]);
    },
    createSegment() {
      SegmentService.postSegmentPromise(
          this.flagId,
          this.newSegment.description,
          this.newSegment.rolloutPercent
      ).then(response => {
        let segment = response.data;
        processSegment(segment);
        segment.constraints = [];
        if (this.variants.length) {
          // default the first variant to 100%
          this.setDistributionValue(100, segment, this.variants[0]);
        }
        this.newSegment = clone(SegmentService.DEFAULT_SEGMENT);
        this.segments.push(segment);
        this.$message.success("new segment created");
        this.dialogCreateSegmentOpen = false;
      }, handleErr.bind(this));
    }
  }
}
</script>