<template>
  <Part
    title="Ваши ожидания должностного роста"
    caption="кем вы хотите стать через год, три, пять лет работы в фирме?"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input v-model="lv.p[partName]['через 1 год']" clearable counter maxlength="1000" autogrow
             label="через год...">
      <template v-slot:prepend><q-icon name="looks_one" /></template>
    </q-input>
    <q-input v-model="lv.p[partName]['через 3 года']" clearable counter maxlength="1000" autogrow
             label="через 3 года...">
      <template v-slot:prepend><q-icon name="looks_3" /></template>
    </q-input>
    <q-input v-model="lv.p[partName]['через 5 лет']" clearable counter maxlength="1000" autogrow
             label="через 5 лет...">
      <template v-slot:prepend><q-icon name="looks_5" /></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
// noinspection JSUnusedGlobalSymbols
export default {
  name: 'Expectations',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
