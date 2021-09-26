<template>
  <tr>
    <td>{{ index + 1 }}</td>
    <td>{{ i.title }}</td>
    <td>{{ f.title }}</td>
    <td>{{ i.price }}</td>
    <td>
      <a style="color: #3b3b3b" @click="reduce()" :disable="item.quantity === 1">
        <font-awesome-icon :icon="['fas', 'minus-circle']"/>
      </a>
      {{ item.quantity }}
      <a style="color: #3b3b3b" @click="add()">
        <font-awesome-icon :icon="['fas', 'plus-circle']"/>
      </a>
    </td>
    <td></td>
    <td>
      <a @click="del(index)">
        <font-awesome-icon :icon="['fas', 'trash']"/>
      </a>
    </td>
  </tr>
</template>

<script>

export default {

  props: ['item', 'index'],
  methods: {
    add: function () {
      this.$store.commit("add", {
            id: this.item.id,
            quantity: 1
          }
      )
    },

    reduce: function () {
      if (this.item.quantity === 1) return;
      this.$store.commit("add", {
            id: this.item.id,
            quantity: -1
          }
      )
    },

    del: function (index) {
      this.$store.commit("removeFromCart", {
            index: index,
          }
      )
    }
  },

  computed: {
    i: function () {
      return this.$store.state.items.find(x => x.id == this.item.id)
    },
    f: function () {
      return this.$store.state.restaurants.find(x => x.id == this.i.restaurantId)
    }
  },
  name: "CartTr"
}
</script>

<style scoped>

</style>