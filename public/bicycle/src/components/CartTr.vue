<template>
  <tr>
    <td>{{ index + 1 }}</td>
    <td>{{ i.name }}</td>
    <td>{{ f.name }}</td>
    <td>{{ i.price }}$</td>
    <td>
      <a style="color: #3b3b3b" @click="addProduct({id: i.id, quantity: -1})"
         :disable="item.quantity === 1">
        <font-awesome-icon :icon="['fas', 'minus-circle']"/>
      </a>
      {{ item.quantity }}
      <a style="color: #3b3b3b" @click="addProduct({id: i.id, quantity: 1})">
        <font-awesome-icon :icon="['fas', 'plus-circle']"/>
      </a>
    </td>
    <td></td>
    <td>
      <a @click="remove(index)">
        <font-awesome-icon :icon="['fas', 'trash']"/>
      </a>
    </td>
  </tr>
</template>

<script>

import {mapActions} from "vuex"

export default {
  props: ['item', 'index'],
  methods: {
    ...mapActions(['remove', 'addProduct']),

  },

  computed: {
    i: function () {
      return this.$store.state.items.find(x => x.id == this.item.id)
    },
    f: function () {
      return this.$store.state.restaurants.find(x => x.id == this.i.supplier_id)
    }
  },
}
</script>

<style scoped>

</style>