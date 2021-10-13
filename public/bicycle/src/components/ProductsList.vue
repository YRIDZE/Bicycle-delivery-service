<template>
  <div>
    <div class="d-flex flex-wrap bd-highlight align-items-stretch">
      <product
          v-for="item in paginatedData"
          :key="item.id"
          :item="item"
      ></product>
    </div>
    <div class="clearfix btn-group offset-md-5 py-4 text-center" v-if="pageCount !== 1">
      <button class="btn btn-sm btn-outline-secondary" @click="prevPage"
              :disabled="this.$store.getters['prod/getPage'] === 0"> Previous
      </button>
      <button class="btn btn-sm btn-outline-secondary" @click="nextPage"
              :disabled="this.$store.getters['prod/getPage'] >= pageCount-1"> Next
      </button>
    </div>
  </div>
</template>

<script>


import {mapActions} from "vuex";

export default {
  props: ["id"],
  methods: {
    ...mapActions('prod', ['changePageNumber']),

    nextPage() {
      this.changePageNumber(1);
    },
    prevPage() {
      this.changePageNumber(-1);
    }
  },
  computed: {
    filteredProdList: function () {
      let products = this.$store.getters["prod/getProducts"]

      if (this.$store.getters["filter/getProdTypeFilter"].length !== 0) {
        products = products
            .filter(value => this.$store.getters["filter/getProdTypeFilter"]
                .includes(value.type))
      }

      if (this.id === "all")
        return products;

      return products.filter(x => x.supplier_id === parseInt(this.id, 10));
    },
    pageCount() {
      let l = this.filteredProdList.length
      return Math.ceil(l / 12);
    },
    paginatedData() {
      const start = this.$store.getters["prod/getPage"] * 12;
      const end = start + 12;
      return this.filteredProdList.slice(start, end);
    }
  },
};
</script>

<style scoped>

</style>