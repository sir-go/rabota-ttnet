<template>
  <Part
    title="Авто"
    caption="автотранспорт, данные ВУ"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <div class="row justify-between">
      <q-input
        v-model="lv.p[partName]['марка']"
        label="марка"
        clearable counter maxlength="50"
        class="col-12 col-sm-4"
      >
        <template v-slot:prepend><q-icon name="las la-car" /></template>
      </q-input>
      <q-input
        v-model="lv.p[partName]['год выпуска']"
        label="год выпуска"
        placeholder="____ г."
        mask="####"
        clearable bottom-slots
        class="col-12 col-sm-3"
      >
        <template v-slot:prepend><q-icon name="las la-calendar-check" /></template>
      </q-input>
      <q-input
        v-model="lv.p[partName]['гос. номер']"
        label="гос. номер"
        clearable bottom-slots counter maxlength="12"
        class="col-12 col-sm-4"
      >
        <template v-slot:prepend><q-icon name="las la-barcode" /></template>
      </q-input>
    </div>
    <div class="row justify-between">
      <q-field borderless label="категории ВУ" stack-label class="col-12 col-sm-6">
        <template v-slot:control>
          <div class="self-center full-width no-outline row justify-between" tabindex="0">
            <q-checkbox v-model="categories"
                        v-for="i in ['A','B','C','D', 'E']" v-bind:key="i"
                        :val="i" :label="i" />
          </div>
        </template>
      </q-field>
      <q-input
        v-model="lv.p[partName]['стаж вождения']"
        label="водительский стаж"
        clearable
        class="col-12 col-sm-4"
        counter maxlength="50"
      >
        <template v-slot:prepend><q-icon name="departure_board" /></template>
      </q-input>
    </div>
  </Part>
</template>

<script>

import Part from './part'
export default {
  name: 'Auto',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: {
    lv: { get () { return this.value }, set (v) { this.$emit('input', v) } },
    categories: {
      get () {
        return this.$_.split(this.value.p[this.partName]['категории ВУ'], '')
      },
      set (v) { this.$_.split(this.value.p[this.partName]['категории ВУ'] = this.$_.join(v.sort(), '')) }
    }
  }
}
</script>
