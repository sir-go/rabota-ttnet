<template>
  <Part
    title="Военная служба"
    caption="где и когда служили, если нет, то почему"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <div class="row justify-between">
      <q-select
        label="полнота прохождения"
        v-model="lv.p[partName]['полнота прохождения']"
        :options="[
          { icon: 'las la-calendar-check', label: 'в полном объёме', value: 'в полном объёме' },
          { icon: 'las la-calendar-minus', label: 'не в полном объёме', value: 'не в полном объёме' },
          { icon: 'las la-calendar-times', label: 'не призывался', value: 'не призывался' }
        ]"
        map-options
        emit-value
        class="col-12 col-sm-5"
      >
        <template v-slot:prepend><q-icon name="las la-tasks" /></template>
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
      <q-input v-if="lv.p[partName]['полнота прохождения'] === 'не призывался'"
               class="full-width" label="причина"
               v-model="lv.p[partName]['причина не службы']"
               autogrow clearable
               bottom-slots counter maxlength="1000"
      >
        <template v-slot:prepend><q-icon name="las la-question-circle" /></template>
      </q-input>
      <q-input
        v-if="lv.p[partName]['полнота прохождения'] &&
        lv.p[partName]['полнота прохождения'] !== 'не призывался'"
        v-model="lv.p[partName]['период']"
        label="период прохождения службы"
        placeholder="год призыва - год дмб"
        mask="#### г. - #### г."
        clearable
        class="col-12 col-sm-6"
      >
        <template v-slot:prepend><q-icon name="las la-calendar-check" /></template>
      </q-input>
      <q-input
        v-if="lv.p[partName]['полнота прохождения'] &&
        lv.p[partName]['полнота прохождения'] !== 'не призывался'"
        class="full-width" label="состою на учёте в РВК"
        v-model="lv.p[partName]['РВК']" clearable
        bottom-slots counter maxlength="100"
      >
        <template v-slot:prepend><q-icon name="las la-landmark" /></template>
      </q-input>
    </div>
  </Part>
</template>

<script>
import Part from './part'
// noinspection JSUnusedGlobalSymbols
export default {
  name: 'Duty',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
