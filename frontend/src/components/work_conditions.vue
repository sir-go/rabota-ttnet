<template>
  <Part
    title="Условия труда"
    caption="з/п, рабочий день, пр. условия"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-input
      label="ожидаемая Вами з/п"
      v-model="lv.p[partName]['ожидаемая зарплата']"
      clearable
      counter
      maxlength="100"
    >
      <template v-slot:prepend><q-icon name="las la-coins"/></template>
    </q-input>
    <q-select
      label="Ненорм. рабочий день"
      v-model="lv.p[partName]['приемлемы переработки']"
      map-options
      emit-value
      true-value="да"
      false-value="нет"
      :options="[
        { label: 'приемлем', value: 'да', icon: 'las la-thumbs-up' },
        { label: 'недопустим', value: 'нет', icon: 'las la-thumbs-down' }
      ]"
      clearable
    >
      <template v-slot:prepend><q-icon name="las la-business-time" /></template>
      <template v-slot:option="scope">
        <q-item
          v-bind="scope.itemProps"
          v-on="scope.itemEvents"
        >
          <q-item-section avatar>
            <q-icon :name="scope.opt.icon" />
          </q-item-section>
          <q-item-section>
            <q-item-label v-html="scope.opt.label" />
            <q-item-label caption>{{ scope.opt.description }}</q-item-label>
          </q-item-section>
        </q-item>
      </template>
    </q-select>
    <q-input
      label="особые условия"
      v-model="lv.p[partName]['особые условия']" autogrow clearable counter maxlength="1000">
      <template v-slot:prepend><q-icon name="las la-universal-access"/></template>
    </q-input>
    <q-input
      label="когда сможете приступить"
      v-model="lv.p[partName]['начало стажировки']" clearable counter maxlength="250">
      <template v-slot:prepend><q-icon name="las la-calendar-day"/></template>
    </q-input>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'WorkConditions',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
