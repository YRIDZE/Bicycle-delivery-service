<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <supplier
        v-for="supplier in filteredSuppList"
        :key="supplier.id"
        :supp="supplier"
    ></supplier>
  </div>
</template>

<script>

export default {
  computed: {
    filteredSuppList: function () {
      let suppliers = this.$store.getters["supp/getSuppliers"];

      if (this.$store.getters["filter/getSuppTypeFilter"].length !== 0) {
        suppliers = suppliers
            .filter(value => this.$store.getters["filter/getSuppTypeFilter"]
                .includes(value.type));
      }

      if (this.getSuppTypeFilter != null) {
        let timestamp = this.getSuppTypeFilter;
        suppliers = suppliers.filter(value => (this.compareTime(timestamp.opening, value.workingHours.opening)
            && this.compareTime(value.workingHours.closing, timestamp.closing)));
      }

      return suppliers;
    },
    getSuppTypeFilter: function () {
      return this.$store.getters["filter/getSuppTimeFilter"];
    },
  },
  methods: {
    compareTime(str1, str2) {
      let time1 = str1.split(':');
      let time2 = str2.split(':');
      if (time1[0] === "00" || time2[0] === "00" || time1[0] === "24" || time2[0] === "24")
        return true

      return time1[0] >= time2[0] && time1[1] >= time2[1];
    },
  },
};
</script>
