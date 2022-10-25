<template>
  <Part
    title="Личные данные"
    caption="Имя, контакты, адреса, возраст"
    :icon="icon"
    no-toggle
  >
    <div class="row justify-between">
      <q-input
        v-model="phone"
        label="* номер телефона"
        requared
        clearable
        mask="phone"
        type="tel"
        placeholder="(___)___-____"
        class="col-12 col-sm-5"
        bottom-slots
      ><template v-slot:prepend><q-icon name="las la-mobile" /></template></q-input>

      <q-input
        v-model="lv.p[partName]['email']"
        label="e-mail"
        clearable
        type="email"
        class="col-12 col-sm-6"
        bottom-slots
        counter
        maxlength="50"
      ><template v-slot:prepend><q-icon name="las la-at" /></template></q-input>
    </div>

    <q-field
        v-model="lv.p[partName]['фио']"
        label="* Фамилия Имя Отчество"
        clearable
        requared
        bottom-slots
        counter
        maxlength="100"
    >
      <template v-slot:prepend><q-icon name="las la-id-badge" /></template>
      <template v-slot:control>
        <dadata-suggestions
          class="self-center full-width no-outline no-border"
          v-model="lv.p[partName]['фио']"
          :options="{ type: 'NAME' }"
        />
      </template>
    </q-field>

    <q-field
        v-model="lv.p[partName]['адрес регистрации']"
        label="адрес прописки"
        clearable
        bottom-slots
        counter
        maxlength="250"
    >
      <template v-slot:prepend><q-icon name="las la-map-marked" /></template>
      <template v-slot:control>
        <dadata-suggestions
          class="self-center full-width no-outline no-border"
          v-model="lv.p[partName]['адрес регистрации']"
          field-value="unrestricted_value"
          :options="{ type: 'ADDRESS' }"
        />
      </template>
    </q-field>

    <q-field
        v-model="lv.p[partName]['адрес проживания']"
        label="адрес проживания"
        clearable
        bottom-slots
        counter
        maxlength="250"
    >
      <template v-slot:prepend><q-icon name="las la-home" /></template>

      <template v-slot:control>
        <dadata-suggestions
          class="self-center full-width no-outline no-border"
          v-model="lv.p[partName]['адрес проживания']"
          field-value="unrestricted_value"
          :options="{ type: 'ADDRESS' }"
        />
      </template>
    </q-field>

    <div class="row justify-between">
      <q-input
        v-model="lv.p[partName]['дата рождения']"
        mask="##.##.####"
        placeholder="__ . __ . ____ г."
        label="дата рождения"
        class="col-5 col-xs-12 col-sm-5"
        clearable
      >
        <template v-slot:prepend><q-icon name="las la-calendar-day" /></template>
      </q-input>
    </div>
  </Part>
</template>

<script>
import Part from './part'
export default {
  name: 'PersonalData',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: {
    lv: { get () { return this.value }, set (v) { this.$emit('input', v) } },
    phone: {
      get () {
        return this.lv.p[this.partName]['телефон']
      },
      set (v) {
        this.lv.p[this.partName]['телефон'] = this.$_.isString(v) ? v.replace(/\D/g, '') : null
      }
    }
  }
}
</script>
