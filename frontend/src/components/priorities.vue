<template>
  <Part
    title="Приоритеты при выборе работы"
    caption="расставьте факторы в зависимости от того, что для вас важнее при выборе работы"
    :icon="icon"
    v-model="lv.enabledParts"
    :part-name="partName"
    ref="priorEl"
  >
    <Container
      @drop="onPriorityDrop"
      lock-axis="y"
      drag-handle-selector=".drag-handler"
      drag-class="dragged"
      :get-child-payload="idx => lv.p[partName][idx]"
    >
      <Draggable v-for="i in lv.p[partName]" v-bind:key="i" class="q-mt-sm q-mb-sm bg-grey-1 outlined">
        <div class="row items-center justify-between q-pt-xs q-pa-xs">
          <q-icon size="sm" color="grey-8" :name="icons[i]" class="col-shrink" />
          <div class="col-9">{{ i }}</div>
          <q-icon size="sm" color="grey-5" name="drag_indicator" class="col-shrink drag-handler"/>
        </div>
      </Draggable>
    </Container>
  </Part>
</template>

<script>
import { Container, Draggable } from 'vue-smooth-dnd'
import Part from './part'

// noinspection JSUnusedGlobalSymbols,NonAsciiCharacters
const IconByLabel = {
  'карьера': 'business_center',
  'деньги, льготы': 'euro',
  'близость к дому': 'home_work',
  'приобретение нового опыта': 'turned_in',
  'престиж компании': 'grade',
  'стабильность, надежность': 'verified_user',
  'самостоятельность и ответственность позиции': 'directions_boat',
  'практика по специальности': 'school',
  'высокая интенсивность работы': 'emoji_symbols',
  'сложность поставленных задач': 'account_tree'
}

export default {
  name: 'Priorities',
  components: { Part, Container, Draggable },
  props: { icon: String, value: Object, partName: String },
  data () {
    return {
      icons: Object.freeze(IconByLabel)
    }
  },
  computed: { lv: { get () { return this.value }, set (v) { this.$emit('input', v) } } },
  methods: {
    onPriorityDrop ({ removedIndex, addedIndex, payload }) {
      if (removedIndex === null && addedIndex === null) return

      const result = [...this.lv.p[this.partName]]
      let itemToAdd = payload

      if (removedIndex !== null) {
        itemToAdd = result.splice(removedIndex, 1)[0]
      }

      if (addedIndex !== null) {
        result.splice(addedIndex, 0, itemToAdd)
      }
      this.lv.p[this.partName] = result
      document.body.classList.remove('smooth-dnd-no-user-select', 'smooth-dnd-disable-touch-action')
    }
  }
}
</script>

<style lang="stylus">
  .drag-handler
    cursor grabbing

  .dragged
    box-shadow 0 5px 30px 0 #808080
</style>
