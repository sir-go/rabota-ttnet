<template>
  <Part
    title="Основная информация"
    :icon="icon"
    no-toggle
  >
    <q-input
      v-model="lv.p[partName]['источник']"
      label="* откуда Вы узнали о наборе?"
      clearable
      required
      bottom-slots
      counter
      maxlength="100"
    >
      <template v-slot:prepend><q-icon name="las la-share-alt-square" /></template>
    </q-input>

    <q-input
      v-model="lv.p[partName]['вакансия']"
      label="* вакансия"
      clearable
      required
      bottom-slots
      counter
      maxlength="100"
    >
      <template v-slot:prepend><q-icon name="las la-briefcase" /></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'MainData',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
