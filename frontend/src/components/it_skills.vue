<template>
  <Part
    title="Владею ПК"
    caption="какими программными пакетами умеете пользоваться"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <div class="row justify-between q-mr-lg">
      <q-checkbox
        class="col-5"
        v-model="lv.p[partName]['офис']"
        :val="i"
        :label="i"
        v-for="i in ['Windows', 'MS Word', '1C', 'MS Excel']"
        v-bind:key="i"
      />
    </div>
    <q-input class="full-width" label="другое ПО"
             v-model="lv.p[partName]['прочее ПО']"
             autogrow clearable counter maxlength="250"
    />
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'ItSkills',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
