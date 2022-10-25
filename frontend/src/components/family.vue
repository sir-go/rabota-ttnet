<template>
  <Part
    title="Семья"
    caption="супруги, родители, дети - имена, место работы"
    :icon="icon"
    v-model="lv.enabledParts"
    no-b-pad
    :part-name="partName"
  >
    <q-card-section class="bg-grey-1 q-mt-sm q-mb-sm outlined"
                    v-for="fm in lv.p[partName]"
                    v-bind:key="fm.id"
    >
      <div  class="row justify-between">
        <q-select
          label="родство"
          v-model="fm['родство']"
          :options="familyMemberTypes"
          map-options
          emit-value
          bottom-slots
          class="col-3 col-xs-12 col-sm-3"
        >
          <template v-slot:prepend><q-icon name="group" /></template>
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

        <q-field
          v-model="fm['фио']" label="Фамилия Имя Отчество" clearable
          class="col-8 col-xs-12 col-sm-8"
          bottom-slots
          counter
          maxlength="100"
        >
          <template v-slot:prepend><q-icon name="las la-id-badge" /></template>
          <template v-slot:control>
            <dadata-suggestions
              class="self-center full-width no-outline no-border"
              v-model="fm['фио']"
              field-value="unrestricted_value"
              :options="{ type: 'NAME' }"
              maxlength="100"
            />
          </template>
        </q-field>
      </div>
      <div v-if="fm['родство'] !== 'ребёнок'">
        <q-input
          v-model="fm['мето работы']" label="место работы" clearable
          counter maxlength="250"
        >
          <template v-slot:prepend><q-icon name="las la-industry" /></template>
        </q-input>
        <q-input
          v-model="fm['должность']" label="должность" clearable
          counter maxlength="100">
          <template v-slot:prepend><q-icon name="las la-briefcase" /></template>
        </q-input>
      </div>
      <q-input v-else
               type="number" min="0" max="75"
               mask="##"
               v-model="fm['возраст']" label="возраст (лет)"
               style="width: 9rem" clearable>
        <template v-slot:prepend><q-icon name="child_care" /></template>
      </q-input>
      <div class="text-grey-8 absolute-top-right">
        <q-btn flat dense icon="close" @click="removeFamilyMember(fm)" />
      </div>
    </q-card-section>
    <q-card-actions>
      <q-btn flat align="right" color="grey-7" class="full-width" icon="las la-user-plus"
             label="добавить члена семьи"
             @click="addFamilyMember"
      />
    </q-card-actions>
  </Part>
</template>

<script>
const FAMILY_MEMBER_TYPES = [
  { icon: 'las la-restroom', value: 'супруг(а)', label: 'супруг(а)' },
  { icon: 'las la-female', value: 'мать', label: 'мать' },
  { icon: 'las la-male', value: 'отец', label: 'отец' },
  { icon: 'child_care', value: 'ребёнок', label: 'ребёнок' }
]
import Part from './part'
import family from '../models/family.json'
import { uid } from 'quasar'

export default {
  name: 'Family',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: {
    lv: { get () { return this.value }, set (v) { this.$emit('input', v) } },
    familyMemberTypes () {
      return FAMILY_MEMBER_TYPES.filter(FT => {
        return this.lv.p[this.partName].filter(f => {
          return FT.value !== 'ребёнок' && FT.value === f['родство']
        }).length < 1
      })
    }
  },
  methods: {
    addFamilyMember () {
      let item = this.$_.clone(family)
      // noinspection NonAsciiCharacters
      this.$_.assign(item, { id: uid(), 'родство': this.familyMemberTypes[0].value })
      this.lv.p[this.partName].push(item)
    },
    removeFamilyMember (i) {
      this.lv.p[this.partName] = this.lv.p[this.partName].filter(_i => _i !== i)
    }
  }
}
</script>
