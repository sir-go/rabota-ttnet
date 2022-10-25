<template>
  <Part
    title="Ваши успехи"
    caption="Что полезного Вы извлекли из опыта на своем последнем месте работы?
    Опишите Ваше продвижение на последнем рабочем месте"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input v-model="lv.p[partName]" autogrow clearable counter maxlength="5000">
      <template v-slot:prepend><q-icon name="sort" /></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'Achievements',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
