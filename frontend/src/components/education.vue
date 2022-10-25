<template>
  <Part
    title="Образование"
    caption="уровень образования"
    :icon="icon"
    v-model="lv.enabledParts"
    no-b-pad
    :part-name="partName"
  >
    <q-card-section class="bg-grey-1 q-mt-sm q-mb-sm outlined"
                    v-for="i in lv.p[partName]" v-bind:key="i.id">
      <q-select
        label="вид"
        v-model="i['уровень']"
        :options="edu_types"
        map-options
        emit-value
      >
        <template v-slot:prepend><q-icon name="las la-brain" /></template>
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
      <q-input v-model="i['учебное заведение']" label="полное название учебного заведения"
               clearable counter maxlength="250"
      >
        <template v-slot:prepend><q-icon name="las la-university" /></template>
      </q-input>
      <q-input v-model="i['специальность']" label="специальность"
               clearable counter maxlength="100"
      >
        <template v-slot:prepend><q-icon name="las la-user-graduate" /></template>
      </q-input>
      <q-input
        v-model="i['год окончания']"
        :type="$q.platform.is.mobile ? 'number' : 'text'"
        placeholder="____ г."
        label="год окончания"
        mask="####"
        style="width: 10rem"
        clearable
      >
        <template v-slot:prepend><q-icon name="las la-calendar-check" /></template>
      </q-input>
      <div class="text-grey-8 absolute-top-right">
        <q-btn flat dense icon="close" @click="removeEdu(i)" />
      </div>
    </q-card-section>
    <q-card-actions>
      <q-btn flat align="right" color="grey-7" class="full-width" icon="las la-school"
             label="+ запись об образовании"
             @click="addEdu"
      />
    </q-card-actions>
  </Part>
</template>

<script>
const EDU_TYPES = [
  { icon: 'las la-hard-hat', label: 'среднее специальное', value: 'среднее специальное' },
  { icon: 'las la-graduation-cap', label: 'высшее', value: 'высшее' },
  { icon: 'las la-certificate', label: 'профессиональные курсы', value: 'профессиональные курсы' }
]
import Part from './part'
import edu from '../models/edu.json'
import { uid } from 'quasar'

export default {
  name: 'Education',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } },
  data () { return { edu_types: EDU_TYPES } },
  methods: {
    addEdu () {
      let item = this.$_.clone(edu)
      // noinspection NonAsciiCharacters
      this.$_.assign(item, { id: uid(), 'уровень': this.edu_types[0].value })
      this.lv.p[this.partName].push(item)
    },
    removeEdu (i) {
      this.lv.p[this.partName] = this.lv.p[this.partName].filter(_i => _i !== i)
    }
  }
}
</script>
