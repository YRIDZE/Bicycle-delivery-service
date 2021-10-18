<template>
  <tr>
    <td class="text-lg">{{ index + 1 }}</td>
    <td class="text-lg">{{ i.name }}</td>
    <td class="text-lg">{{ f.name }}</td>
    <td class="text-lg">{{ i.price }}$</td>
    <td class="text-lg">
      <a style="color: #3b3b3b" @click="addItem({product_id: i.id, quantity: -1})"
         :disable="item.quantity === 1">
        <font-awesome-icon :icon="['fas', 'minus-circle']"/>
      </a>
      {{ item.quantity }}
      <a style="color: #3b3b3b" @click="addItem({product_id: i.id, quantity: 1})">
        <font-awesome-icon :icon="['fas', 'plus-circle']"/>
      </a>
    </td>
    <td></td>
    <td class="text-lg">
      <a @click="removeItem(item)">
        <font-awesome-icon :icon="['fas', 'trash']"/>
      </a>
    </td>
  </tr>
</template>

<script>

import {mapActions} from "vuex"

export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
    index: {
      type: Number,
      required: true,
    }
  },
  methods: {
    ...mapActions('cart', ['removeItem', 'addItem']),
  },

  computed: {
    i: function () {
      return this.$store.getters["prod/getProducts"].find(x => x.id == this.item.product_id)
    },
    f: function () {
      return this.$store.getters["supp/getSuppliers"].find(x => x.id == this.i.supplier_id)
    }
  },
}
</script>

<style scoped>

</style>