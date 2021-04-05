<template>
  <div class="about">
    <div class="list-products">
      Productos comprados por {{ customer.name }}

      <Carousel v-if="products.length != 0" :products="products"></Carousel>
    </div>

    <div class="list-products">
      Productos que te pueden interesar 

      <Carousel v-if="products.length != 0" :products="suggestions"></Carousel>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Carousel from "../components/Carousel";

export default {
  components: {
    Carousel,
  },
  data: function() {
    return {
      newProduct:{
        nameProduct:"",
        idProduct:"",
        price:0,
      },
      customer: {},
      transactions: [],
      idProducts: [],
      products: [],
      buyers: [],
      suggestions: [],
    };
  },
  async mounted() {

    this.customer = JSON.parse(localStorage.getItem("customerInfo"));
    await axios
      .get("http://localhost:3000/transaction/" + this.customer.id)
      .then((response) => {
        this.transactions = response.data["customer"];
      })
      .catch((error) => {
        console.error(error.message);
      });
     // console.log(this.transactions)
      this.transactions.forEach(transaction => {
        transaction.products.forEach(product=>{
          this.idProducts.push(product)
        })
      })
      this.idProducts.forEach(async id=>{
        await axios
          .get("http://localhost:3000/product/" + id)
          .then( response => {
            this.newProduct={
              nameProduct:response.data["product"][0].nameProduct,
              idProduct:response.data["product"][0].idProduct,
              price:response.data["product"][0].price
            }
            

             this.products.push(this.newProduct);
          });
      })
      console.log(this.products)
/*
    .map((transaction) => {
      transaction.products.map(async (product) => {
        await axios
          .get("http://localhost:3000/product/" + product)
          .then( (response) => {
             this.products.push(response.data["product"][0]);
          });
      });
    });*/

    await axios
      .get("http://localhost:3000/customer/ip/" + this.transactions[0].ip)
      .then((response) => {
        response.data["customer"].map((customer) => {
          axios
            .get("http://localhost:3000/customer/" + customer.buyerId)
            .then((response) => {
              this.buyers.push(response.data["customer"][0]);
            });
        });
      });

    const prices = await this.products.map((product) => product.price);
    const min = Math.min.apply(null, prices);
    const max = Math.max.apply(null, prices);
    await axios
      .get("http://localhost:3000/product/suggestion/" + min + "/" + max)
      .then(response => {
        this.suggestions=(response.data["products"])
      });
  },
};
</script>

<style scoped>
.list-products {
  font-size: 40px;
  margin-top: 40px;
}
</style>
