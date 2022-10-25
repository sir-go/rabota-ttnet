<template>
  <Part
    title="Качества успешного сотрудника"
    caption="какими тремя качествами должен обладать успешный сотрудник,
    работающий в должности на которую Вы претендуете?"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input v-model="lv.p[partName]['качество ' + i]" label="должен..." autogrow clearable counter maxlength="1000"
             v-for="i in [1, 2, 3]" v-bind:key="i">
      <template v-slot:prepend><q-icon name="done" /></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'MustHavePerks',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
