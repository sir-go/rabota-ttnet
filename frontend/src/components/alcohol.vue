<template>
  <Part
    title="Отношение к упортреблению алкоголя"
    caption="выберите оптимальный для Вас вариант"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-list>
      <q-item class="q-pl-none" tag="label" v-ripple v-for="i in [
        'считаю, что умеренное потребление вне производства не мешает успешной работе',
        'считаю, что человек имеет право «расслабиться» по своему усмотрению',
        'считаю, что употребление спиртного негативно отражается на профессиональной деятельности'
      ]" v-bind:key="i">
        <q-item-section avatar>
          <q-radio v-model="lv.p[partName]" :val="i" />
        </q-item-section>
        <q-item-section>
          <q-item-label>
            {{ i }}
          </q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
  </Part>
</template>

<!--suppress JSUnusedGlobalSymbols -->
<script>
import Part from './part'
export default {
  name: 'Alcohol',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
