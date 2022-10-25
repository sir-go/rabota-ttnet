<template>
  <Part
    title="Опыт работы"
    caption="места работы, должности, задачи"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
    no-b-pad
  >
    <q-card-section class="bg-grey-1 q-mt-sm q-mb-sm outlined" v-for="i in lv.p[partName]" v-bind:key="i.id">
      <div class="row justify-between">
        <q-input
          v-model="i['дата трудоустройства']"
          mask="##.##.####"
          placeholder="__ . __ . ____ г."
          label="дата трудоустройства"
          class="col-12 col-sm-5"
          clearable
        >
          <template v-slot:prepend><q-icon name="las la-calendar-check"/></template>
        </q-input>
        <q-input
          v-model="i['дата увольнения']"
          placeholder="__ . __ . ____ г."
          mask="##.##.####"
          label="дата увольнения"
          class="col-12 col-sm-5"
          clearable
        >
          <template v-slot:prepend><q-icon name="las la-calendar-times"/></template>
        </q-input>
      </div>
      <q-input
        v-model="i['название организации']"
        label="название организации"
        clearable counter maxlength="250"
      >
        <template v-slot:prepend><q-icon name="las la-industry"/></template>
      </q-input>
      <div class="row justify-between">
        <q-input
          class="col-12 col-sm-7"
          label="должность"
          v-model="i['должность']"
          autogrow clearable counter maxlength="100"
        ><template v-slot:prepend><q-icon name="las la-user-tie"/></template></q-input>
        <q-input
          v-model="i['зарплата']"
          label="з/п (р/месяц)"
          :type="$q.platform.is.mobile ? 'number' : 'text'"
          mask="######"
          class="col-12 col-sm-4"
          clearable
          bottom-slots
        >
          <template v-slot:prepend><q-icon name="las la-coins"/></template>
        </q-input>
      </div>
      <q-input
        class="full-width"
        label="основные должностные обязанности"
        v-model="i['обязанности']"
        autogrow clearable counter maxlength="1000"
      ><template v-slot:prepend><q-icon name="las la-tasks"/></template></q-input>
      <q-input
        v-if="i['дата увольнения']"
        class="full-width"
        label="причина увольнения"
        v-model="i['причина увольнения']"
        autogrow clearable counter maxlength="1000"
      >
        <template v-slot:prepend><q-icon name="help_outline" /></template>
      </q-input>

      <div class="text-h6 text-grey-9 q-mt-sm q-mb-sm">меня могут рекомендовать:</div>
      <q-card-section class="q-mb-sm bg-white outlined" v-for="r in i['рекомендации']" v-bind:key="r.id">
        <q-field v-model="r['фио']" label="Фамилия Имя Отчество" clearable counter maxlength="100">
          <template v-slot:prepend><q-icon name="las la-id-badge" /></template>
          <template v-slot:control>
            <dadata-suggestions
              class="self-center no-outline no-border"
              v-model="r['фио']"
              field-value="unrestricted_value"
              :options="{ type: 'NAME' }"
            />
          </template>
        </q-field>
        <div class="row justify-between">
          <q-input
            class="col-12 col-sm-6"
            v-model="r['должность']" label="должность" clearable counter maxlength="100">
            <template v-slot:prepend><q-icon name="las la-briefcase" /></template>
          </q-input>
          <q-input
            v-model="r['контакты']"
            label="контакты"
            class="col-12 col-sm-5"
            clearable counter maxlength="100"
          >
            <template v-slot:prepend><q-icon name="las la-mobile" /></template>
          </q-input>
        </div>
        <div class="text-grey-8 absolute-top-right">
          <q-btn flat dense icon="close" @click="removeExpRecommend(i, r)" />
        </div>
      </q-card-section>
      <q-btn flat align="right" color="grey-7" class="full-width" icon="las la-user-plus"
             label="добавить контакт"
             @click="addExpRecommend(i)"
      />
      <div class="text-grey-8 absolute-top-right">
        <q-btn flat dense icon="close" @click="removeExp(i)" />
      </div>
    </q-card-section>
    <q-card-actions>
      <q-btn flat align="right" color="grey-7" class="full-width" icon="domain"
             label="+ добавить место работы"
             @click="addExp"
      />
    </q-card-actions>
  </Part>
</template>

<script>
import Part from './part'
import exp from '../models/exp.json'
import recs from '../models/recs.json'
import { uid } from 'quasar'

// noinspection JSUnusedGlobalSymbols
export default {
  name: 'Experience',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: {
    lv: { get () { return this.value }, set (v) { this.$emit('input', v) } }
  },
  methods: {
    addExp () {
      let item = this.$_.clone(exp)
      item.id = uid()
      this.lv.p[this.partName].push(item)
    },
    removeExp (i) {
      this.lv.p[this.partName] = this.lv.p[this.partName].filter(_i => _i !== i)
    },
    addExpRecommend (expItem) {
      let item = this.$_.clone(recs)
      item.id = uid()
      expItem['рекомендации'].push(item)
    },
    removeExpRecommend (expItem, i) {
      expItem['рекомендации'] = expItem['рекомендации'].filter(_i => _i !== i)
    }
  }
}
</script>
