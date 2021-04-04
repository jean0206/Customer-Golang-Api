<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
</template>

<script>
export default {
  beforeMount() {
    const customer = JSON.parse(localStorage.getItem("customerInfo"))
    const cosas = []
    console.log(customer.name)
  },
}
</script>