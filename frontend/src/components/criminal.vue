<template>
  <Part
    title="Имею судимость"
    caption="статья, под следствием ли сейчас"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <div class="row justify-between">
      <q-input
        v-model="lv.p[partName]['статья']"
        label="статья"
        class="col-12 col-sm-7"
        clearable counter maxlength="100"
      >
        <template v-slot:prepend><q-icon name="las la-paragraph" /></template>
      </q-input>
      <q-checkbox
        true-value="да"
        false-value="нет"
        v-model="lv.p[partName]['под следствием']"
        label="нахожусь под следствием"
      />
    </div>
  </Part>
</template>

<script>
import Part from './part'
// noinspection JSUnusedGlobalSymbols
export default {
  name: 'Criminal',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
