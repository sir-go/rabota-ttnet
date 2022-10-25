<template>
  <Part
    title="Хобби"
    caption="опишите Ваши интересы и увлечения"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input class="full-width" v-model="lv.p[partName]" autogrow clearable
             counter maxlength="1000"
    ><template v-slot:prepend><q-icon name="sort" /></template></q-input>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'Hobbies',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
