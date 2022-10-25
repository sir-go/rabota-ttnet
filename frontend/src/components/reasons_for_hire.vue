<template>
  <Part
    title="Почему подходите именно Вы?"
    caption="каковы три основные причины, по которым Вы должны
    быть приняты на эту должность? Что отличает Вас от других?"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input
      v-model="lv.p[partName]['причина ' + i]"
      :label="'причина ' + i"
      autogrow
      clearable
      v-for="i in [1, 2, 3]"
      v-bind:key="i"
      counter
      maxlength="1000"
    >
      <template v-slot:prepend><q-icon name="las la-hand-point-right" /></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
// noinspection JSUnusedGlobalSymbols
export default {
  name: 'ReasonsForHire',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
