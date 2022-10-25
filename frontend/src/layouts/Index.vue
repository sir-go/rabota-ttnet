<template>
  <q-layout view="hHh LpR fFf">
    <q-header reveal elevated class="bg-grey-8 text-white cursor-pointer" v-ripple @click="scrollToTop">
      <q-toolbar>
        <q-toolbar-title class="col-12 text-center">
          {{ page.title }}
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container>
      <q-page padding class="row justify-center">
        <q-form @submit="submit" class="full-width" style="max-width: 40rem">
          <div class="full-width q-pa-lg">
            <p>
              Заполните анкету как можно подробней.
            </p>
            <p>Все поля, кроме помеченных звёздочкой *, не обязательные к
              заполнению, но если информации о Вас будут меньше, чем от других кандидатов выбор будет не в Вашу пользу.
            </p>
            <p>
              После заполнения анкеты я свяжусь с Вами повторно.
            </p>
          </div>
          <MainData v-model="form" part-name="основная информация" icon="las la-exclamation"/>
          <PersonalData v-model="form" part-name="основная информация" icon="las la-user-edit"/>
          <Passport v-model="form" part-name="паспорт" icon="las la-passport"/>
          <Family v-model="form" part-name="семья" icon="las la-users"/>
          <Education v-model="form" part-name="образование" icon="las la-graduation-cap"/>
          <ItSkills v-model="form" part-name="ит навыки" icon="las la-robot"/>
          <Auto v-model="form" part-name="авто" icon="las la-truck-pickup"/>
          <Hobbies v-model="form" part-name="увлечения" icon="las la-gamepad"/>
          <Duty v-model="form" part-name="армия" icon="las la-user-astronaut"/>
          <Experience v-model="form" part-name="опыт работы" icon="las la-history"/>
          <Criminal v-model="form" part-name="судимость" icon="las la-balance-scale"/>
          <Alcohol v-model="form" part-name="алкоголь" icon="las la-wine-bottle"/>
          <Smoking v-model="form" part-name="курение" icon="las la-joint"/>
          <Expectations v-model="form" part-name="ожидания" icon="las la-business-time"/>
          <Priorities v-model="form" part-name="приоритеты" icon="las la-stream"/>
          <MustHavePerks v-model="form" part-name="обязательные качества" icon="las la-thumbs-up"/>
          <ReasonsForHire v-model="form" part-name="причины нанять" icon="las la-user-ninja"/>
          <Achievements v-model="form" part-name="успехи" icon="las la-award"/>
          <Goals v-model="form" part-name="достижения" icon="las la-flag"/>
          <Leading v-model="form" part-name="инициативность" icon="las la-user-tie"/>
          <WorkConditions v-model="form" part-name="условия" icon="las la-handshake"/>
          <DPA
            v-model="form.DPA"
            :fullname="form.p['фио']"
            :s="form.p['паспорт']['серия']"
            :n="form.p['паспорт']['номер']"
            :w="form.p['паспорт']['кем выдан']"
            :a="form.p['адрес регистрации']"
            icon="las la-signature"
          />
          <div class="full-width q-pa-lg q-pb-xl">
            <q-btn :disabled="!formIsValid" type="submit" size="lg" color="primary" class="full-width">отправить</q-btn>
          </div>
        </q-form>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import { scroll, uid } from 'quasar'
import profile from '../models/profile.json'
import priors from '../models/priors.json'
import PersonalData from '../components/personal_data'
import MainData from '../components/main_data'
import Passport from '../components/passport'
import Family from '../components/family'
import Education from '../components/education'
import ItSkills from '../components/it_skills'
import Auto from '../components/auto'
import Hobbies from '../components/hobbies'
import Experience from '../components/experience'
import Criminal from '../components/criminal'
import Alcohol from '../components/alcohol'
import Smoking from '../components/smoking'
import Expectations from '../components/expectations'
import Priorities from '../components/priorities'
import MustHavePerks from '../components/must_have_perks'
import ReasonsForHire from '../components/reasons_for_hire'
import Achievements from '../components/achievements'
import Goals from '../components/goals'
import Leading from '../components/leading'
import WorkConditions from '../components/work_conditions'
import Duty from '../components/duty'
import DPA from '../components/DPA'

const { setScrollPosition } = scroll
const SERVER_TIMEOUT = 300

// noinspection JSUnusedGlobalSymbols
export default {
  name: 'MainLayout',
  components: {
    Duty,
    WorkConditions,
    Leading,
    Goals,
    Achievements,
    ReasonsForHire,
    MustHavePerks,
    Priorities,
    Expectations,
    Smoking,
    Alcohol,
    Criminal,
    Experience,
    Hobbies,
    Auto,
    ItSkills,
    Education,
    Family,
    Passport,
    MainData,
    PersonalData,
    DPA
  },
  data () {
    return {
      page: { title: 'Анкета соискателя' },
      form: {
        p: profile,
        enabledParts: ['основная информация'],
        DPA: {
          pd: true,
          s: true
        }
      }
    }
  },
  mounted () {
    this.parseModel()
    // this.getProfileData()
  },
  computed: {
    formIsValid () {
      return this.form.DPA.pd && this.form.DPA.s &&
        this.form.p['основная информация']['вакансия'] &&
        this.form.p['основная информация']['источник']
    }
  },
  methods: {
    idsAddArray (arr, arrNames) { arr.forEach(oi => this.idsAdd(oi, arrNames)) },
    idsAdd (o, arrNames) {
      if (this.$_.isEmpty(o)) return
      arrNames.forEach(an => {
        if (o.hasOwnProperty(an) && !this.$_.isEmpty(o[an])) {
          o[an].forEach(i => { i.id = uid() })
        }
      })
    },
    idsDelArray (arr, arrNames) { arr.forEach(oi => this.idsDel(oi, arrNames)) },
    idsDel (o, arrNames) {
      if (this.$_.isEmpty(o)) return
      arrNames.forEach(an => {
        if (o.hasOwnProperty(an) && !this.$_.isEmpty(o[an])) {
          o[an].forEach(i => { delete i.id })
        }
      })
    },
    isPartDisabled (partName) {
      return !this.$_.includes(this.form.enabledParts, partName)
    },
    isObjEmpty (o) {
      return this.$_.isObject(o) ? this.$_.values(o).every(this.isObjEmpty) : this.$_.values(o).every(this.$_.isEmpty)
    },
    modelFill () {
      let result = this.$_.cloneDeep(Object.freeze(this.form.p))
      this.idsDel(result, ['семья', 'образование', 'опыт работы'])
      this.idsDelArray(result['опыт работы'], ['рекомендации'])
      if (result['судимость']['под следствием'] === 'нет') {
        result['судимость']['под следствием'] = null
      }
      if (result['приоритеты'] === priors) {
        result['приоритеты'] = []
      }
      return this.$_.omitBy(result, (v, k) => this.isObjEmpty(v) || this.isPartDisabled(k))
    },
    parseModel () {
      this.form.p = this.$_.cloneDeep(profile)
      this.$_.keys(this.form.p).forEach(i => {
        if (!this.isObjEmpty(this.form.p[i])) {
          this.form.enabledParts.push(i)
        }
      })
      this.idsAdd(this.form.p, ['семья', 'образование', 'опыт работы'])
      this.idsAddArray(this.form.p['опыт работы'], ['рекомендации'])
      if (this.$_.isNull(this.form.p['судимость']['под следствием'])) {
        this.form.p['судимость']['под следствием'] = 'нет'
      }
      if (this.$_.isEmpty(this.form.p['приоритеты'])) {
        this.form.p['приоритеты'] = priors
      }
    },
    scrollToTop () { setScrollPosition(window, 0, 250) },
    reset () {
      this.parseModel()
      this.form.enabledParts = ['основная информация']
    },
    submit () {
      this.$axios.post('/api/v1/profile/',
        this.modelFill(this.form.p),
        { timeout: SERVER_TIMEOUT * 1000 })
        .then(() => {
          alert('Сведения успешно обновлены!')
          this.reset()
        })
        .catch((e) => {
          this.$q.notify({
            position: 'top',
            type: 'negative',
            message: e.toString()
          })
        })
        .finally(() => {
          this.$q.loading.hide()
          this.scrollToTop()
        })
    }
  }
}
</script>

<style lang="stylus">
  .dadata-input
    height 2.5rem
    padding 1rem 0 0 0
    margin -1rem 0 0 0
    background none

  .outlined
    outline 1px solid lightgrey
</style>
