<template>
  <div class="card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <div class="accordion-item">
        <h2 class="accordion-header" id="panelsStayOpen-headingOne">
          <button class="accordion-button" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelsStayOpen-collapseOne" aria-expanded="true"
                  aria-controls="panelsStayOpen-collapseOne">
            Accordion Item #1
          </button>
        </h2>
        <div id="panelsStayOpen-collapseOne" class="accordion-collapse collapse show"
             aria-labelledby="panelsStayOpen-headingOne">
          <div class="accordion-body">
            <div class="d-flex flex-row form-row">
              <div class="form-group m-1">
                <label>Min</label>
                <input class="form-control" placeholder="$0" type="number">
              </div>
              <div class="form-group text-right m-1">
                <label>Max</label>
                <input class="form-control" placeholder="$1,0000" type="number">
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panelsStayOpen-headingTwo">
          <button class="accordion-button" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelsStayOpen-collapseTwo" aria-expanded="true"
                  aria-controls="panelsStayOpen-collapseTwo">
            Accordion Item #2
          </button>
        </h2>
        <div id="panelsStayOpen-collapseTwo" class="accordion-collapse collapse show"
             aria-labelledby="panelsStayOpen-headingTwo">
          <div class="accordion-body">
            <div class="filter-content">
              <div class="card-body">
                <Check name="More cheese"></Check>
                <Check name="More vegetables"></Check>
                <Check name="More sauce"></Check>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Check from './Check'

export default {
  components: {
    Check
  }
}
</script>

<style scoped>

</style>