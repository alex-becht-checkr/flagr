<template>
  <div>
    <el-dialog title="Delete variant" :visible.sync="dialogDeleteVariantVisible">
      <span>{{`Are you sure you want to delete variant #${variantToDelete.id} [${variantToDelete.key}]`}}</span>
      <span slot="footer" class="dialog-footer">
            <el-button @click="cancelVariantDeletion">Cancel</el-button>
            <el-button type="primary" @click="confirmVariantDeletion">Confirm</el-button>
          </span>
    </el-dialog>

    <el-card class="variants-container base-card">
      <div slot="header" class="clearfix">
        <h2>Variants
          <el-tooltip placement="bottom" effect="light">
            <div slot="content">Potential outputs being evaluated<br>
              by this flag.
            </div>
            <span class="el-icon-info" style="vertical-align: middle; font-size: 1rem;"></span>
          </el-tooltip>
        </h2>
      </div>
      <div class="variants-container-inner" v-if="variants.length">
        <div v-for="variant in variants" :key="variant.id">
          <el-card shadow="never">
            <el-form label-position="left" label-width="100px">
              <div class="flex-row id-row">
                <el-tag type="primary" :disable-transitions="true">
                  Variant ID:
                  <b>{{ variant.id }}</b>
                </el-tag>
                <el-input class="variant-key-input" size="small" placeholder="Key" v-model="variant.key">
                  <template slot="prepend">Key</template>
                </el-input>
                <div class="flex-row-right save-remove-variant-row">
                  <el-button @click="deleteVariant(variant)" size="small">
                    <span class="el-icon-delete" />
                  </el-button>
                </div>
              </div>
              <el-collapse class="flex-row">
                <el-collapse-item title="Variant attachment" class="variant-attachment-collapsable-title">
                  <p class="variant-attachment-title">You can add JSON in key/value pairs format.</p>
                  <vue-json-editor v-model="variant.attachment" :showBtns="false" :mode="'code'"
                                   v-on:has-error="variant.attachmentValid = false" v-on:input="variant.attachmentValid = true"
                                   class="variant-attachment-content"></vue-json-editor>
                </el-collapse-item>
              </el-collapse>
            </el-form>
          </el-card>
        </div>
      </div>
      <div class="card--error" v-else>No variants created for this feature flag yet</div>
      <div class="variants-input">
        <div class="flex-row equal-width constraints-inputs-container">
          <div>
            <el-input placeholder="Variant key" v-model="newVariant.key"></el-input>
          </div>
        </div>
        <el-button class="width--full" :disabled="!newVariant.key" @click.prevent="createVariant">Create
          variant</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
import VariantService from "@/services/variantService";
import helpers from "@/helpers/helpers";
import clone from "lodash.clone";
import vueJsonEditor from "vue-json-editor";

const { handleErr } = helpers;

export default {
  name: "variants",
  props: {
    flagId: Number,
    variants: Array,
    segments: Array
  },
  emits: ["fetchFlag"],
  components: {vueJsonEditor},
  data() {
    return {
      dialogDeleteVariantVisible: false,
      variantToDelete: {},
      newVariant: clone(VariantService.DEFAULT_VARIANT),
    };
  },
  methods: {
    cancelVariantDeletion() {
      this.dialogDeleteVariantVisible = false;
      this.variantToDelete = {};
    },
    confirmVariantDeletion() {
      VariantService.deleteVariantPromise(this.flagId, this.variantToDelete.id).then(() => {
        this.variantToDelete = {};
        this.dialogDeleteVariantVisible = false;
        this.$message.success("variant deleted");
        this.$emit("fetch-flag");
      }, handleErr.bind(this));
    },
    createVariant() {
      VariantService.postVariantPromise(this.flagId, this.newVariant).then(response => {
        let variant = response.data;
        this.newVariant = clone(VariantService.DEFAULT_VARIANT);
        this.variants.push(variant);
        this.$message.success("new variant created");
      }, handleErr.bind(this));
    },
    deleteVariant(variant) {
      const isVariantInUse = this.segments.some(segment =>
          segment.distributions.some(
              distribution => distribution.variantID === variant.id
          )
      );

      if (isVariantInUse) {
        this.$message.error("This variant is being used by a segment distribution. Please remove the segment or edit the distribution in order to remove this variant.");

        return;
      }

      this.dialogDeleteVariantVisible = true;
      this.variantToDelete = variant;
    },
  }
}
</script>