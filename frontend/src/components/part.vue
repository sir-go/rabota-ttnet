<template>
  <q-card bordered flat class="q-mb-sm">
    <q-item :tag="noToggle ? 'div' : 'label'" :v-ripple="!noToggle" class="bg-grey-2">
      <q-item-section avatar class="xs-hide">
        <q-icon size="lg" :name="icon" color="blue-grey-7" />
      </q-item-section>
      <q-item-section>
        <q-item-label class="text-h6">{{ title }}</q-item-label>
        <q-item-label caption>{{ caption }}</q-item-label>
      </q-item-section>
      <q-item-section side v-if="!noToggle">
        <q-item-label>
          <q-toggle :val="partName" keep-color color="primary" unchecked-icon="clear" checked-icon="check" v-model="lv"/>
        </q-item-label>
      </q-item-section>
    </q-item>
    <div v-if="noToggle || this.$_.includes(value, partName)"
         :class="['q-ml-lg', 'q-mr-lg', noBPad ? '' : 'q-mb-lg']">
      <slot></slot>
    </div>
  </q-card>
</template>

<script>
export default {
  name: 'Part',
  props: {
    title: String,
    caption: String,
    partName: String,
    value: Array,
    noBPad: Boolean,
    noToggle: Boolean,
    icon: String
  },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
