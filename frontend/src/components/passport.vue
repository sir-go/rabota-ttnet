<template>
  <Part
    title="Паспортные данные"
    caption="Имя, контакты, адреса, возраст"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <div class="row justify-between">
      <q-input
        v-model="lv.p[partName]['серия и номер']"
        label="серия и номер"
        mask="## ## ######"
        placeholder="__  __  _____"
        clearable
        class="col-12 col-sm-7"
      >
        <template v-slot:prepend><q-icon name="las la-hashtag" /></template>
      </q-input>

      <q-input
          v-model="lv.p[partName]['дата выдачи']"
        mask="##.##.####"
        placeholder="__ . __ . ____ г."
        label="дата выдачи"
        class="col-5 col-xs-12 col-sm-4"
        clearable
      >
        <template v-slot:prepend><q-icon name="las la-calendar-day" /></template>
      </q-input>
    </div>

    <q-field
      v-model="lv.p[partName]['кем выдан']"
      label="кем выдан"
      clearable
      bottom-slots
      counter
      maxlength="250"
    >
      <template v-slot:prepend><q-icon name="las la-landmark" /></template>
      <template v-slot:control>
        <dadata-suggestions
          class="self-center full-width no-outline no-border"
          v-model="lv.p[partName]['кем выдан']"
          field-value="unrestricted_value"
          :options="{ type: 'fms_unit' }"
          maxlength="250"
        />
      </template>
    </q-field>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'Passport',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
