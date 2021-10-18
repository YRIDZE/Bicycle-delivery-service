<template>
  <div>
    <div class="d-flex flex-wrap bd-highlight align-items-stretch">
      <product
          v-for="item in paginatedData"
          :key="item.id"
          :item="item"
      ></product>
    </div>
    <div class="d-flex bd-highlight">
      <div class="clearfix btn-group m-auto mt-8" v-if="pageCount > 1">
        <button class="btn btn-sm btn-outline-secondary" @click="changePageNumber(-1)"
                :disabled="this.$store.getters['prod/getPage'] === 0"> Previous
        </button>
        <button class="btn btn-sm btn-outline-secondary" @click="changePageNumber(1)"
                :disabled="this.$store.getters['prod/getPage'] >= pageCount-1"> Next
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  props: {
    id: {
      type: Number,
      required: true,
    },
  },
  methods: {
    ...mapActions('prod', ['changePageNumber']),
  },
  computed: {
    filteredProdList: function () {
      let products = this.$store.getters["prod/getProducts"];

      if (this.$store.getters["filter/getProdTypeFilter"].length !== 0) {
        products = products
            .filter(value => this.$store.getters["filter/getProdTypeFilter"]
                .includes(value.type));
      }

      if (this.id === "all")
        return products;

      return products.filter(x => x.supplier_id === parseInt(this.id, 10));
    },
    pageCount() {
      let l = this.filteredProdList.length;
      return Math.ceil(l / 12);
    },
    paginatedData() {
      if (this.$router.currentRoute.fullPath === '/all') {
        const start = this.$store.getters["prod/getPage"] * 12;
        const end = start + 12;
        return this.filteredProdList.slice(start, end);
      } else {
        return this.filteredProdList;
      }
    },
  },
  mounted() {
    let path = this.$router.currentRoute.fullPath;
    this.$store.dispatch('prod/getProductTypesBySupp', path[path.length - 1]);
  }
};
</script>

<style scoped>

</style>