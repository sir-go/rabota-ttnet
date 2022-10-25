<template>
  <Part
    title="Отношение к курению"
    caption="выберите оптимальный для Вас вариант"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
  >
    <q-list>
      <q-item class="q-pl-none" tag="label" v-ripple v-for="i in [
        'я не курю и считаю это негативным явлением',
        'курю умеренно, но это не мешает работе',
        'могу закурить, если для этого есть настроение'
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
  name: 'Smoking',
  components: { Part },
  props: { icon: String, value: Object, partName: String },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } }
}
</script>
